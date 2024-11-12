package handlers

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"

    "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/justcgh9/zvukovat/services/users/internal/domain/models"
	"github.com/justcgh9/zvukovat/services/users/internal/lib/jwt"
)

const (
    TOKEN_INDEX_IN_HEADER = 1
    EMPTY_COOKIE_MAX_AGE = -1
)

type Registrator interface {
    SignUp(ctx context.Context, usr models.User, domainName string) (map[string]string, models.UserDTO, error)
}

func NewSignUp(
        log *slog.Logger,
        registrator Registrator,
        timeout time.Duration,
        domainName string,
        accessSecret string,
        refreshSecret string,
    ) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        const op = "http.auth.NewSignUp"

        log := log.With(
            slog.String("op", op),
            slog.String("request_id", middleware.GetReqID(r.Context())),
        )
        var user models.User
        err := render.DecodeJSON(r.Body, &user)
        if err != nil {
            log.Error("error decoding request body", slog.String("error", err.Error()))
            render.Status(r, 400)
            render.JSON(w, r, err.Error())
            return
        }

        ctx, cancel := context.WithTimeout(context.Background(), timeout)
        defer cancel()

        tkns, usr, err := registrator.SignUp(ctx, user, domainName)
        if err != nil {
            log.Error("error signing up", slog.String("error", err.Error()))
            render.Status(r, 400)
            render.JSON(w, r, err.Error())
            return
        }

        expires, err := jwt.GetExpiryDate(tkns["refreshToken"])
        if err != nil {
            log.Error("error signing up", slog.String("error", err.Error()))
            render.Status(r, 400)
            render.JSON(w, r, err.Error())
            return
        }
        cookie := http.Cookie{
            Name: "refreshToken",
            Value: tkns["refreshToken"],
            HttpOnly: true,
            Secure: true,
            Path: "/",
            Expires: expires,
        }

        w.Header().Add("Set-Cookie", fmt.Sprintf("%s;Partitioned", cookie.String()))
        render.Status(r, 201)
        response := make(map[string]interface{})
        response["tokens"] = tkns
        response["user"] = usr
        render.JSON(w, r, response)
    }
}


type Loginer interface {
    SignIn(ctx context.Context, user models.User) (map[string]string, models.UserDTO, error)
}

func NewSignIn(
    log *slog.Logger,
    loginer Loginer,
    timeout time.Duration,
    accessSecret string,
    refreshSecret string,
) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        const op = "http.auth.NewSignIn"

        log := log.With(
            slog.String("op", op),
            slog.String("request_id", middleware.GetReqID(r.Context())),
        )
        var user models.User
        err := render.DecodeJSON(r.Body, &user)
        if err != nil {
            log.Error("error decoding request body", slog.String("error", err.Error()))
            render.Status(r, 400)
            render.JSON(w, r, err.Error())
            return
        }

        ctx, cancel := context.WithTimeout(context.Background(), timeout)
        defer cancel()

        tkns, usr, err := loginer.SignIn(ctx, user)
        if err != nil {
            log.Error("error decoding request body", slog.String("error", err.Error()))
            render.Status(r, 400)
            render.JSON(w, r, err.Error())
            return
        }

        expires, err := jwt.GetExpiryDate(tkns["refreshToken"])
        if err != nil {
            log.Error("error signing up", slog.String("error", err.Error()))
            render.Status(r, 400)
            render.JSON(w, r, err.Error())
            return
        }
        cookie := http.Cookie{
            Name: "refreshToken",
            Value: tkns["refreshToken"],
            HttpOnly: true,
            Secure: true,
            Path: "/",
            Expires: expires,
        }

        w.Header().Add("Set-Cookie", fmt.Sprintf("%s;Partitioned", cookie.String()))
        render.Status(r, 201)
        response := make(map[string]interface{})
        response["tokens"] = tkns
        response["user"] = usr
        render.JSON(w, r, response)
    }
}

type SignOuter interface {
    SignOut(ctx context.Context, token string) error
}

func NewSignOut(
    log *slog.Logger,
    signOuter SignOuter,
    timeout time.Duration,
)   http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        const op = "http.auth.NewSignOut"

        log := log.With(
            slog.String("op", op),
            slog.String("request_id", middleware.GetReqID(r.Context())),
        )

        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            log.Error("empty Authorization header")
            render.Status(r, 401)
            render.JSON(w, r, "missing Authorization header")
            return
        }

        tkn := strings.Split(authHeader, "Bearer ")[TOKEN_INDEX_IN_HEADER]
        ctx, cancel := context.WithTimeout(context.Background(), timeout)
        defer cancel()

        err := signOuter.SignOut(ctx, tkn)
        if err != nil {
            log.Error(err.Error())
            render.Status(r, 401)
            render.JSON(w, r, err.Error())
            return
        }

        cookie := &http.Cookie{
            Name: "refreshToken",
            Value: "",
            Path: "/",
            MaxAge: EMPTY_COOKIE_MAX_AGE,
        }

        w.Header().Add("Set-Cookie", fmt.Sprintf("%s;Partitioned", cookie.String()))
        render.Status(r, 200)
    }
}

type UserActivator interface {
    ActivateUser(ctx context.Context, link string) (models.UserDTO, error)
}

func NewActivateUser(
    log *slog.Logger,
    userActivator UserActivator,
    timeout time.Duration,
) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        const op = "http.auth.NewActivateUser"

        log := log.With(
            slog.String("op", op),
            slog.String("request_id", middleware.GetReqID(r.Context())),
        )

        link := chi.URLParam(r, "link")
        log.Info("", slog.String("link", link))
        ctx, cancel := context.WithTimeout(context.Background(), timeout)
        defer cancel()

        usr, err := userActivator.ActivateUser(ctx, link)
        if err != nil {
            log.Error(err.Error())
            render.Status(r, 400)
            render.JSON(w, r, err.Error())
            return
        }

        render.Status(r, 200)
        render.JSON(w, r, usr)
    }
}

type TokenRefresher interface {
    Refresh(ctx context.Context, refreshToken string) (map[string]string, error)
}

func NewRefreshAccessToken(
    log *slog.Logger,
    tokenRefresher TokenRefresher,
    timeout time.Duration,
) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        const op = "http.auth.NewRefreshAccessToken"

        log := log.With(
            slog.String("op", op),
            slog.String("request_id", middleware.GetReqID(r.Context())),
        )

        refreshCookie, err := r.Cookie("refreshToken")
        if err != nil {
            log.Error(err.Error())
            render.Status(r, 403)
            render.JSON(w, r, err.Error())
            return
        }

        ctx, cancel := context.WithTimeout(context.Background(), timeout)
        defer cancel()

        tkns, err := tokenRefresher.Refresh(ctx, refreshCookie.Value)
        if err != nil {
            log.Error(err.Error())
            render.Status(r, 403)
            render.JSON(w, r, err.Error())
            return
        }

        cookie := &http.Cookie{
            Name:     "refreshToken",
            Value:    tkns["refreshToken"],
            HttpOnly: true,
            Secure:   true,
            Path:     "/",
            Expires:  time.Now().Add(30 * 24 * time.Hour),
        }

        w.Header().Add("Set-Cookie", fmt.Sprintf("%s;Partitioned", cookie.String()))
        w.Header().Set("Content-Type", "application/json")
        render.Status(r, 200)
    }
}
