package mongo_go_driver_mock

import "go.mongodb.org/mongo-driver/bson/primitive"

type user struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name"`
	Email string             `bson:"email"`
}
