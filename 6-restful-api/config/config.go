package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config interface {
	Get(key string) string
}

type config struct{}

func New(filesnames ...string) Config {
	err := godotenv.Load(filesnames...)
	if err != nil {
		panic(err)
	}
	return &config{}
}

func (c *config) Get(key string) string {
	return os.Getenv(key)
}
