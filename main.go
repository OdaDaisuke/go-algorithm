package main

import (
	"github.com/OdaDaisuke/go-algorithm/dijkstra"
	"github.com/OdaDaisuke/go-algorithm/goroutine"
	"github.com/OdaDaisuke/go-algorithm/image_processor"
	"github.com/OdaDaisuke/go-algorithm/go_tour"
	"github.com/OdaDaisuke/go-algorithm/net_http"
)

func main() {
	startGoroutine()
	startDijkstra()
	startImageProcess()
	startGoTour()
	startNetHttp()
}

func startGoroutine() {
	goroutine.Start()
}

func startImageProcess() {
	image_processor.Start()
}

func startDijkstra() {
	dijkstra.Start()
}

func startGoTour() {
	go_tour.Start()
}

func startNetHttp() {
	net_http.Start()
}