package discover

import (
	"testing"

	"github.com/spf13/afero"
	"gotest.tools/assert"
)

func TestScanFolderNoTargetFound(t *testing.T) {
	fs := afero.NewMemMapFs()
	afero.WriteFile(fs, "a.txt", []byte("My amazing text file"), 0644)
	afero.WriteFile(fs, "b.txt", []byte("Another life-changing file"), 0644)

	d := &Detector{
		Fs: fs,
	}

	_, err := d.Detect()
	assert.Error(t, err, "unable to detect project type")
}
