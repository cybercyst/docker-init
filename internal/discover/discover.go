package discover

import (
	"docker-init/internal/types"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func getTargetType(file string) types.TargetType {
	switch file {
	case "go.mod":
		return types.Go
	}
	return types.None
}

func getInput(targetType types.TargetType, targetPath string) (map[string]interface{}, error) {
	input := map[string]interface{}{}

	switch targetType {
	case types.Go:
		gomodBytes, err := os.ReadFile(targetPath)
		if err != nil {
			return nil, err
		}

		for _, line := range strings.Split(string(gomodBytes), "\n") {
			if strings.HasPrefix(line, "module") {
				module := strings.Split(line, " ")[1]
				input["binary"] = module
			}

			// HACK to setup for gin projects
			if strings.Contains(line, "gin-gonic/gin") {
				input["port"] = 8080
			}
		}

		return input, nil
	}

	err := fmt.Errorf("no input generator found for target type %v", targetType.ToString())
	return nil, err
}

func ScanFolderForTargets(fsys fs.FS) ([]types.Target, error) {
	targets := []types.Target{}

	files, err := fs.ReadDir(fsys, ".")
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		targetType := getTargetType(file.Name())
		if targetType == types.None {
			continue
		}

		targetPath, err := filepath.Abs(file.Name())
		if err != nil {
			return nil, err
		}
		input, err := getInput(targetType, targetPath)
		if err != nil {
			return nil, err
		}
		target := types.Target{
			TargetType: targetType,
			Path:       filepath.Dir(targetPath),
			Input:      input,
		}
		targets = append(targets, target)
	}

	return targets, err
}
