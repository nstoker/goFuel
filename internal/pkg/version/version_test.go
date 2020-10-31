package version_test

import (
	"testing"

	"github.com/nstoker/gofuel/internal/pkg/version"
)

func TestVersionExists(t *testing.T) {
	v := version.Version()

	if v == "" {
		t.Error("version string explicitly blank")
	}
}
