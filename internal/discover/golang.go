package discover

import (
	"github.com/spf13/afero"
	"golang.org/x/mod/modfile"
)

func NewGolangProject(fs afero.Fs) (*Info, error) {
	gomodBytes, err := afero.ReadFile(fs, "go.mod")
	if err != nil {
		return nil, err
	}

	resp, err := modfile.Parse("go.mod", gomodBytes, nil)
	if err != nil {
		return nil, err
	}

	return &Info{
		BuildRuntime: Runtime{
			Type:    Go,
			Version: resp.Go.Version,
		},
		Runtime: Runtime{
			Type:    None,
			Version: "",
		},
	}, nil
}
