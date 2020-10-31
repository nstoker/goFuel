package localenv_test

import (
	"testing"

	"github.com/nstoker/gofuel/internal/pkg/localenv"
)

func Test_RootDir(t *testing.T) {
	p := localenv.RootDir()
	if p == "" {
		t.Fatalf("failed to get the root dir")
	}
}
