package config

import (
	"os"
	"time"

	"github.com/spf13/viper"
)

const defaultConfigPath = "./config/config.local.yaml"

type Config struct {
	Logger   LoggerConfig   `yaml:"logger"`
	Auth     AuthConfig     `yaml:"auth"`
	HTTP     HTTPConfig     `yaml:"http"`
	Database DatabaseConfig `yaml:"database"`
}

type LoggerConfig struct {
	Level string `yaml:"level"`
	File  string `yaml:"file"`
}

type AuthConfig struct {
	SigningKey     string        `yaml:"signingKey"`
	AccessTokenTTL time.Duration `yaml:"accessTokenTTL"`
}

type HTTPConfig struct {
	Port int `yaml:"port"`
}

type DatabaseConfig struct {
	Postgres PostgresConfig
}

type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func NewConfig() (*Config, error) {
	var err error
	var config Config

	viper.SetConfigFile(defaultConfigPath)
	if path := os.Getenv("CONFIG_PATH"); path != "" {
		viper.SetConfigFile(path)
	}

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	viper.BindEnv("database.postgres.password", "POSTGRES_PASSWORD")
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
