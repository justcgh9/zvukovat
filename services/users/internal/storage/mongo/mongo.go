package mongo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/justcgh9/zvukovat/services/users/internal/domain/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
    client *mongo.Client
    userCollection *mongo.Collection
    tokenCollection *mongo.Collection
}

func New(ctx context.Context, uri, dbName string) (*Storage, error) {
    const op = "storage.mongo.New"

    clientOptions := options.Client().ApplyURI(uri)
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        return nil, fmt.Errorf("%s: %w", op, err)
    }

    pingCtx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
    defer cancel()
    err = client.Ping(pingCtx, nil)
    if err != nil {
        return nil, fmt.Errorf("%s: %w", op, err)
    }

    return &Storage{
        client: client,
        userCollection: client.Database(dbName).Collection("users"),
        tokenCollection: client.Database(dbName).Collection("tokens"),
    }, nil
}

func (s *Storage) SaveUser(ctx context.Context, usr models.User) (models.UserDTO, error) {
    const op = "storage.mongo.SaveUser"

    result, err := s.userCollection.InsertOne(ctx, usr)
    if err != nil {
        return models.UserDTO{}, fmt.Errorf("%s: %w", op, err)
    }

    return models.UserDTO{
        Id: result.InsertedID.(primitive.ObjectID).Hex(),
        Email: usr.Email,
        Username: usr.Username,
        IsActivated: false,
        FavouriteTracks: make([]string, 0),
    }, nil

}

func (s *Storage) User(ctx context.Context, email string) (models.User, error) {
    const op = "storage.mongo.User"
    var usr models.User

    filter := bson.M{"userdto.email": email}
    err := s.userCollection.FindOne(ctx, filter).Decode(&usr)
    if err != nil {
        return models.User{}, fmt.Errorf("%s: %w", op, err)
    }

    return usr, nil
}

func (s *Storage) UpdateFavourites(ctx context.Context, user models.UserDTO) error {
    const op = "storage.mongo.UpdateFavourites"
    filter := bson.M{"userdto.email": user.Email}
    update := bson.M{"$set": bson.M{
        "userdto.favouriteTracks": user.FavouriteTracks,
    }}

    _, err := s.userCollection.UpdateOne(ctx, filter, update)
    if err != nil {
        return fmt.Errorf("%s: %w", op, err)
    }


    return nil
}

func (s *Storage) ActivateUser(ctx context.Context, link string) (models.UserDTO, error) {
    const op = "storage.mongo.ActivateUser"
    var usr models.UserDTO

    filter := bson.M{"activationLink": link}
    err := s.userCollection.FindOne(ctx, filter).Decode(&usr)
    if err != nil {
        return models.UserDTO{}, fmt.Errorf("%s: %w", op, err)
    }

    update := bson.M{
        "$set": bson.M{
            "userdto.isActivated": true,
        },
    }

    _, err = s.userCollection.UpdateOne(ctx, filter, update)
    if err != nil {
        return models.UserDTO{}, fmt.Errorf("%s: %w", op, err)
    }

    usr.IsActivated = true

    return usr, nil
}

func (s *Storage) SaveToken(ctx context.Context, tkn models.Token) (models.Token, error) {
    const op = "storage.mongo.SaveToken"
    var oldTkn models.Token
    var err error

    filter := bson.M{"user": tkn.UserId}
    err = s.tokenCollection.FindOne(ctx, filter).Decode(&oldTkn)
    if err != nil {
        if !errors.Is(err, mongo.ErrNoDocuments) {
            return models.Token{}, fmt.Errorf("%s: %w", op, err)
        }

        res, err := s.tokenCollection.InsertOne(ctx, tkn)
        if err != nil {
            return models.Token{}, fmt.Errorf("%s: %w", op, err)
        }

        tkn.Id = res.InsertedID.(primitive.ObjectID).Hex()
        return tkn, nil
    }

    update := bson.M{
        "$set": bson.M{
            "user": tkn.UserId,
            "refreshToken": tkn.RefreshToken,
        },
    }
    _, err = s.tokenCollection.UpdateOne(ctx, filter, update)
    if err != nil {
        return models.Token{}, fmt.Errorf("%s: %w", op, err)
    }

    return tkn, nil
}

func (s *Storage) Token(ctx context.Context, tknStr string) (models.Token, error) {
    const op = "storage.mongo.Token"

    var token models.Token
    filter := bson.M{"refreshToken": tknStr}
    err := s.tokenCollection.FindOne(ctx, filter).Decode(&token)
    if err != nil {
        return models.Token{}, fmt.Errorf("%s: %w", op, err)
    }

    return token, nil
}

func (s *Storage) DeleteToken(ctx context.Context, usrId string) error {
    const op = "storage.mongo.DeleteToken"

    filter := bson.M{"user": usrId}
    _, err := s.tokenCollection.DeleteOne(ctx, filter)
    if err != nil {
        return fmt.Errorf("%s: %w", op, err)
    }

    return nil
}
