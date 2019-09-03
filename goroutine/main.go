package goroutine

import (
	"fmt"
	"sync"
	"time"
)

/*
goroutineは最小で2048byte。

OSのスレッドごとに、タスクであるgoroutineのスケジューラ・リストがある。
現代のカーネルはCPUのM個のコア数に対して同時にN個の複数の処理ができるM:Nモデルになっていて
goroutineもそれと同じM:Nモデル。OSのスレッド数Mに対してM:Nになってる。
スレッドの種類はM:N以外に
- N:1(複数のユーザ空間スレッドが１つのOSスレッドで実行)
- 1:1(1つの実行スレッドが１つのOSスレッドと一致)
がある。

[Goroutineスケジューリングの話]
- M(machine)
- G(goroutine)
- P(Processor)
  - Gをためておくキューを持ってる。
  - GをMに割り当てる責務がある
- グローバルキュー
  - Pが持つGのキューとは別に定義
  - 通常Gが実行待ち状態になるとPのキューに入るが、いくつかの状況ではグローバルキューに入る
  - グローバルキューが取り出されるタイミング
    - Mの自分のキューがからのとき
    - Mの自分のキューがまだあったとしても61回に一度取り出す(これはグローバルキューがずっと実行されないことを防ぐ)
- sysmon
  - GOMAXPROCSの数だけMとPのセットが動いてるが、これと別にsysmonという関数を実行し続けるだけの特別なMが存在する。これはPなしで実行される。
- P idle list
  - Pが暇になったら行く場所
- M idle list
  - Mが暇になったら行く場所

GOMAXPROCSの数だけ、M, Pのセットが用意される。
PがGをMに割り当てて、実行されたらGはGCされる。
この時実行中のどこかのPがGを実行しきって暇になったら、他のPからGを半分だけ奪い取る。(work stealingアルゴリズム)。
work stealingアルゴリズムは、CPUバウンドな処理は非常に効率が良いが、io waitなどの無駄にブロックする処理とは相性が悪い。
のでいくつかの工夫がなされてる。
↓

[syscall実行時]
sysclallをして時間がかかっている場合、sysmonがそれを検知して対象のMとGをPから切り離し、
別にMとPを紐づけて処理を進める。
このとき紐づけるMとPは、なければ作成する。
syscallが終わったら最初にP idle listの中身を調べ、暇なPが存在したらそいつを取ってきて処理を進める。
それが無理ならグローバルキューに追加し、MはM idle listに追加される。

[ネットワーク処理時]
ネットワーク待ち状態になったら、対象のGがMから切り離されnetpollerに登録される。
その後sysmonがnetpollerを定期的にチェックし、ネットワーク処理が準備完了してたら
そのGをグローバルキューにエンキューする。

[channelのreceive/send時]
channel1というチャンネルがあるとする。
Gがchannel1のreceiveを待つと、channl1の待ちGリストにエンキューされる。
別にGがchannel1にsendすると、channel1の待ちGリストからデキューされて、
send元のGのPのキューにエンキューされる。

ちなみにgo runコマンドにオプションで-raceをつけるとデータ競合をチェックしてくれる。
  -> https://golang.org/doc/articles/race_detector.html
 */

func Start() {
	number := []int{1,2,3,4,5}
	serial(number)
	parallel(number)
	parallelChannel(number)
}

func serial(num []int) {
	for _, n := range num {
		fmt.Printf("serial: %d\n", n)
	}
}

func parallel(num []int) {
	fmt.Println()
	var wp sync.WaitGroup

	for _, n := range num {
		wp.Add(1)

		go func(n int) {
			fmt.Printf("parallel: %d\n", n)
			wp.Done()
		}(n)

	}

	wp.Wait()
}

func parallelChannel(num []int) {
	res := make(chan int, len(num))
	go func() {
		defer close(res)
		for _, n := range num {
			fmt.Printf("parallel channel push: %d\n", n)
			res <- n
			time.Sleep(time.Millisecond)
		}
	}()

	for r := range res {
		fmt.Printf("parallel channel pull: %d\n", r)
	}
}
