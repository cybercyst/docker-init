package main

import (
	"docker-new/cmd"

	"github.com/docker/cli/cli-plugins/manager"
	"github.com/docker/cli/cli-plugins/plugin"
	"github.com/docker/cli/cli/command"
	"github.com/spf13/cobra"
)

func pluginMain() {
	plugin.Run(func(dockerCli command.Cli) *cobra.Command {
		return cmd.RootCmd
	}, manager.Metadata{
		SchemaVersion: "0.1.0",
		Vendor:        "Docker Inc.",
		Version:       cmd.Version,
	})
}

func main() {
	pluginMain()
}
