package main

import (
	"fmt"
	"log"
	"net/http"
	"prometheus-test-task/internals/api"
	"prometheus-test-task/internals/cache"
	"prometheus-test-task/internals/config"
	"prometheus-test-task/internals/metrics"
)

func main() {

	conf := config.New()

	log.Printf("config: %+v", conf)

	cache := cache.New()
	metrics := metrics.New()
	handlers := api.New(cache, metrics)

	go func() {
		log.Print("Starting metrics server")

		metricsMux := http.NewServeMux()
		metricsMux.HandleFunc("/metrics", handlers.HandleMetrics)
		metricsServer := http.Server{
			Addr:    fmt.Sprintf(":%d", conf.MetricsPort),
			Handler: metricsMux,
		}
		err := metricsServer.ListenAndServe()
		if err != nil {
			log.Fatal("Cannot start metrics server")
		}
	}()

	log.Print("Starting logs server")
	apiMux := http.NewServeMux()
	apiMux.HandleFunc("/logs", handlers.HandleLogs)
	apiServer := http.Server{
		Addr:    fmt.Sprintf(":%d", conf.LogsPort),
		Handler: apiMux,
	}
	err := apiServer.ListenAndServe()
	if err != nil {
		log.Fatal("Cannot start api server")
	}
}
