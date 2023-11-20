package util

import (
	"github.com/spf13/viper"
)

// stores all configuration variables for the application
// the values are read by viper from a config or environment variables
type Config struct {
	DBDriver      string `mapstructure:"DB_Driver"`
	DBSource      string `mapstructure:"DB_Source"`
	ServerAddress string `mapstructure:"Server_Address"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
