package repositories

import (
	"context"
	"justcgh9/spotify_clone/server/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateComment(comment models.Comment) (models.Comment, error) {

    comment.Id = ""
    result, err := commentCollection.InsertOne(context.TODO(), comment)
    if err != nil {
        return models.Comment{}, err
    }

    comment.Id = result.InsertedID.(primitive.ObjectID).Hex()
    return comment, nil

}

func EditComment(comment models.Comment) (models.Comment, error) {

    return models.Comment{}, nil
}

func GetComment(commentID string) (models.Comment, error) {

    var comment models.Comment
    objId, err := primitive.ObjectIDFromHex(commentID)
    if err != nil {
        return models.Comment{}, err
    }
    filter := bson.D{{"_id", objId}}
    err = commentCollection.FindOne(context.TODO(), filter).Decode(&comment)
    if err != nil {
        return models.Comment{}, err
    }

    return models.Comment{}, nil
}

func DeleteComment(commentID string) (models.Comment, error) {
    return models.Comment{}, nil
}

