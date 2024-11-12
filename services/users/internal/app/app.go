package app

import (
	"context"
	"log/slog"
	"time"

	grpcapp "github.com/justcgh9/zvukovat/services/users/internal/app/grpc"
	httpapp "github.com/justcgh9/zvukovat/services/users/internal/app/http"
	"github.com/justcgh9/zvukovat/services/users/internal/services/auth"
	"github.com/justcgh9/zvukovat/services/users/internal/services/email"
	"github.com/justcgh9/zvukovat/services/users/internal/storage/mongo"
)

type App struct {
    GRPCSrv grpcapp.App
    HTTPSrv httpapp.App
}

func New(
    log *slog.Logger,
    grpcPort int,
    httpPort int,
    storagePath, dbName string,
    emailHst, emailUsr, emailPsswd string,
    emailPort int,
    accessTokenTTL time.Duration,
    refreshTokenTTL time.Duration,
    accessSecret, refreshSecret string,
    grpcTimeout time.Duration,
    httpTimeout time.Duration,
    httpIdleTimeout time.Duration,

) *App {
    //TODO: init storage
    ctx, cancel := context.WithTimeout(context.Background(), httpTimeout)
    defer cancel()
    storage, err := mongo.New(
        ctx,
        storagePath,
        dbName,
    )
    if err != nil {
        panic(err)
    }
    //TODO: init service
    emailSrvc := email.New(emailUsr, emailHst, emailPsswd, emailPort)
    authSrvc := auth.New(
        log,
        storage,
        storage,
        storage,
        storage,
        storage,
        storage,
        emailSrvc,
        accessTokenTTL,
        refreshTokenTTL,
    )
    grpcApp := grpcapp.New(log, grpcPort)
    httpApp := httpapp.New(log, httpPort, httpTimeout, httpIdleTimeout, authSrvc, accessSecret, refreshSecret)

    return &App{
        GRPCSrv: *grpcApp,
        HTTPSrv: *httpApp,
    }
}
