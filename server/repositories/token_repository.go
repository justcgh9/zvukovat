package repositories

import (
	"context"
	"errors"
	"justcgh9/spotify_clone/server/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func SaveToken(token models.Token) (models.Token,error) {
    var tokenData models.Token
    var err error
    filter := bson.D{{"user", token.UserId}}
    err = TokenCollection.FindOne(context.TODO(), filter).Decode(&tokenData)
    if err != nil {
        if !errors.Is(err, mongo.ErrNoDocuments){
            return models.Token{}, err
        }

        result, err := TokenCollection.InsertOne(context.TODO(), token)
        if err != nil {
            return models.Token{}, err
        }
        token.Id = result.InsertedID.(primitive.ObjectID).Hex()
        return token, nil


    }

    update := bson.D{
        {"$set", bson.D{
            {"user", token.UserId},
            {"refreshToken", token.RefreshToken},
        }},
    }
    _, err = TokenCollection.UpdateOne(context.TODO(), filter, update)
    if err != nil {
        return models.Token{}, err
    }


    return token, nil
}

func FindToken(tokenString string) (error) {
    var token models.Token
    filter := bson.M{"refreshToken": tokenString}
    err := TokenCollection.FindOne(context.TODO(), filter).Decode(&token)
    return err
}

func DeleteToken(userId string) error {
    filter := bson.M{"user": userId}
    _, err := TokenCollection.DeleteOne(context.TODO(), filter)
    return err
}
