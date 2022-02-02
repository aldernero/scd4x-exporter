package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	sensorCollector := newScd4xCollector()
	prometheus.MustRegister(sensorCollector)
	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("Beginning to serve on port :9110")
	http.ListenAndServe(":9110", nil)
}
