package dijkstra

import "fmt"

func Start() {
	// 有向グラフを作る
	g := NewDirectedGraph()

	// グラフ定義
	g.Add("s", "a", 2)
	g.Add("s", "b", 5)
	g.Add("a", "b", 2)
	g.Add("a", "c", 5)
	g.Add("b", "c", 4)
	g.Add("b", "d", 2)
	g.Add("c", "z", 7)
	g.Add("d", "c", 5)
	g.Add("d", "z", 2)

	// sノードからzノードへの最短経路を得る
	path, err := g.ShortestPath("s", "z")

	// 経路が見つからなければ終了
	if err != nil {
		fmt.Println("Goal not found")
		return
	}

	// 見つかった経路からノードとコストを表示する
	for _, node := range path {
		fmt.Printf("ノード: %v, コスト: %v\n", node.name, node.cost)
	}
}
