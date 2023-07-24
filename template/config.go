package template

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func loadConfig() {
	env := os.Getenv("env")
	path, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("Failed to get current working directory: %v", err))
	}

	viper.SetConfigType("yaml")

	if env == "" {
		env = "dev"
	}

	viper.SetConfigFile(filepath.Join(path, "config", fmt.Sprintf("config-%s", env)))

	if err = viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("Failed to load config file: %v", err))
	}

}
