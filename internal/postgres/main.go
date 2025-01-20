package postgres

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/Madinab99999/blogging-platform/internal/config"
	"github.com/Madinab99999/blogging-platform/internal/postgres/model"
	_ "github.com/lib/pq"
)

type Postgres struct {
	*model.Model
}

func New(conf *config.PostgresConfig, logger *slog.Logger) (*Postgres, error) {
	if err := validateConfig(conf); err != nil {
		return nil, err
	}
	db, err := NewDB(conf)
	if err != nil {
		return nil, err
	}

	return &Postgres{
		Model: model.New(conf, logger.With(slog.String("module", "model")), db),
	}, nil
}

func NewDB(conf *config.PostgresConfig) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Password, conf.Name,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	return db, nil
}

func (p *Postgres) Close(db *sql.DB) error {
	if err := db.Close(); err != nil {
		return fmt.Errorf("failed to close database connection: %w", err)
	}
	return nil
}

func validateConfig(conf *config.PostgresConfig) error {
	required := map[string]string{
		"host":     conf.Host,
		"port":     fmt.Sprintf("%d", conf.Port),
		"name":     conf.Name,
		"user":     conf.User,
		"password": conf.Password,
	}
	for key, value := range required {
		if value == "" || (key == "port" && conf.Port == 0) {
			return fmt.Errorf("missing config PostgreSQL %s", key)
		}
	}
	port := conf.Port
	if port < 1 || port > 65535 {
		return fmt.Errorf("invalid port number: %d", port)
	}
	return nil
}
