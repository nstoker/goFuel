package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/nstoker/gofuel/internal/pkg/version"
	"github.com/sirupsen/logrus"
)

// HealthCheck fields returned by the HealthCheckHandler
type HealthCheck struct {
	Alive   bool   `json:"alive"`
	Version string `json:"version"`
}

// HealthCheckHandler returns a status if the system is still running
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	appVersion := version.Version()
	hc := &HealthCheck{
		Alive:   true,
		Version: appVersion,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(hc)
	if err != nil {
		logrus.Printf("HealthCheckHandler: %s", err)
	}
}
