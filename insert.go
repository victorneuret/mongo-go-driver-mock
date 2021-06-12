package mongo_go_driver_mock

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func insert(userData user) (*user, error) {
	insertedResult, err := userCollection.InsertOne(context.Background(), userData)
	if err != nil {
		return nil, err
	}

	userData.ID = insertedResult.InsertedID.(primitive.ObjectID)
	return &userData, nil
}

func insertMany(usersData []user) error {
	users := make([]interface{}, len(usersData))
	for i, userData := range usersData {
		users[i] = userData
	}

	if _, err := userCollection.InsertMany(context.Background(), users); err != nil {
		return err
	}
	return nil
}
