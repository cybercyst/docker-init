package discover

import (
	"docker-init/internal/types"
	"fmt"

	"github.com/spf13/afero"
	"golang.org/x/mod/modfile"
)

func (d *Detector) isGolang() bool {
	_, err := afero.ReadFile(d.Fs, "go.mod")
	if err != nil {
		return false
	}

	return true
}

func (d *Detector) getGolangInfo() (*types.TemplateInfo, error) {
	gomodBytes, err := afero.ReadFile(d.Fs, "go.mod")
	if err != nil {
		return nil, err
	}

	resp, err := modfile.Parse("go.mod", gomodBytes, nil)
	if err != nil {
		return nil, err
	}

	return &types.TemplateInfo{
		Label: "Go",
		Name:  "gomod",
		Input: map[string]interface{}{
			"image":  fmt.Sprintf("golang:%s-alpine", resp.Go.Version),
			"module": resp.Module.Mod.Path,
		},
	}, nil
}
