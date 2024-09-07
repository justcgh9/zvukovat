package services

import (
	"errors"
	"fmt"
	"justcgh9/spotify_clone/server/models"
	"justcgh9/spotify_clone/server/repositories"
	"net/http"
	"os"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func Registration(user models.User) (map[string]string, models.User, error) {

	_, err := repositories.GetUser(user.Email)

    if err == nil {
        return nil, models.User{}, errors.New("user with this email already exists")
    } else if !errors.Is(err, mongo.ErrNoDocuments) {
        fmt.Println(err.Error())
        return nil, models.User{}, err
	}

	password, err := hashPassword(user.Password)
	if err != nil {
		return nil, models.User{}, err
	}

	user.Password = password

	activationLink := uuid.New()
	user.ActivationLink = activationLink.String()
	createdUser, err := repositories.CreateUser(user)
	if err != nil {
		fmt.Println(createdUser)
		return nil, models.User{}, err
	}

	err = SendActivationMail(createdUser.Email, os.Getenv("API_URL")+"/activate/"+createdUser.ActivationLink)
	if err != nil {
		return nil, models.User{}, err
	}

	tokens, err := generateTokens(createdUser)
	if err != nil {
		return nil, models.User{}, err
	}

	userDTO := models.Token{
		RefreshToken: tokens["refreshToken"],
		UserId:       createdUser.Id,
	}

	_, err = repositories.SaveToken(userDTO)
	if err != nil {
		return nil, models.User{}, err
	}

	return tokens, createdUser, nil
}

func Login(user models.User) (map[string]string, models.User, error) {
    var tokens map[string]string

    fetchedUser, err := repositories.GetUser(user.Email)
    if err != nil {
        return tokens, models.User{}, err
    }

    if !checkPasswordHash(user.Password, fetchedUser.Password) {
        return tokens, models.User{}, errors.New("Incorrect password!")
    }

    if !fetchedUser.IsActivated {
        return tokens, models.User{}, errors.New("Account is not activated!")
    }
    tokens, err = generateTokens(fetchedUser)
	if err != nil {
		return tokens, models.User{}, err
	}

	userDTO := models.Token{
		RefreshToken: tokens["refreshToken"],
		UserId:       fetchedUser.Id,
	}

	_, err = repositories.SaveToken(userDTO)
	if err != nil {
		return tokens, models.User{}, err
	}


	return tokens, fetchedUser, nil
}

func Refresh(cookie http.Cookie) (map[string] string, error) {
    userData, err := ValidateRefreshToken(cookie.Value)
    if err != nil {
        return nil, err
    }

    err = repositories.FindToken(cookie.Value)
    if err != nil {
        return nil, err
    }

    user, err := repositories.GetUser(userData.Payload.Email)
    if err != nil {
        return nil, err
    }

    tokens, err := generateTokens(user)
	if err != nil {
		return tokens, err
	}

	userDTO := models.Token{
		RefreshToken: tokens["refreshToken"],
		UserId:       user.Id,
	}

	_, err = repositories.SaveToken(userDTO)
	if err != nil {
		return tokens, err
	}
    return tokens, nil
}

func Logout(userId string) error {
    return repositories.DeleteToken(userId)
}

func Refresh(cookie http.Cookie) (map[string] string, error) {
    userData, err := ValidateRefreshToken(cookie.Value)
    if err != nil {
        return nil, err
    }

    err = repositories.FindToken(cookie.Value)
    if err != nil {
        return nil, err
    }

    user, err := repositories.GetUser(userData.Payload.Email)
    if err != nil {
        return nil, err
    }

    tokens, err := generateTokens(user)
	if err != nil {
		return tokens, err
	}

	userDTO := models.Token{
		RefreshToken: tokens["refreshToken"],
		UserId:       user.Id,
	}

	_, err = repositories.SaveToken(userDTO)
	if err != nil {
		return tokens, err
	}
    return tokens, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}



