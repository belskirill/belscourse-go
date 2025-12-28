package main

import (
	"belscourrsego/internal/app"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	app, err := app.NewApp(ctx)
	if err != nil {
		log.Fatal(err)
	}

	server := app.BuilderHTTPServer()

	go func() {
		app.Logger.Info("HTTP server on 127.0.0.1:8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		app.Logger.Error("http shutdown error:", zap.Error(err))
	}

	app.Close()

}
