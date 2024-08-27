package repositories

import (
	"context"
	"justcgh9/spotify_clone/server/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreatePlaylist(playlist models.Playlist) (models.Playlist, error) {
    result, err := playlistCollection.InsertOne(context.TODO(), playlist)
    if err != nil {
        return models.Playlist{}, err
    }
    playlist.Id = result.InsertedID.(primitive.ObjectID).Hex()
    return playlist, nil
}

func GetPlaylist(playlistID string) (models.Playlist, error) {
    var playlist models.Playlist
    objId, err := primitive.ObjectIDFromHex(playlistID)
    if err != nil {
        return models.Playlist{}, err
    }
    filter := bson.D{{"_id", objId}}
    err = playlistCollection.FindOne(context.TODO(), filter).Decode(&playlist)
    if err != nil {
        return models.Playlist{}, err
    }
    return playlist, nil
}

func GetMyPlaylists(userId string) ([]models.Playlist, error) {

    filter := bson.M{"owner": userId}
	findOptions := options.Find()
    findOptions.SetLimit(10)

    var playlists []models.Playlist
	cursor, err := playlistCollection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var playlist models.Playlist
		if err := cursor.Decode(&playlist); err != nil {
			return nil, err
		}
		playlists = append(playlists, playlist)
	}

	return playlists, nil
}

func AddTrackToPlaylist(playlist models.Playlist, trackID string) (error) {

    objId, err := primitive.ObjectIDFromHex(playlist.Id)
    if err != nil {
        return err
    }

    filter := bson.D{{"_id", objId}}

    update := bson.D{{"$set", bson.D {{"tracks", append(playlist.Tracks, trackID)}}}}
    _, err = playlistCollection.UpdateOne(context.TODO(), filter, update)
    if err != nil {
        return err
    }

    return nil
}

func RemovePlaylist(playlistID string) (error) {
    objId, err := primitive.ObjectIDFromHex(playlistID)
    if err != nil {
        return err
    }

    filter := bson.D{{"_id", objId}}
    _, err = playlistCollection.DeleteOne(context.TODO(), filter)

    if err != nil {
        return err
    }
    return nil
}

func SetPlaylistVisibility(playlistId string, visibility bool) error {

    objId, err := primitive.ObjectIDFromHex(playlistId)
    if err != nil {
        return err
    }

    filter := bson.D{{"_id", objId}}

    update := bson.M{"$set": bson.M{
        "isPrivate": visibility,
    }}
    _, err = playlistCollection.UpdateOne(context.TODO(), filter, update)
    if err != nil {
        return err
    }

    return nil
}
