package cmd

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig() error {
	// Set config file name and paths
	viper.SetConfigName(".envcheck")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME")

	//Set default
	viper.SetDefault("schema", "schema.yaml")
	viper.SetDefault("strict", false)
	viper.SetDefault("verbose", false)

	// Read config file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil
		}
		return fmt.Errorf("error reading config file: %w", err)
	}

	return nil
}