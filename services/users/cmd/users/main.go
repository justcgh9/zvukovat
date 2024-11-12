package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/justcgh9/zvukovat/services/users/internal/app"
	"github.com/justcgh9/zvukovat/services/users/internal/config"
)

const (
    envLocal = "local"
    envProd = "prod"
)

func main() {
    err := godotenv.Load()

    cfg := config.MustLoad()

    log := setupLogger(cfg.Env)

    if err != nil {
        log.Error("Error loading .env file")
    }


    appl := app.New(
        log,
        cfg.GRPC.Port,
        cfg.HTTP.Port,
        cfg.StoragePath,
        cfg.DBName,
        cfg.Email.Host,
        cfg.Email.User,
        cfg.Email.Password,
        cfg.Email.Port,
        cfg.AccessTokenTTL,
        cfg.RefreshTokenTTL,
        cfg.JwtAccessSecret,
        cfg.JwtRefreshSecret,
        cfg.GRPC.Timeout,
        cfg.HTTP.Timeout,
        cfg.HTTP.IdleTimeout,
    )

    go appl.GRPCSrv.MustRun()
    go appl.HTTPSrv.MustRun()

    stop := make(chan os.Signal, 1)
    signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

    <-stop
    appl.GRPCSrv.Stop()
    appl.HTTPSrv.Stop()
    log.Info("application stopped")
}

func setupLogger(env string) *slog.Logger {
    var log *slog.Logger

    switch env {
    case envLocal:
        log = slog.New(
            slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
            )
    case envProd:
        log = slog.New(
            slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
            )
    }
    return log
}
