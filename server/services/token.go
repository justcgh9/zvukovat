package services

import (
	"justcgh9/spotify_clone/server/models"
	"os"
	"time"
    "fmt"
    "errors"
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


func ValidateAccessToken(tokenString string) (*models.UserClaims, error) {
	secretKey := []byte(os.Getenv("JWT_ACCESS_SECRET"))

	token, err := jwt.ParseWithClaims(tokenString, &models.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok && token.Method.Alg() == jwt.SigningMethodHS256.Alg() {
			return secretKey, nil
		}
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*models.UserClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func ValidateRefreshToken(tokenString string) (*models.UserClaims, error) {
	secretKey := []byte(os.Getenv("JWT_REFRESH_SECRET"))

	token, err := jwt.ParseWithClaims(tokenString, &models.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok && token.Method.Alg() == jwt.SigningMethodHS256.Alg() {
			return secretKey, nil
		}
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*models.UserClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
