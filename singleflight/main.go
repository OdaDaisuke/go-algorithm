package sf

import (
	"fmt"

	"golang.org/x/sync/singleflight"
)

func Run() {
	var g singleflight.Group
	// 例えばHTTPサーバで同時にリクエストがきて同じ関数が実行された時
	// 最初だけ実行して、他のリクエストは待機させ、最初のリクエストが完了した時に同じ結果を渡してくれる。
	// HOL Blocking気になる。
	res, err, _ := g.Do("key", func() (interface{}, error) {
		return nil, nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("res ", res)
}
