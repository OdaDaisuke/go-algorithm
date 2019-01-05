package main

import (
	"github.com/OdaDaisuke/go-algorithm/dijkstra"
	"github.com/OdaDaisuke/go-algorithm/go_tour"
	"github.com/OdaDaisuke/go-algorithm/goroutine"
	"github.com/OdaDaisuke/go-algorithm/image_processor"
	"github.com/OdaDaisuke/go-algorithm/net_http"
	"github.com/OdaDaisuke/go-algorithm/read_json"
	"github.com/OdaDaisuke/go-algorithm/redis"
)

func main() {
	app_redis.Start()
	net_http.Start()
	goroutine.Start()
	dijkstra.Start()
	image_processor.Start()
	go_tour.Start()
	readjson.Start()
}