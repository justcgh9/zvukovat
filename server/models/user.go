package models


type User struct {
    Id              string  `bson:"_id,omitempty" json:"id"`
    Email           string  `json:"email" bson:"email"`
    Password        string  `json:"password" bson:"password"`
    IsActivated     bool    `json:"isActivated" bson:"isActivated"`
    ActivationLink  string  `json:"activationLink" bson:"activationLink"`
}
