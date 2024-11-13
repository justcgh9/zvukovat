package grpc

import (
	"fmt"
	"log/slog"
	"net"

	authrpc "github.com/justcgh9/zvukovat/services/users/internal/grpc/auth"
	"google.golang.org/grpc"
)

type App struct {
    log *slog.Logger
    gRPCServer *grpc.Server
    port int
}

func New(
    log *slog.Logger,
    port int,
    secret string,
) *App {
    gRPCServer := grpc.NewServer()

    authrpc.Register(gRPCServer, secret)

    return &App{
        log: log,
        gRPCServer: gRPCServer,
        port: port,
    }
}

func (a *App) MustRun() {
    if err := a.Run(); err != nil {
        panic(err)
    }
}

func (a *App) Run() error {
    const op = "grpc.Run"

    log := a.log.With(
            slog.String("op", op),
            slog.Int("port", a.port),
        )

    listener, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
    if err != nil {
        log.Error("error connecting to tcp", slog.String("error", err.Error()))
        return fmt.Errorf("%s: %w", op, err)
    }

    log.Info("starting gRPC server", slog.String("addr", listener.Addr().String()))

    if err := a.gRPCServer.Serve(listener); err != nil {
        log.Error("error listening tcp", slog.String("error", err.Error()))
        return fmt.Errorf("%s: %w", op, err)
    }

    return nil
}

func (a *App) Stop() {
    const op = "grpc.Stop"

    a.log.With(slog.String("op", op)).
        Info("stopping gRPC server", slog.Int("port", a.port))

    a.gRPCServer.GracefulStop()
}
