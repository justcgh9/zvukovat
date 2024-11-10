package models

import "github.com/golang-jwt/jwt/v5"


type UserDTO struct {
	Id              string   `bson:"_id,omitempty" json:"id"`
    Username        string   `json:"username" bson:"username"`
	Email           string   `json:"email" bson:"email"`
	IsActivated     bool     `json:"isActivated" bson:"isActivated"`
	FavouriteTracks []string `json:"favouriteTracks" bson:"favouriteTracks"`
}

type User struct {
	Password        string   `json:"password" bson:"password"`
	ActivationLink  string   `json:"activationLink" bson:"activationLink"`
    UserDTO
}

type UserClaims struct {
	Payload User `json:"payload"`
	jwt.RegisteredClaims
}

func (u UserDTO) ContainsTrack(id string) bool {
	if len(u.FavouriteTracks) <= 0 {
		return false
	}
	for _, track := range u.FavouriteTracks {
		if track == id {
			return true
		}
	}

	return false
}
