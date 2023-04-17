package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/mercanil/simple-todo-app/model"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
)

var cfgFile string
var TodosStored model.Todos
var StorageFile = "simple-todo-app.json"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "simple-todo-app",
	Short: "Todo is a sample todo manager ",
	Long:  `This project is just for fun.`,
}

func Execute() {
	file, err := os.Open(StorageFile)
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		os.Exit(1)
	}
	json.Unmarshal(bytes, &TodosStored)

	err = rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

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

		// Search config in home directory with name ".simple-todo-app" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".simple-todo-app")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
