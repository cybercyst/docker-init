package cmd

import (
	"docker-init/internal/discover"
	"docker-init/internal/template"
	"fmt"
	"os"

	"github.com/bclicn/color"
	"github.com/enescakir/emoji"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "docker-init",
	Short: "Use Docker with your existing projects",
	Long: `Do you want to have all the great benefits of using Docker with an existing
project and no idea where to start?

- Set up a Docker dev environment for this project
- Bootstrap and configure your production-facing build artifact

Docker init makes it simple!`,
	Run: func(cmd *cobra.Command, args []string) {
		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Printf("unable to get current directory! %v ", err)
			os.Exit(1)
		}

		targets, err := discover.ScanFolderForTargets(os.DirFS(currentDir))
		if err != nil {
			fmt.Printf("unable to determine targets from current directory! %v ", err)
			os.Exit(1)
		}

		if len(targets) == 0 {
			fmt.Println(emoji.SeeNoEvilMonkey, "Well this is embarassing. We were unable to initialize Docker for this directory")
			os.Exit(1)
		}

		for _, target := range targets {
			fmt.Println()
			fmt.Println(emoji.PartyPopper, color.Green("SUCCESS"), "We found a", color.BBlue(target.TargetType.ToString()), "project!", emoji.PartyPopper)
			fmt.Println()
			err := template.Generate(target)
			if err != nil {
				fmt.Printf("error while generating files: %v", err)
				os.Exit(1)
			}
			fmt.Println()
			fmt.Println(emoji.CheckBoxWithCheck, " Finished setting up Docker for your", color.BBlue(target.TargetType.ToString()), "project.")
		}

		fmt.Println()
		fmt.Println(emoji.Rocket, "Run", color.BBlue("docker compose up"), "to get started!", emoji.Rocket)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.docker-init.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".docker-init" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".docker-init")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
