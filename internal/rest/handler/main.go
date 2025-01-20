package handler

import "log/slog"

type Handler struct {
}

func New(logger *slog.Logger) *Handler {
	return &Handler{}
}
