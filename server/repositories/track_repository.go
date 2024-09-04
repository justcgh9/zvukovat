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


var PlaylistCollection, TokenCollection, TrackCollection, AlbumCollection, UserCollection *mongo.Collection

func Initialize(client *mongo.Client) {
	TrackCollection = client.Database(config.DBName).Collection("tracks")
	AlbumCollection = client.Database(config.DBName).Collection("albums")
	UserCollection = client.Database(config.DBName).Collection("users")
	TokenCollection = client.Database(config.DBName).Collection("tokens")
	PlaylistCollection = client.Database(config.DBName).Collection("playlists")

}

func GetAllTracks(params *models.TrackPaginationParams) ([]models.Track, error) {

    tracks := make([]models.Track, 0)
	findOptions := options.Find()
	if params != nil {
		findOptions.SetSkip(int64(params.Offset))
		findOptions.SetLimit(int64(params.Count))
	} else {
		findOptions.SetLimit(int64(10))
	}
	cursor, err := TrackCollection.Find(context.TODO(), bson.D{}, findOptions)
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

func SearchTrack(name, artist string) ([]models.Track, error) {
    tracks := make([]models.Track, 0)

	filter := bson.D{
		{
			"name", bson.D{
				{"$regex", name},
				{"$options", "i"},
			},
		},
		{
			"artist", bson.D{
				{"$regex", artist},
				{"$options", "i"},
			},
		},
	}

	cursor, err := TrackCollection.Find(context.TODO(), filter)
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
	err = TrackCollection.FindOne(context.TODO(), filter).Decode(&track)
	if err != nil {
		return models.Track{}, err
	}
	return track, nil
}

func AddTrack(track models.Track) (models.Track, error) {
	track.Id = ""
	result, err := TrackCollection.InsertOne(context.TODO(), track)
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
	_, err = TrackCollection.DeleteOne(context.TODO(), filter)

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
		}},
	}

	_, err = TrackCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return models.Track{}, err
	}
	return track, nil
}

func GetArtists() ([]string, error) {

    pipeline := mongo.Pipeline{
        bson.D{{"$group", bson.D{
            {"_id", bson.D{{"$toLower", "$artist"}}},
        }}},
    }


    cursor, err := TrackCollection.Aggregate(context.TODO(), pipeline)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.TODO())

    var artists []string
    for cursor.Next(context.TODO()) {
        var result struct {
            Id string `bson:"_id"`
        }
        if err := cursor.Decode(&result); err != nil {
            return nil, err
        }
        artists = append(artists, result.Id)
    }
    return artists, nil
}
