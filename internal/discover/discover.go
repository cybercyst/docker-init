package discover

import (
	"docker-init/internal/types"
	"errors"

	"github.com/spf13/afero"
)

type Detector struct {
	Fs afero.Fs
}

func NewDetector(projectPath string) (*Detector, error) {
	fs := afero.NewBasePathFs(afero.NewOsFs(), projectPath)

	isDir, err := afero.IsDir(fs, ".")
	if err != nil {
		return nil, err
	}

	if isDir == false {
		return nil, errors.New("provided project path must be a directory")
	}

	return &Detector{
		Fs: fs,
	}, nil
}

func (d *Detector) Detect() (*types.TemplateInfo, error) {
	switch {
	case d.isGolang():
		return d.getGolangInfo()
	default:
		return nil, errors.New("unable to detect project type")
	}
}
