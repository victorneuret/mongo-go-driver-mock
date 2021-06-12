package mongo_go_driver_mock

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func update(userData user) (*user, error) {
	if err := userCollection.FindOneAndUpdate(
		context.Background(),
		bson.D{
			{"_id", userData.ID},
		},
		bson.D{{"$set", userData}},
		options.FindOneAndUpdate().SetReturnDocument(1),
	).Decode(&userData); err != nil {
		return nil, err
	}
	return &userData, nil
}
