package models

import "github.com/golang-jwt/jwt/v5"


type User struct {
    Id              string  `bson:"_id,omitempty" json:"id"`
    Email           string  `json:"email" bson:"email"`
    Password        string  `json:"password" bson:"password"`
    IsActivated     bool    `json:"isActivated" bson:"isActivated"`
    ActivationLink  string  `json:"activationLink" bson:"activationLink"`
}

type UserClaims struct {
    Payload User `json:"payload"`
    jwt.RegisteredClaims
}
