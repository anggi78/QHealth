package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type (
	DBConfig struct {
		DBPORT     string
		DBHOST     string
		DBUSER     string
		DBPASSWORD string
		DBNAME     string
	}
	AppConfig struct {
		SERVERPORT string
	}
	Configs struct {
		DB  DBConfig
		APP AppConfig
	}
)

func LoadConfig() Configs {
	godotenv.Load(".env")

	db := DBConfig{
		DBPORT:     os.Getenv("DB_PORT"),
		DBHOST:     os.Getenv("DB_HOST"),
		DBUSER:     os.Getenv("DB_USER"),
		DBPASSWORD: os.Getenv("DB_PASSWORD"),
		DBNAME:     os.Getenv("DB_NAME"),
	}

	app := AppConfig{
		SERVERPORT: os.Getenv("SERVER_PORT"),
	}

	return Configs{
		APP: app,
		DB:  db,
	}
}
