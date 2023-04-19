package util

import "github.com/spf13/viper"

type Config struct {
	SERVER_ADDRESS   string `mapstructure:"SERVER_ADDRESS"`
	AUTH_SRV_ADDRESS string `mapstructure:"AUTH_SRV_ADDRESS"`
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
