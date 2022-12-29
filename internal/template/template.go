package template

import (
	"docker-init/internal/template/generators"
	"docker-init/internal/types"
	"fmt"
	"os"
	"path/filepath"
)

func Generate(target types.Target) error {
	body, err := generateContent(target)
	if err != nil {
		return err
	}

	os.WriteFile(filepath.Join(target.Path, "Dockerfile"), []byte(body), 0666)
	if err != nil {
		return err
	}

	return nil
}

func generateContent(target types.Target) (string, error) {
	switch target.TargetType {
	case types.Go:
		return generators.GenerateGo(nil)
	}

	return "", fmt.Errorf("no generator found for target type %s at path %s", target.TargetType.ToString(), target.Path)
}
