package config

import (
	"context"
	"flag"
	"os"
	"strconv"
)

type PostgresConfig struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
}

func newPostgresConfig(ctx context.Context) *PostgresConfig {
	host := os.Getenv("PG_HOST")
	portStr := os.Getenv("PG_PORT")
	port, _ := strconv.Atoi(portStr)
	name := os.Getenv("PG_NAME")
	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASSWORD")

	flag.StringVar(&host, "pg-host", host, "postgreSQL host [PG_HOST]")
	flag.IntVar(&port, "pg-port", port, "postgreSQL port [PG_PORT]")
	flag.StringVar(&name, "pg-name", name, "postgreSQL name [PG_NAME]")
	flag.StringVar(&user, "pg-user", user, "postgreSQL user [PG_USER]")
	flag.StringVar(&password, "pg-password", password, "postgreSQL password [PG_PASSWORD]")

	return &PostgresConfig{
		Host:     host,
		Port:     port,
		Name:     name,
		User:     user,
		Password: password,
	}
}
