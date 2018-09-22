package main

import (
	"github.com/OdaDaisuke/go-algorithm/dijkstra"
	"github.com/OdaDaisuke/go-algorithm/goroutine"
)

func main() {
	number := []int{1,2,3,4,5}
	goroutine.Serial(number)
	goroutine.Parallel(number)
	goroutine.ParallelChannel(number)

	dijkstra.Start()
}