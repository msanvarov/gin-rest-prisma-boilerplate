// Package config provides configuration presets.
package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

var (
	once sync.Once
	yaml *viper.Viper
)

// Configure returns a pointer to the yaml configuration variables.
func Configure(fileName string) {
	yaml = viper.New()
	yaml.SetConfigType("yaml")
	yaml.SetConfigName(fileName)
	yaml.AddConfigPath(".")
	err := yaml.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}

// GetConfiguration retrieves the configuration structure.
func GetConfiguration() *viper.Viper {
	return yaml
}
