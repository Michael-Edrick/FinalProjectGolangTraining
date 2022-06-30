package utils

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Port             string `env:"PORT"`
	SecretKey        string `env:"SECRET_KEY"`
	PostgresHost     string `env:"POSTGRES_HOST"`
	PostgresPort     string `env:"POSTGRES_PORT"`
	PostgresUser     string `env:"POSTGRES_USER"`
	PostgresPassword string `env:"POSTGRES_PASSWORD"`
	PostgresName     string `env:"POSTGRES_NAME"`
}

func InitConfig() *Config {
	var config Config
	cleanenv.ReadConfig(".env", &config)
	return &config
}
