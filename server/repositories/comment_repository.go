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

    objId, err := primitive.ObjectIDFromHex(comment.Id)
    if err != nil {
        return models.Comment{}, err
    }

    filter := bson.D{{"_id", objId}}

    update := bson.D{
        {"$set", bson.D{
            {"track_id", comment.Track_id},
            {"username", comment.Username},
            {"text", comment.Text},
        }},
    }

    _, err = commentCollection.UpdateOne(context.TODO(), filter, update)
    if err != nil {
        return models.Comment{}, err
    }
    return comment, nil

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

    return comment, nil
}

func DeleteComment(commentID string) (error) {
    objId, err := primitive.ObjectIDFromHex(commentID)
    if err != nil {
        return err
    }

    filter := bson.D{{"_id", objId}}
    _, err = commentCollection.DeleteOne(context.TODO(), filter)
    if err != nil {
        return err
    }
    return nil
}

