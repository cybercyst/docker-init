package template

import (
	"docker-init/internal/types"
	"os"
	"testing"
)

func TestGenerateGo(t *testing.T) {
	tmpDir := t.TempDir()
	os.Chdir("../..")

	target := types.Target{
		TargetType: types.Go,
		Path:       tmpDir,
	}

	input := map[string]interface{}{
		"binary": "go-api",
	}

	err := Generate(target, input)
	if err != nil {
		t.Fatalf("no error expected when initializing a Go project: %v", err)
	}
}
