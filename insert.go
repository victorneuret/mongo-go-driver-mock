package mongo_go_driver_mock

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func insert(userData user) (*user, error) {
	insertedResult, err := userCollection.InsertOne(context.Background(), userData)
	if err != nil {
		return nil, err
	}

	fmt.Println("ID", insertedResult.InsertedID)
	userData.ID = insertedResult.InsertedID.(primitive.ObjectID)
	return &userData, nil
}
