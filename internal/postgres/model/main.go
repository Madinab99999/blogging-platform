package model

import (
	"database/sql"
	"log/slog"

	"github.com/Madinab99999/blogging-platform/internal/config"
)

type Model struct {
}

func New(conf *config.PostgresConfig, logger *slog.Logger, db *sql.DB) *Model {
	return &Model{}
}
