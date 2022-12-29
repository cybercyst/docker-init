package discover

import (
	"docker-init/internal/types"
	"fmt"
	"io/fs"
)

func getTarget(file string) types.TargetType {
	switch file {
	case "go.mod":
		return types.Go
	}
	return types.None
}

func ScanFolderForTargets(fsys fs.FS, path string) ([]types.Target, error) {
	targets := []types.Target{}

	files, err := fs.ReadDir(fsys, path)
	for _, file := range files {
		if file.IsDir() {
			fmt.Printf("%q is a directory, skipping", path)
			continue
		}

		targetType := getTarget(file.Name())
		if targetType != types.None {
			target := types.Target{
				TargetType: targetType,
				Path:       path,
			}
			targets = append(targets, target)
		}
	}

	return targets, err
}
