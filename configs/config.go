package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB   DBconfig
	Auth AuthConfig
}

type DBconfig struct {
	Dsn  string
	Port string
}

type AuthConfig struct {
	Secret string
}

func LoadConfig() *Config {
	path := "../.env"
	err := godotenv.Load(path)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &Config{
		DB: DBconfig{
			Dsn:  os.Getenv("DSN"),
			Port: os.Getenv("PORT"),
		},
		Auth: AuthConfig{
			Secret: os.Getenv("SECRET"),
		},
	}
}
