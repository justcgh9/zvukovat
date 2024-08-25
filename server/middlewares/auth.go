package middlewares

import (
	"context"
	"justcgh9/spotify_clone/server/services"
	"net/http"
	"strings"
)

func JwtAuthenticationMiddleware(next http.Handler) http.Handler{
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Missing Authorization Header", http.StatusUnauthorized)
            return
        }
        tokenString := strings.Split(authHeader, "Bearer ")[1]
        userData, err := services.ValidateAccessToken(tokenString)
        if err != nil {
            http.Error(w, err.Error(), http.StatusUnauthorized)
        }

        ctx := context.WithValue(r.Context(), "user", userData)
        r = r.WithContext(ctx)

        next.ServeHTTP(w, r)
    })
}
