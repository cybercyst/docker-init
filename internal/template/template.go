package template

import (
	"docker-new/internal/types"
	"fmt"
	"os"
	"path/filepath"

	"github.com/bclicn/color"
	"github.com/cybercyst/go-cookiecutter/pkg/cookiecutter"
)

func Generate(target types.Target) error {
	templateDir, err := getTemplateDir(target.TargetType)
	if err != nil {
		return err
	}

	metadata, err := cookiecutter.Generate(templateDir, target.Input, target.Path)
	if err != nil {
		return err
	}

	for _, file := range *metadata.CreatedFiles {
		fmt.Println(color.Green("CREATE"), file)
	}

	return nil
}

func getTemplateDir(targetType types.TargetType) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	rootTemplateDir := filepath.Join(homeDir, ".docker-new", "templates")
	templateDir := ""

	switch targetType {
	case types.Go:
		templateDir = "gomod"
	case types.Angular:
		templateDir = "angular"
	case types.Python:
		templateDir = "pyproject"
	case types.React:
		templateDir = "react"
	}

	if templateDir == "" {
		err = fmt.Errorf("no generator found for target type %s", targetType.ToString())
		return "", err
	}

	return filepath.Join(rootTemplateDir, templateDir), nil
}
