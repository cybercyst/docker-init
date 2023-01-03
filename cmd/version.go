package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	Version        = "dev"
	CommitHash     = "n/a"
	BuildTimestamp = "n/a"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get the current version of docker-init",
	Run: func(cmd *cobra.Command, args []string) {
		version := fmt.Sprintf(`
docker-init: Bootstrap your project with Docker

Version: %s
GoVersion: %s
Architecture: %s
GitCommit: %s
Created-at: %s`,
			Version, runtime.Version(), runtime.GOARCH, CommitHash, BuildTimestamp)

		fmt.Println(version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
