package mongo_go_driver_mock

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestFind(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		userCollection = mt.Coll
		expectedUser := user{
			ID:    primitive.NewObjectID(),
			Name:  "john",
			Email: "john.doe@test.com",
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{"_id", expectedUser.ID},
			{"name", expectedUser.Name},
			{"email", expectedUser.Email},
		}))
		userResponse, err := getFromID(expectedUser.ID)
		assert.Nil(t, err)
		assert.Equal(t, &expectedUser, userResponse)
	})
}
