package template

import (
	"docker-init/internal/types"
	"os"
	"testing"
)

func TestGenerateGo(t *testing.T) {
	tmpDir := t.TempDir()
	os.Chdir(tmpDir)

	target := types.Target{
		TargetType: types.Go,
		Path:       tmpDir,
	}
	err := Generate(target)
	if err != nil {
		t.Fatalf("no error expected when initializing a Go project")
	}
}
