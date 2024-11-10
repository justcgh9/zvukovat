package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/justcgh9/zvukovat/services/users/internal/domain/models"
)


func GenerateTokens(payload models.UserDTO, accessSecret, refreshSecret string) (map[string]string, error) {

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


func ValidateAccessToken(tokenString, secret string) (*models.UserClaims, error) {
	secretKey := []byte(secret)

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

	return nil, fmt.Errorf("invalid token")
}

func ValidateRefreshToken(tokenString, secret string) (*models.UserClaims, error) {
	secretKey := []byte(secret)

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

	return nil, fmt.Errorf("invalid token")
}

func GetExpiryDate(tokenString string) (time.Time, error) {
    token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
    if err != nil {
        return time.Time{}, fmt.Errorf("failed to parse token: %w", err)
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok {
        if exp, ok := claims["exp"].(float64); ok {
            expiryDate := time.Unix(int64(exp), 0)
            return expiryDate, nil
        }
        return time.Time{}, fmt.Errorf("exp claim not found in token")
    }

    return time.Time{}, fmt.Errorf("invalid token claims")
}
