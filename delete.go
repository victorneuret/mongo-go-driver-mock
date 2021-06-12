package mongo_go_driver_mock

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func delete(id primitive.ObjectID) error {
	result, err := userCollection.DeleteOne(
		context.Background(),
		bson.D{
			{Key: "_id", Value: id},
		},
	)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("no document deleted")
	}
	return nil
}
