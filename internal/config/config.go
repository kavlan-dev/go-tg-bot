package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type config struct {
	Environment string
	Token       string
}

func InitConfig() (*config, error) {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")

	v.AddConfigPath("./config")
	v.AddConfigPath("../../config")

	v.SetDefault("env", "prod")

	if err := v.ReadInConfig(); err != nil {
		return &config{}, err
	}

	config := config{
		Environment: v.GetString("env"),
		Token:       v.GetString("token"),
	}

	if config.Environment != "dev" && config.Environment != "prod" {
		return nil, fmt.Errorf("Окружение %s не найдено", config.Environment)
	}

	return &config, nil
}
