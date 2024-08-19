package repositories

import (
	"context"
	"justcgh9/spotify_clone/server/config"
	"justcgh9/spotify_clone/server/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var tokenCollection, trackCollection, commentCollection, albumCollection, userCollection *mongo.Collection

func Initialize(client *mongo.Client) {
	trackCollection = client.Database(config.DBName).Collection("tracks")
	commentCollection = client.Database(config.DBName).Collection("comments")
	albumCollection = client.Database(config.DBName).Collection("albums")
	userCollection = client.Database(config.DBName).Collection("users")
	tokenCollection = client.Database(config.DBName).Collection("tokens")
}

func GetAllTracks(params *models.TrackPaginationParams) ([]models.Track, error) {
	var tracks []models.Track

	findOptions := options.Find()
	if params != nil {
		findOptions.SetSkip(int64(params.Offset))
		findOptions.SetLimit(int64(params.Count))
	} else {
		findOptions.SetLimit(int64(10))
	}
	cursor, err := trackCollection.Find(context.TODO(), bson.D{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var track models.Track
		if err := cursor.Decode(&track); err != nil {
			return nil, err
		}
		tracks = append(tracks, track)
	}

	return tracks, nil
}

func SearchTrack(name string) ([]models.Track, error) {
	var tracks []models.Track

	filter := bson.D{
		{
			"name", bson.D{
				{"$regex", name},
				{"$options", "i"},
			},
		},
	}

	cursor, err := trackCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var track models.Track
		if err := cursor.Decode(&track); err != nil {
			return nil, err
		}
		tracks = append(tracks, track)
	}

	return tracks, nil
}

func GetOneTrack(id string) (models.Track, error) {
	var track models.Track
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Track{}, err
	}
	filter := bson.D{{"_id", objId}}
	err = trackCollection.FindOne(context.TODO(), filter).Decode(&track)
	if err != nil {
		return models.Track{}, err
	}
	return track, nil
}

func AddTrack(track models.Track) (models.Track, error) {
	track.Id = ""
	result, err := trackCollection.InsertOne(context.TODO(), track)
	if err != nil {
		return models.Track{}, err
	}
	track.Id = result.InsertedID.(primitive.ObjectID).Hex()
	return track, nil
}

func DeleteTrack(id string) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.D{{"_id", objId}}
	_, err = trackCollection.DeleteOne(context.TODO(), filter)

	if err != nil {
		return err
	}
	return nil
}

func UpdateTrack(track models.Track) (models.Track, error) {
	objId, err := primitive.ObjectIDFromHex(track.Id)
	if err != nil {
		return models.Track{}, err
	}

	filter := bson.D{{"_id", objId}}

	update := bson.D{
		{"$set", bson.D{
			{"name", track.Name},
			{"artist", track.Artist},
			{"text", track.Text},
			{"listens", track.Listens},
			{"picture", track.Picture},
			{"audio", track.Audio},
			{"comments", track.Comments},
		}},
	}

	_, err = trackCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return models.Track{}, err
	}
	return track, nil
}
