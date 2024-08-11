package repositories

import (
	"context"
	"justcgh9/spotify_clone/server/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateAlbum(album models.Album) (models.Album, error) {
    result, err := albumCollection.InsertOne(context.TODO(), album)
    if err != nil {
        return models.Album{}, err
    }
    album.Id = result.InsertedID.(primitive.ObjectID).Hex()
    return album, nil
}

func GetAlbum(albumID string) (models.Album, error) {
    var album models.Album
    objId, err := primitive.ObjectIDFromHex(albumID)
    if err != nil {
        return models.Album{}, err
    }
    filter := bson.D{{"_id", objId}}
    err = albumCollection.FindOne(context.TODO(), filter).Decode(&album)
    if err != nil {
        return models.Album{}, err
    }
    return album, nil
}

func AddTrackToAlbum(album models.Album, trackID string) (error) {

    objId, err := primitive.ObjectIDFromHex(album.Id)
    if err != nil {
        return err
    }

    filter := bson.D{{"_id", objId}}

    update := bson.D{{"$set", bson.D {{"tracks", append(album.Tracks, trackID)}}}}
    _, err = albumCollection.UpdateOne(context.TODO(), filter, update)
    if err != nil {
        return nil
    }

    return nil
}

func RemoveAlbum(albumID string) (error) {
    objId, err := primitive.ObjectIDFromHex(albumID)
    if err != nil {
        return err
    }

    filter := bson.D{{"_id", objId}}
    _, err = albumCollection.DeleteOne(context.TODO(), filter)

    if err != nil {
        return err
    }
    return nil
}
