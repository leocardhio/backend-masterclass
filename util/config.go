package util

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string, env string) (config Config, err error) {
	viper.SetConfigFile(fmt.Sprintf("%s/%s.env", path, env))

	fmt.Println(viper.GetViper().ConfigFileUsed())

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}

	return
}
