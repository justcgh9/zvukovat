package models

type Token struct {
    Id       string `bson:"_id,omitempty" json:"id"`
	UserId       string `bson:"user" json:"user_id"`
	RefreshToken string `bson:"refreshToken" json:"refreshToken"`
}

func NewToken(id, token string) Token {
    return Token{
        UserId: id,
        RefreshToken: token,
    }
}
