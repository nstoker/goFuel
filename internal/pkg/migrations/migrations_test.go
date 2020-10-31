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

	err = m.MigrateUp(dsn)
	if err != nil {
		t.Fatal(err)
	}
}
