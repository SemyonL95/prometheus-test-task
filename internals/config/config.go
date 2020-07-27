package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	MetricsPort int
	LogsPort    int
}

func New() *Config {
	rawmp := os.Getenv("METRICS_PORT")
	if rawmp == "" {
		log.Fatal("please provide METRICS_PORT env variable")
	}

	mp, err := strconv.Atoi(rawmp)
	if err != nil {
		log.Fatal("METRIC_PORT should to be valid int")
	}

	rawlp := os.Getenv("LOGS_PORT")
	if rawmp == "" {
		log.Fatal("please provide LOGS_PORT env variable")
	}

	lp, err := strconv.Atoi(rawlp)
	if err != nil {
		log.Fatal("LOGS_PORT should to be valid int")
	}

	return &Config{
		MetricsPort: mp,
		LogsPort:    lp,
	}
}
