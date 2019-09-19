package sequence

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoNextId struct {
	id    string `bson:"id"`
	Value int    `bson:"value"`
}

// GetNextID return auto inc id
func GetNextID(collection *mongo.Collection, connectString string) (int, error) {
	var nextId MongoNextId
	ops := options.FindOneAndUpdate()
	ops.SetUpsert(true)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection.FindOneAndUpdate(
		ctx,
		bson.M{"id": connectString},
		bson.M{"$inc": bson.M{"value": 1}},
		ops,
	).Decode(&nextId)

	return nextId.Value, nil
}
