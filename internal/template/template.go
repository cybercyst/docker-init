package template

import (
	"docker-init/internal/types"
	"fmt"
	"os"
	"path/filepath"

	"github.com/bclicn/color"
	"github.com/cybercyst/go-scaffold/pkg/scaffold"
)

func Generate(info *types.TemplateInfo, outputPath string) error {
	templateDir, err := getTemplateDir(info.Name)
	if err != nil {
		return err
	}

	template, err := scaffold.Download(templateDir)
	if err != nil {
		return err
	}

	metadata, err := scaffold.Generate(template, &info.Input, outputPath)
	if err != nil {
		return err
	}

	for _, file := range *metadata.CreatedFiles {
		fmt.Println(color.Green("CREATE"), file)
	}

	return nil
}

func getTemplateDir(templateName string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	// TODO: read from config value for this path
	rootTemplateDir := filepath.Join(homeDir, ".docker-init", "templates")
	return filepath.Join(rootTemplateDir, templateName), nil
}
