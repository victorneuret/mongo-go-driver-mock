package mongo_go_driver_mock

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestFindOneAndUpdate(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		userCollection = mt.Coll
		userData := user{
			ID:    primitive.NewObjectID(),
			Name:  "john",
			Email: "john.doe@test.com",
		}
		mt.AddMockResponses(bson.D{
			{"ok", 1},
			{"value", bson.D{
				{"_id", userData.ID},
				{"name", userData.Name},
				{"email", userData.Email},
			}},
		})

		updatedUser, err := update(userData)

		assert.Nil(t, err)
		assert.Equal(t, &userData, updatedUser)
	})
}
