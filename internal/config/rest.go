package config

import (
	"context"
	"flag"
	"os"
	"strconv"
)

type APIRestConfig struct {
	Host string
	Port int
}

func newApiRestConfig(_ context.Context) *APIRestConfig {
	host := os.Getenv("API_REST_HOST")
	portStr := os.Getenv("API_REST_PORT")
	port, _ := strconv.Atoi(portStr)

	c := &APIRestConfig{
		Host: host,
		Port: port,
	}

	flag.StringVar(&c.Host, "api-rest-host", c.Host, "api rest host [API_REST_HOST]")
	flag.IntVar(&c.Port, "api-rest-port", c.Port, "api rest port [API_REST_PORT]")

	return c
}
