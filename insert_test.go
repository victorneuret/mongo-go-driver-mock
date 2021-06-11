package mongo_go_driver_mock

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestInsert(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		userCollection = mt.Coll
		id := primitive.NewObjectID()
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		mtest.CreateWriteErrorsResponse()

		insertedUser, err := insert(user{
			ID:    id,
			Name:  "john",
			Email: "john.doe@test.com",
		})

		assert.Nil(t, err)
		assert.Equal(t, &user{
			ID:    id,
			Name:  "john",
			Email: "john.doe@test.com",
		}, insertedUser)
	})
}
