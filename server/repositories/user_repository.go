package repositories

import (
	"context"
	"fmt"
	"justcgh9/spotify_clone/server/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUser(email string) (models.User, error) {

	var user models.User

	filter := bson.D{{"email", email}}

	err := UserCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func CreateUser(CreateUserDTO models.User) (models.User, error) {

	var user models.User

	result, err := UserCollection.InsertOne(context.TODO(), CreateUserDTO)
	if err != nil {
		return models.User{}, err
	}
	user = models.User{
		Id:             result.InsertedID.(primitive.ObjectID).Hex(),
		Email:          CreateUserDTO.Email,
		Password:       CreateUserDTO.Password,
        Username:       CreateUserDTO.Username,
		ActivationLink: CreateUserDTO.ActivationLink,
        FavouriteTracks: make([]string, 0),
	}

	return user, nil
}

func ActivateUser(link string) (models.User, error) {

	filter := bson.M{"activationLink": link}

	var user models.User

	err := UserCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		fmt.Println(user)
		return models.User{}, err
	}

	update := bson.D{
		{"$set", bson.D{
			{"isActivated", true},
		}},
	}

	_, err = UserCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return models.User{}, err
	}

    user.IsActivated = true

	return user, nil
}

func GetAllUsers() ([]models.User, error) {

    users := make([]models.User, 0)

	cursor, err := UserCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func UpdateFavourites(user models.User) error {
    filter := bson.M{"email": user.Email}
    update := bson.M{"$set": bson.M{
        "favouriteTracks": user.FavouriteTracks,
    }}
    _, err := UserCollection.UpdateOne(context.TODO(), filter, update)
    return err
}
