package migrate

import (
	"testing"
)

func TestBadMigrationFile(t *testing.T) {
	oldMigrate := globalMigrate
	defer func() {
		globalMigrate = oldMigrate
	}()
	globalMigrate = NewMigrate(nil, nil)

	err := Register(migration{})
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
	MustRegister(migration{})
}
