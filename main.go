package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"time"

	"github.com/Madinab99999/blogging-platform/internal/config"
	"github.com/Madinab99999/blogging-platform/internal/postgres"
	"github.com/Madinab99999/blogging-platform/internal/rest"
	"github.com/Madinab99999/blogging-platform/pkg/logger"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// config
	conf := config.New(ctx)

	// logger
	log := logger.New(logger.Options{
		Format:     logger.TextFormat,
		Level:      slog.LevelDebug,
		AddSource:  false,
		TimeFormat: time.DateTime,
	})

	//postgres
	_, err := postgres.New(conf.Postgres, log.With(slog.String("service", "postgre")))
	if err != nil {
		log.ErrorContext(ctx, "failed to start postgre", slog.Any("error", err))
		panic(err)
	}

	//rest
	r := rest.New(conf.API.Rest, log.With(slog.String("service", "rest")))
	go func(ctx context.Context, cancelFunc context.CancelFunc) {
		if err := r.Start(conf.API.Rest, ctx); err != nil {
			log.ErrorContext(ctx, "failed to start rest", slog.Any("error", err))
		}

		cancelFunc()
	}(ctx, cancel)

	go func(cancelFunc context.CancelFunc) {
		shutdown := make(chan os.Signal, 1)   // Create channel to signify s signal being sent
		signal.Notify(shutdown, os.Interrupt) // When an interrupt is sent, notify the channel

		sig := <-shutdown
		log.WarnContext(ctx, "signal received - shutting down...", slog.Any("signal", sig))

		cancelFunc()
	}(cancel)

	<-ctx.Done()

	fmt.Println("shutting down gracefully")
}
