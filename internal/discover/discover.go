package discover

import (
	"docker-init/internal/types"
	"fmt"
	"io/fs"
	"path/filepath"
)

func getTarget(file string) types.TargetType {
	switch file {
	case "go.mod":
		return types.Go
	}
	return types.None
}

func ScanFolderForTargets(fsys fs.FS) ([]types.Target, error) {
	targets := []types.Target{}

	files, err := fs.ReadDir(fsys, ".")
	for _, file := range files {
		if file.IsDir() {
			fmt.Printf("%q is a directory, skipping", file.Name())
			continue
		}

		targetType := getTarget(file.Name())
		filePath, err := filepath.Abs(file.Name())
		if err != nil {
			return nil, err
		}

		if targetType != types.None {
			target := types.Target{
				TargetType: targetType,
				Path:       filepath.Dir(filePath),
			}
			targets = append(targets, target)
		}
	}

	return targets, err
}
