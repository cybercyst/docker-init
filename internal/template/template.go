package template

import (
	"docker-init/internal/types"
	"fmt"
	"os"
	"path/filepath"

	"github.com/bclicn/color"
	"github.com/cybercyst/go-scaffold/pkg/scaffold"
)

func Generate(target types.Target) error {
	templateDir, err := getTemplateDir(target.TargetType)
	if err != nil {
		return err
	}

	template, err := scaffold.Download(templateDir)
	if err != nil {
		return err
	}

	metadata, err := scaffold.Generate(template, target.Input, target.Path)
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
	// TODO: read from config value for this path
	rootTemplateDir := filepath.Join(homeDir, ".docker-init", "templates")
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
