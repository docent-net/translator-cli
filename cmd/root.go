package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile         string
	credentialsFile string
	gcpProjectName  string

	rootCmd = &cobra.Command{
		Use:   "translate-cli",
		Short: "CLI translator",
		Long: `CLI translator using Google Translate
			as backend`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("translate-cli: run help command for more information")
		},
	}
)

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName(".translator-cli") // name of config file (without extension)
		viper.AddConfigPath("$HOME/")          // call multiple times to add many search paths
		viper.AddConfigPath("$HOME/.translator-cli/")          // call multiple times to add many search paths
		viper.AddConfigPath(".")               // optionally look for config in the working directory
	}

	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	credentialsFile = viper.GetString("credentialsFile")
	if credentialsFile == "" {
		panic(fmt.Errorf("Google Translate service account config JSON file path not found in config file!"))
	}

	gcpProjectName = viper.GetString("gcpProject")
	if gcpProjectName == "" {
		panic(fmt.Errorf("Google Cloud Project name not found in config file!"))
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.translator-cli.yml)")

	cobra.OnInitialize(initConfig)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
