package config

import (
	"context"
	"flag"
	"os"
)

type APIConfig struct {
	TokenSecret string
	Rest        *APIRestConfig
}

func newApiConfig(ctx context.Context) *APIConfig {
	tokenSecret := os.Getenv("API_TOKEN_SECRET")

	flag.StringVar(&tokenSecret, "api-token-secret", tokenSecret, "api token secret [API_TOKEN_SECRET]")

	return &APIConfig{
		TokenSecret: tokenSecret,
		Rest:        newApiRestConfig(ctx),
	}
}
