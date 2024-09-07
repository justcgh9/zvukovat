package repositories_test

import (
	"context"
	"justcgh9/spotify_clone/server/config"
	"justcgh9/spotify_clone/server/models"
	"justcgh9/spotify_clone/server/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectToMongoDB(t *testing.T) *mongo.Client {

	clientOptions := options.Client().ApplyURI(config.MongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
        t.Fatalf("Failed to connect to MongoDB: %v", err)
	}

    if err := client.Ping(context.TODO(), nil); err != nil{
        t.Fatalf("Failed to ping MongoDB: %v", err)
    }

    return client
}

func setupTestRepo(t *testing.T) (*mongo.Client, map[string]*mongo.Collection) {
    client := connectToMongoDB(t)
    collectionNames := []string{"users", "albums", "tokens", "playlists", "tracks"}
    collections := make(map[string]*mongo.Collection, 5)

    for _, name := range collectionNames {
        collection := client.Database("testdb").Collection(name)

        if err := collection.Drop(context.TODO()); err != nil {
            t.Fatalf("Failed to deop collection %s: %v", name, err)
        }

        collections[name] = collection
    }

    return client, collections
}

func TestInsertUser(t *testing.T) {

    client, collections := setupTestRepo(t)
    repositories.UserCollection = collections["users"]
    user := models.User{
		Email:      "john.doe@example.com",
        Password:   "12345678",
        Username:   "exampleuser",
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			t.Fatalf("Failed to disconnect MongoDB client: %v", err)
		}
	}()

    createdUser, err := repositories.CreateUser(user)
    assert.NoError(t, err)
    assert.Equal(t, user.Username, createdUser.Username)
    assert.Equal(t, user.Password, createdUser.Password)
    assert.Equal(t, user.Email, createdUser.Email)

}

func TestGetUser(t *testing.T) {

    client, collections := setupTestRepo(t)
    repositories.UserCollection = collections["users"]
    user := models.User{
		Email:      "john.doe@example.com",
        Password:   "12345678",
        Username:   "exampleuser",
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			t.Fatalf("Failed to disconnect MongoDB client: %v", err)
		}
	}()

    _, err := repositories.CreateUser(user)
    assert.NoError(t, err)

    recievedUser, err := repositories.GetUser(user.Email)
    assert.NoError(t, err)

    assert.Equal(t, user.Username, recievedUser.Username)
    assert.Equal(t, user.Password, recievedUser.Password)
    assert.Equal(t, user.Email, recievedUser.Email)
}

func TestActivateUser(t *testing.T) {

    client, collections := setupTestRepo(t)
    repositories.UserCollection = collections["users"]
    user := models.User{
		Email:      "john.doe@example.com",
        Password:   "12345678",
        Username:   "exampleuser",
        ActivationLink: "example.com/activate",
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			t.Fatalf("Failed to disconnect MongoDB client: %v", err)
		}
	}()

    createdUser, err := repositories.CreateUser(user)
    assert.NoError(t, err)
    assert.Equal(t, createdUser.ActivationLink, user.ActivationLink)
    activatedUser, err := repositories.ActivateUser(user.ActivationLink)
    assert.NoError(t, err)
    assert.Equal(t, activatedUser.IsActivated, true)
}

func TestAddTrack(t *testing.T) {
    client, collections := setupTestRepo(t)
    repositories.TrackCollection = collections["tracks"]

    track := models.Track{
        Name: "some name",
        Artist: "some Artist",
        Text: "some Lyrics",
        Picture: "/path/to/picure",
        Audio: "/path/to/audio",
    }

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			t.Fatalf("Failed to disconnect MongoDB client: %v", err)
		}
	}()

    createdTrack, err := repositories.AddTrack(track)
    assert.NoError(t, err)
    assert.Equal(t, track.Name, createdTrack.Name)
    assert.Equal(t, track.Artist, createdTrack.Artist)
    assert.Equal(t, track.Text, createdTrack.Text)
    assert.Equal(t, track.Picture, createdTrack.Picture)
    assert.Equal(t, track.Audio, createdTrack.Audio)
}

func TestGetOneTrack(t *testing.T) {

    client, collections := setupTestRepo(t)
    repositories.TrackCollection = collections["tracks"]

    track := models.Track{
        Name: "some name",
        Artist: "some Artist",
        Text: "some Lyrics",
        Picture: "/path/to/picure",
        Audio: "/path/to/audio",
    }

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			t.Fatalf("Failed to disconnect MongoDB client: %v", err)
		}
	}()

    createdTrack, err := repositories.AddTrack(track)
    assert.NoError(t, err)

    fetchedTrack, err := repositories.GetOneTrack(createdTrack.Id)
    assert.NoError(t, err)
    assert.Equal(t, track.Name, fetchedTrack.Name)
    assert.Equal(t, track.Artist, fetchedTrack.Artist)
    assert.Equal(t, track.Text, fetchedTrack.Text)
    assert.Equal(t, track.Picture, fetchedTrack.Picture)
    assert.Equal(t, track.Audio, fetchedTrack.Audio)
}


func TestGetAllTracks(t *testing.T) {

    client, collections := setupTestRepo(t)
    repositories.TrackCollection = collections["tracks"]

    track := models.Track{
        Name: "some name",
        Artist: "some Artist",
        Text: "some Lyrics",
        Picture: "/path/to/picure",
        Audio: "/path/to/audio",
    }

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			t.Fatalf("Failed to disconnect MongoDB client: %v", err)
		}
	}()

    _, err := repositories.AddTrack(track)
    assert.NoError(t, err)

    _, err = repositories.AddTrack(track)
    assert.NoError(t, err)
    fetchedTracks, err := repositories.GetAllTracks(nil)
    assert.NoError(t, err)
    assert.Equal(t, len(fetchedTracks), 2)
    for _, fetchedTrack := range fetchedTracks {
        assert.Equal(t, track.Name, fetchedTrack.Name)
        assert.Equal(t, track.Artist, fetchedTrack.Artist)
        assert.Equal(t, track.Text, fetchedTrack.Text)
        assert.Equal(t, track.Picture, fetchedTrack.Picture)
        assert.Equal(t, track.Audio, fetchedTrack.Audio)
    }
}

func TestSearchTrack(t *testing.T) {

    client, collections := setupTestRepo(t)
    repositories.TrackCollection = collections["tracks"]

    track := models.Track{
        Name: "some name",
        Artist: "some Artist",
        Text: "some Lyrics",
        Picture: "/path/to/picure",
        Audio: "/path/to/audio",
    }
    anotherTrack := models.Track{
        Name: "another name",
        Artist: "another Artist",
        Text: "some Lyrics",
        Picture: "/path/to/picure",
        Audio: "/path/to/audio",
    }

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			t.Fatalf("Failed to disconnect MongoDB client: %v", err)
		}
	}()

    _, err := repositories.AddTrack(track)
    assert.NoError(t, err)

    _, err = repositories.AddTrack(anotherTrack)
    assert.NoError(t, err)

    fetchedTracks, err := repositories.SearchTrack("sOme", "somE")
    assert.NoError(t, err)
    assert.Equal(t, len(fetchedTracks), 1)
    fetchedTrack := fetchedTracks[0]
    assert.Equal(t, track.Name, fetchedTrack.Name)
    assert.Equal(t, track.Artist, fetchedTrack.Artist)
    assert.Equal(t, track.Text, fetchedTrack.Text)
    assert.Equal(t, track.Picture, fetchedTrack.Picture)
    assert.Equal(t, track.Audio, fetchedTrack.Audio)
}


func TestUpdateTrack(t *testing.T) {
    client, collections := setupTestRepo(t)
    repositories.TrackCollection = collections["tracks"]

    track := models.Track{
        Name: "some name",
        Artist: "some Artist",
        Text: "some Lyrics",
        Picture: "/path/to/picure",
        Audio: "/path/to/audio",
    }

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			t.Fatalf("Failed to disconnect MongoDB client: %v", err)
		}
	}()

    createdTrack, err := repositories.AddTrack(track)
    assert.NoError(t, err)

    track.Id = createdTrack.Id
    track.Name = "another name"

    updatedTrack, err := repositories.UpdateTrack(track)
    assert.NoError(t, err)
    assert.Equal(t, createdTrack.Id, updatedTrack.Id)
    assert.Equal(t, track.Name, updatedTrack.Name)
    assert.Equal(t, track.Artist, updatedTrack.Artist)
    assert.Equal(t, track.Text, updatedTrack.Text)
    assert.Equal(t, track.Picture, updatedTrack.Picture)
    assert.Equal(t, track.Audio, updatedTrack.Audio)
}

func TestGetArtists(t *testing.T) {

    client, collections := setupTestRepo(t)
    repositories.TrackCollection = collections["tracks"]

    track := models.Track{
        Name: "some name",
        Artist: "some Artist",
        Text: "some Lyrics",
        Picture: "/path/to/picure",
        Audio: "/path/to/audio",
    }

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			t.Fatalf("Failed to disconnect MongoDB client: %v", err)
		}
	}()
    _, err := repositories.AddTrack(track)
    assert.NoError(t, err)

    track.Artist = "sOmE ArTiSt"
    _, err = repositories.AddTrack(track)
    assert.NoError(t, err)

    track.Artist = "Another Artist"
    _, err = repositories.AddTrack(track)
    assert.NoError(t, err)

    artists, err := repositories.GetArtists()
    assert.NoError(t, err)
    expected_artists := []string{"some artist", "another artist"}
    assert.ObjectsAreEqualValues(artists, expected_artists)
}

func TestUpdateFavourites(t *testing.T) {

    client, collections := setupTestRepo(t)
    repositories.UserCollection = collections["users"]
    repositories.TrackCollection = collections["tracks"]
/*    user := models.User{
		Email:      "john.doe@example.com",
        Password:   "12345678",
        Username:   "exampleuser",
        ActivationLink: "example.com/activate",
	}
*/
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			t.Fatalf("Failed to disconnect MongoDB client: %v", err)
		}
	}()
}
