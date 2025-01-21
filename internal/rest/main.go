package rest

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"time"

	"github.com/Madinab99999/blogging-platform/internal/config"
	"github.com/Madinab99999/blogging-platform/internal/rest/handler"
	"github.com/Madinab99999/blogging-platform/internal/rest/router"
)

type Rest struct {
	conf   *config.APIRestConfig
	logger *slog.Logger
	router *router.Router
}

func New(conf *config.APIRestConfig, logger *slog.Logger) *Rest {
	h := handler.New(logger.With(slog.String("module", "handler")))
	r := router.New(h)
	return &Rest{
		conf:   conf,
		logger: logger,
		router: r,
	}
}

func (a *Rest) Start(conf *config.APIRestConfig, ctx context.Context) error {
	if err := validateConfig(conf); err != nil {
		return err
	}
	mux := a.router.Start(ctx)

	errLogger := slog.NewLogLogger(a.logger.Handler(), slog.LevelError)
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", a.conf.Port),
		Handler:           mux,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
		ErrorLog: errLogger,
	}

	a.logger.InfoContext(ctx, "starting server", slog.Int("port", a.conf.Port))

	if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func validateConfig(conf *config.APIRestConfig) error {
	required := map[string]string{
		"host": conf.Host,
		"port": fmt.Sprintf("%d", conf.Port),
	}
	for key, value := range required {
		if value == "" || (key == "port" && conf.Port == 0) {
			return fmt.Errorf("missing config REST %s", key)
		}
	}
	port := conf.Port
	if port < 1 || port > 65535 {
		return fmt.Errorf("invalid port number: %d", port)
	}
	return nil
}
