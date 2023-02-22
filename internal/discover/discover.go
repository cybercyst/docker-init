package discover

import (
	"docker-init/internal/types"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
)

func getPackageJsonKey(fs afero.Fs, targetPath string, key string) (interface{}, error) {
	packageJsonBytes, err := afero.ReadFile(fs, filepath.Join(targetPath, "package.json"))
	if err != nil {
		return nil, err
	}

	packageJson := make(map[string]interface{})
	err = json.Unmarshal(packageJsonBytes, &packageJson)
	if err != nil {
		return nil, err
	}

	return packageJson[key], nil
}

func getTargetType(file string) types.TargetType {
	switch file {
	case "go.mod":
		return types.Go
	case "angular.json":
		return types.Angular
	case "pyproject.toml":
		return types.Python
	case "package.json":
		packageJsonBytes, err := os.ReadFile(file)
		if err != nil {
			panic(err)
		}
		if strings.Contains(string(packageJsonBytes), "react") {
			return types.React
		}
	}
	return types.None
}

func getInput(fs afero.Fs, targetType types.TargetType, targetPath string) (*map[string]interface{}, error) {
	input := make(map[string]interface{})

	switch targetType {
	case types.Go:
		gomodBytes, err := afero.ReadFile(fs, filepath.Join(targetPath, "go.mod"))
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

		return &input, nil
	case types.Angular:
		projectName, err := getPackageJsonKey(fs, targetPath, "name")
		if err != nil {
			return nil, err
		}
		input["project_name"] = projectName

		return &input, nil
	case types.React:
		projectName, err := getPackageJsonKey(fs, targetPath, "name")
		if err != nil {
			return nil, err
		}
		input["project_name"] = projectName

		return &input, nil
	case types.Python:
		// TODO: Actually detect these values and handle more Python project types
		input["app"] = "main:app"
		input["port"] = 8000

		return &input, nil
	}

	err := fmt.Errorf("no input generator found for target type %v", targetType.ToString())
	return nil, err
}

func ScanFolderForTargets(fs afero.Fs) ([]types.Target, error) {
	targets := []types.Target{}

	files, err := afero.ReadDir(fs, ".")
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
		targetPathDir := filepath.Dir(targetPath)
		input, err := getInput(fs, targetType, targetPathDir)
		if err != nil {
			return nil, err
		}
		target := types.Target{
			TargetType: targetType,
			Path:       targetPathDir,
			Input:      input,
		}
		targets = append(targets, target)
	}

	return targets, err
}

type Info struct {
	BuildRuntime Runtime
	Runtime      Runtime
}

type Detector struct {
	Fs afero.Fs
}

func (d *Detector) Detect() (*Info, error) {
	info, err := d.detectRuntime()
	if err != nil {
		return nil, err
	}

	return info, nil
}

func (d *Detector) detectRuntime() (*Info, error) {
	switch {
	case d.isGolang():
		return d.getGolangInfo()
	default:
		return nil, nil
	}
}

func (d *Detector) isGolang() bool {
	_, err := d.Fs.Stat("go.mod")
	if err != nil {
		return false
	}

	return true
}

func (d *Detector) getGolangInfo() (*Info, error) {
	return nil, nil
}
