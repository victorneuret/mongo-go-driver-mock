package mongo_go_driver_mock

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getFromID(id primitive.ObjectID) (*user, error) {
	filter := bson.D{{Key: "id", Value: id}}
	var object user

	if err := userCollection.FindOne(context.Background(), filter).Decode(&object); err != nil {
		return nil, err
	}
	return &object, nil
}

func find(name string) ([]user, error) {
	filter := bson.D{{Key: "name", Value: name}}
	var users []user

	cursor, err := userCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.Background(), &users); err != nil {
		return nil, err
	}
	return users, nil
}
