package http

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/justcgh9/zvukovat/services/users/internal/http/auth/handlers"
	"github.com/justcgh9/zvukovat/services/users/internal/services/auth"
)

type App struct{
    log *slog.Logger
    httpServer *http.Server
    port int
}

func New(
    log *slog.Logger,
    port int,
    timeout time.Duration,
    iddleTimeout time.Duration,
    authSrvc *auth.Auth,
) *App {
    router := chi.NewRouter()

    router.Route("/users", func(router chi.Router) {
        router.Post("/signup", handlers.NewSignUp(log, authSrvc, timeout, fmt.Sprintf("http://localhost:%d", port)))
        router.Post("/signin", handlers.NewSignIn(log, authSrvc, timeout))
    })

    return &App{
        log: log,
        port: port,
        httpServer: &http.Server{
            Addr: fmt.Sprintf(":%d", port),
            Handler: router,
            ReadTimeout: timeout,
            WriteTimeout: timeout,
            IdleTimeout: iddleTimeout,
        },

    }

}

func (a *App) MustRun() {
    if err := a.Run(); err != nil {
        panic(err)
    }
}

func (a *App)  Run() error {
    const op = "http.Run"

    log := a.log.With(
            slog.String("op", op),
        slog.Int("port", a.port),
        )

    log.Info("starting HTTP server", slog.String("addr", a.httpServer.Addr))
    if err := a.httpServer.ListenAndServe(); err != nil {
        log.Error("error listening http", slog.String("error", err.Error()))
        return fmt.Errorf("%s: %w", op, err)
    }

    return nil
}

func (a *App) Stop() {
    const op = "http.Stop"

    a.log.With(slog.String("op", op)).
        Info("stopping HTTP server", slog.Int("port", a.port))

    a.httpServer.Shutdown(context.Background())
}
