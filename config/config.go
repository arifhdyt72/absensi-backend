package config

import "github.com/spf13/viper"

type Config struct {
	DBHost        string `mapstructure:"DB_HOST"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	PORT          string `mapstructure:"PORT"`
	BASE_URL      string `mapstructure:"BASE_URL"`
}

func LoadConfig(path string, arguments []string) (config Config, err error) {

	if len(arguments) <= 1 {
		return
	}

	viper.AddConfigPath("./config/" + path)
	viper.SetConfigName(arguments[1])
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
