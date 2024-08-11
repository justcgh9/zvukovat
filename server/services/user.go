package services

import (
	"errors"
	"justcgh9/spotify_clone/server/models"
	"justcgh9/spotify_clone/server/repositories"

	"go.mongodb.org/mongo-driver/mongo"
)

func Registration(user models.User) (error) {
    _, err := repositories.GetUser(user.Email)
    if !errors.Is(err, mongo.ErrNoDocuments){
       return errors.New("User with this email already exists")
    }

    return nil
}
