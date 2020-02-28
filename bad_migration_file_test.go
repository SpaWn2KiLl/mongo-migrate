package migrate

import (
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

func TestBadMigrationFile(t *testing.T) {
	oldMigrate := globalMigrate
	defer func() {
		globalMigrate = oldMigrate
	}()
	globalMigrate = NewMigrate(nil, nil)

	err := Register(func(client *mongo.Client, db *mongo.Database) error {
		return nil
	}, func(client *mongo.Client, db *mongo.Database) error {
		return nil
	})
	if err == nil {
		t.Errorf("Unexpected nil error")
	}
}

func TestBadMigrationFilePanic(t *testing.T) {
	oldMigrate := globalMigrate
	defer func() {
		globalMigrate = oldMigrate
		if r := recover(); r == nil {
			t.Errorf("Unexpectedly no panic recovered")
		}
	}()
	globalMigrate = NewMigrate(nil, nil)
	MustRegister(func(client *mongo.Client, db *mongo.Database) error {
		return nil
	}, func(client *mongo.Client, db *mongo.Database) error {
		return nil
	})
}
