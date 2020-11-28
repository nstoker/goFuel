package database_test

import (
	"testing"

	"github.com/joho/godotenv"
	"github.com/nstoker/gofuel/internal/pkg/database"
)

const testEnv = "../../../.test.env"

func TestFailInitialiseDatabase(t *testing.T) {
	err := database.InitialiseDatabase()
	if err.Error() != "database url environment variable missing" {
		t.Fatal("expected error message")
	}
}

func TestInitialiseDatabase(t *testing.T) {
	godotenv.Load(testEnv)

	err := database.InitialiseDatabase()
	if err != nil {
		t.Fatal(err)
	}
}
