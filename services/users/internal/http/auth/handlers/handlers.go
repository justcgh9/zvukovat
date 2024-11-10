package handlers

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/justcgh9/zvukovat/services/users/internal/domain/models"
	"github.com/justcgh9/zvukovat/services/users/internal/lib/jwt"
)

type Registrator interface {
    SignUp(ctx context.Context, usr models.User, domainName string) (map[string]string, models.UserDTO, error)
}

func NewSignUp(
        log *slog.Logger,
        registrator Registrator,
        timeout time.Duration,
        domainName string,
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
