package template

import (
	"docker-new/internal/types"
	"os"
	"testing"
)

func TestGenerateGo(t *testing.T) {
	tmpDir := t.TempDir()
	os.Chdir("../..")

	target := types.Target{
		TargetType: types.Go,
		Path:       tmpDir,
		Input: &map[string]interface{}{
			"binary": "go-api",
		},
	}

	err := Generate(target)
	if err != nil {
		t.Fatalf("no error expected when initializing a Go project: %v", err)
	}
}
