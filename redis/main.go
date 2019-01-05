package app_redis

import (
	"github.com/gomodule/redigo/redis"
	"fmt"
	"time"
)

// cf. http://redis.shibu.jp/commandreference/lists.html
// SET 文字列型の追加
// LPUSH 新しい要素をリストの左(先頭)に追加
// RPUSH 新しい要素をリストの右(最後尾)に追加
// LINDEX リストのindexが指している要素を取得
// LSET リストの指定されたインデックスの値を新しい値に更新
// LREM リストの指定されたインデックスの値を削除
// LRANGE リストの指定された範囲のインデックスの要素を取得
// SADD セット型に要素を追加
// SISMEMBER セット型の任意の要素が存在するか
// SMEMBERS セット型の全体の要素を取得
// WATCH 並行処理などする場合に、WATCHで値を監視しといて競合があったらUNWATCHコマンドを送ってトランザクションキューをフラッシュする。
// HGETALL ハッシュ型の全てのフィールドとあたいのペアを得る
// HSET ハッシュにkeyValを追加
// PUBLISH チャンネルをpublishする

// ↓トランザクション↓
// MULTI
// ...コマンド... -> conn.Sendでクライアントの送信バッファにコマンドを貯める。キューみたいな感じ。
// DISCARD -> トランザクションキューの内容がフラッシュされてトランザクションが終了する。故意に終了したいときに使う。
// EXEC

// アクティブコネクション数 = 使用中+アイドル

const addr = "localhost:6379"

func newPool(addr string)  *redis.Pool {
	return &redis.Pool{
		MaxIdle: 3,
		MaxActive: 0,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "addr")
		},
	}
}

func getPool() redis.Conn {
	pool := newPool(addr)
	return pool.Get()
}

func Start() {
	//conn := getPool()
	//defer conn.Close()
	conn, err := redis.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	const key, listKey, setKey = "programs", "lessons", "numbers"

	// Write value
	r, err := conn.Do("SET", key, "[{\"name\": \"pilates\"}]")
	if err != nil {
		panic(err)
	}
	fmt.Println(r)

	// With list
	_, err = conn.Do("LPUSH", listKey, "a")
	_, err = conn.Do("RPUSH", listKey, "b")
	_, err = conn.Do("LPUSH", listKey, "c")

	// With set
	_, err = conn.Do("SADD", setKey, "one")
	_, err = conn.Do("SADD", setKey, "two")
	_, err = conn.Do("SADD", setKey, "three")


	// Read value
	s, err := redis.String(conn.Do("GET", key))
	if err != nil {
		panic(err)
	}
	fmt.Println(s)

	a, err := redis.String(conn.Do("LINDEX", listKey, 1))
	if err != nil {
		panic(err)
	}
	fmt.Println(a)

	b, err := redis.Strings(conn.Do("LRANGE", listKey, 0, -1))
	if err != nil {
		panic(err)
	}
	fmt.Println(b)

	m, err := redis.Int(conn.Do("SREM", setKey, "two"))
	if err != nil {
		panic(err)
	}
	fmt.Println(m)

	// If true, then 1.
	l, err := redis.Int(conn.Do("SISMEMBER", setKey, "three"))
	if err != nil {
		panic(err)
	}
	fmt.Println(l)

	allSet, err := redis.Strings(conn.Do("SMEMBERS", setKey))
	if err != nil {
		panic(err)
	}
	fmt.Println(allSet)

	// Start transaction
	err = conn.Send("MULTI")
	if err != nil {
		panic(err)
	}

	err = conn.Send("SET", "money", 1000)
	if err != nil {
		panic(err)
	}

	err = conn.Send("SADD", "fruits", "apple")
	if err != nil {
		panic(err)
	}

	// exec command
	txVal, err := redis.Values(conn.Do("EXEC"))
	if err != nil {
		panic(err)
	}
	fmt.Println(txVal) // If OK, then 1

	// Pub, Sub

	// subscriber
	psc := redis.PubSubConn{Conn :conn }
	psc.Subscribe("channel_1", "channel_2", "channel_3")
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			fmt.Printf("%s: message: $s\n", v.Channel, v.Data)
		case redis.Subscription:
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			return
		}
	}

	// publisher
	p, err := redis.Int(conn.Do("PUBLISH", "channel_1", "hello"))
	if err != nil {
		panic(err)
	}
	fmt.Println(p)
}