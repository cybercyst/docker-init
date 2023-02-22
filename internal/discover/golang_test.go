package discover

import (
	"testing"

	"github.com/spf13/afero"
	"gotest.tools/assert"
)

func TestShouldDetectModNameAndVersion(t *testing.T) {
	fs := afero.NewMemMapFs()
	afero.WriteFile(fs, "go.mod", []byte("module go-test\n\ngo 1.19\n"), 0644)

	info, err := NewGolangProject(fs)
	if err != nil {
		t.Fatalf("got unexpected error when parsing golang project: %s", err)
	}

	assert.Equal(t, info.BuildRuntime.Type, Go)
	assert.Equal(t, info.BuildRuntime.Version, "1.19")
	assert.Equal(t, info.Runtime.Type, None)
	assert.Equal(t, info.Runtime.Version, "")
}
