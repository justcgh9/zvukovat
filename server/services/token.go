package services

import (
	"justcgh9/spotify_clone/server/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func generateTokens(payload models.User) (map[string]string, error) {
    accessSecret := os.Getenv("JWT_ACCESS_SECRET")
    refreshSecret := os.Getenv("JWT_REFRESH_SECRET")

    accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "exp":      time.Now().Add(15 * time.Minute).Unix(),
        "payload":  payload,
    }).SignedString([]byte(accessSecret))
    if err != nil {
        return nil, err
    }

    refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "exp":      time.Now().Add(60 * 24 * time.Hour).Unix(),
        "payload":  payload,
    }).SignedString([]byte(refreshSecret))
    if err != nil {
        return nil, err
    }

    return map[string]string{
        "accessToken":  accessToken,
        "refreshToken": refreshToken,
    }, nil
}
