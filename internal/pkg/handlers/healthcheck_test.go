package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nstoker/gofuel/internal/pkg/handlers"
	h "github.com/nstoker/gofuel/internal/pkg/handlers"
	"github.com/nstoker/gofuel/internal/pkg/version"
)

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health-check", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.HealthCheckHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	result := &handlers.HealthCheck{}
	err = json.Unmarshal(rr.Body.Bytes(), &result)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if !result.Alive {
		t.Errorf("expected to get alive: true, instead got %v", result.Alive)
	}

	if result.Version != version.Version() {
		t.Errorf("expected version %s, got %s", version.Version(), result.Version)
	}
}
