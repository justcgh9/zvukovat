package repositories_test

import (
	"context"
	"justcgh9/spotify_clone/server/config"
	"justcgh9/spotify_clone/server/models"
	"justcgh9/spotify_clone/server/repositories"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
    "github.com/stretchr/testify/assert"
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

func TestUpdateFavourites()
