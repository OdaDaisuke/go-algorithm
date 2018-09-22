package main

import (
	"github.com/OdaDaisuke/go-algorithm/image_processor"
	"github.com/OdaDaisuke/go-algorithm/dijkstra"
	"github.com/OdaDaisuke/go-algorithm/goroutine"
	"fmt"
)

func main() {
	startGoroutine()
	startDijkstra()
	startImageProcess()
}

func startGoroutine() {
	number := []int{1,2,3,4,5}
	goroutine.Serial(number)
	goroutine.Parallel(number)
	goroutine.ParallelChannel(number)
}

func startImageProcess() {
	ip := image_processor.NewImageProcessor("./assets/flower_1.jpeg")
	ip.Init()
	defer ip.CloseFile()

	ip.GrayScale()
	fmt.Println(ip.Exif)
}

func startDijkstra() {
	dijkstra.Start()
}