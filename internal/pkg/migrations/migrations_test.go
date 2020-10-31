package migrations_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	m "github.com/nstoker/gofuel/internal/pkg/migrations"
)

const testEnv = "../../../.test.env"

func TestDownUpMigrate(t *testing.T) {
	godotenv.Load(testEnv)

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		t.Fatalf("did not get DATABASE_URL from environment")
	}

	err := m.MigrateDown(dsn)
	if err != nil {
		t.Fatal(err)
	}

	err = m.MigrateDown(dsn)
	if err != nil {
		t.Fatalf("expected no error on multiple down migrations: %s", err)
	}

	err = m.MigrateUp(dsn)
	if err != nil {
		t.Fatal(err)
	}

	err = m.MigrateUp(dsn)
	if err != nil {
		t.Fatalf("expected no error on multiple up migrations %s", err)
	}
}

func TestMigrateWithoutDSN(t *testing.T) {
	dsn := ""
	err := m.MigrateDown(dsn)
	if err == nil {
		t.Errorf("failed to detect an empty dsn")
	}
	err = m.MigrateUp(dsn)
	if err == nil {
		t.Errorf("failed to detect an empty up dsn")
	}
}
