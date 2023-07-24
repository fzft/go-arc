package goarc

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

var confFiles = []string{"questions.yaml"}

// LoadConf loads the configuration files
func LoadConf() {
	path, err := os.Getwd()

	if err != nil {
		panic(fmt.Sprintf("Failed to get current working directory: %v", err))
	}
	viper.SetConfigType("yaml")

	for _, configFile := range confFiles {
		viper.SetConfigFile(filepath.Join(path, "conf", configFile))

		err = viper.MergeInConfig()
		if err != nil {
			panic(fmt.Sprintf("Failed to load config file: %v", err))
		}
	}
}
