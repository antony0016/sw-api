package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	DB_HOST     string
	DB_PORT     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_DATABASE string
}

func (env *Env) EnvLoad() {
	godotenv.Load()

	env.DB_HOST = os.Getenv("DB_HOST")
	env.DB_PORT = os.Getenv("DB_PORT")
	env.DB_USERNAME = os.Getenv("DB_USERNAME")
	env.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	env.DB_DATABASE = os.Getenv("DB_DATABASE")
}
