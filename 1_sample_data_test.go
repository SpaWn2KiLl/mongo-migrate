// +build integration

package migrate

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const globalTestCollection = "test-global"

func init() {
	Register(migration{})
}

type migration struct {
}

func (migration) Up(client *mongo.Client, db *mongo.Database) error {
	_, err := db.Collection(globalTestCollection).InsertOne(context.TODO(), bson.D{{"a", "b"}})
	if err != nil {
		return err
	}
	return nil
}

func (migration) Down(client *mongo.Client, db *mongo.Database) error {
	_, err := db.Collection(globalTestCollection).DeleteOne(context.TODO(), bson.D{{"a", "b"}})
	if err != nil {
		return err
	}
	return nil
}
