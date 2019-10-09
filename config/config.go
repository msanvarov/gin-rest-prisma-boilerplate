package config

import (
	"github.com/spf13/viper"
	"log"
	"sync"
)

var (
	once sync.Once
	yaml *viper.Viper)

func Configuration(env string) *viper.Viper {
	once.Do(func() {
		yaml = viper.New()
		yaml.SetConfigType("yaml")
		yaml.SetConfigName(env)
		yaml.AddConfigPath(".")
		err := yaml.ReadInConfig()
		if err != nil {
			log.Fatal(err)
		}
	})
	return yaml
}