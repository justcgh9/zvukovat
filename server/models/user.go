package models

import "github.com/golang-jwt/jwt/v5"

type User struct {
	Id              string   `bson:"_id,omitempty" json:"id"`
    Username        string   `json:"username" bson:"username"`
	Email           string   `json:"email" bson:"email"`
	Password        string   `json:"password" bson:"password"`
	IsActivated     bool     `json:"isActivated" bson:"isActivated"`
	ActivationLink  string   `json:"activationLink" bson:"activationLink"`
	FavouriteTracks []string `json:"favouriteTracks" bson:"favouriteTracks"`
}

type UserClaims struct {
	Payload User `json:"payload"`
	jwt.RegisteredClaims
}

type UserDTO struct {
	Id              string   `bson:"_id,omitempty" json:"id"`
    Username        string   `json:"username" bson:"username"`
	Email           string   `json:"email" bson:"email"`
	IsActivated     bool     `json:"isActivated" bson:"isActivated"`
	FavouriteTracks []string `json:"favouriteTracks" bson:"favouriteTracks"`
}

func (user User) ContainsTrack(trackId string) bool {
	if len(user.FavouriteTracks) <= 0 {
		return false
	}
	for _, track := range user.FavouriteTracks {
		if track == trackId {
			return true
		}
	}

	return false
}
