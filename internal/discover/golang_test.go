package discover

import (
	"docker-init/internal/types"
	"strings"
	"testing"

	"github.com/spf13/afero"
	"gotest.tools/assert"
)

func TestDetectGolang119(t *testing.T) {
	fs := afero.NewMemMapFs()
	afero.WriteFile(fs, "go.mod", []byte(strings.TrimSpace(`
module go-test

go 1.19
`)), 0644)

	detector := &Detector{
		Fs: fs,
	}

	got, err := detector.Detect()
	assert.NilError(t, err)

	want := &types.TemplateInfo{
		Name:  "gomod",
		Label: "Go",
		Input: map[string]interface{}{
			"module": "go-test",
			"image":  "golang:1.19-alpine",
		},
	}
	assert.DeepEqual(t, got, want)
}

func TestDetectGolang120(t *testing.T) {
	fs := afero.NewMemMapFs()
	afero.WriteFile(fs, "go.mod", []byte(strings.TrimSpace(`
module go-test

go 1.20
`)), 0644)

	detector := &Detector{
		Fs: fs,
	}

	got, err := detector.Detect()
	assert.NilError(t, err)

	want := &types.TemplateInfo{
		Name:  "gomod",
		Label: "Go",
		Input: map[string]interface{}{
			"module": "go-test",
			"image":  "golang:1.20-alpine",
		},
	}
	assert.DeepEqual(t, got, want)
}
