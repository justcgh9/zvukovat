package models

type Token struct {
    UserId       string   `bson:"user_id,omitempty" json:"user_id"`
    RefreshToken string   `bson:"refreshToken" json:"refreshToken"`
}
