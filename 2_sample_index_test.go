// +build integration

package migrate

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const globalTestIndexName = "test_idx_2"

func init() {
	Register(migration{})
}

type migration struct{}

func (migration) Up(client *mongo.Client, db *mongo.Database) error {
	keys := bson.D{{"a", 1}}
	opt := options.Index().SetName(globalTestIndexName)
	model := mongo.IndexModel{Keys: keys, Options: opt}
	_, err := db.Collection(globalTestCollection).Indexes().CreateOne(context.TODO(), model)
	if err != nil {
		return err
	}
	return nil
}

func (migration) Down(client *mongo.Client, db *mongo.Database) error {
	_, err := db.Collection(globalTestCollection).Indexes().DropOne(context.TODO(), globalTestIndexName)
	if err != nil {
		return err
	}
	return nil
}