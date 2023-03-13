package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Environment string
	App         string
	Version     string
}

func NewConfig() *Config {
	config := Config{}
	err := config.loadEnvs()
	if err != nil {
		panic(err)
	}
	return &config
}

func (c *Config) Boot() {
	c.ReadConfig("APP")
	return
}

func (c *Config) ReadConfig(Key string) string {
	env := os.Getenv(Key)
	if env == "" {
		panic(fmt.Sprintf("[Config] Unable to read environment variable: %v", Key))
	}
	return env
}

func (c *Config) loadEnvs() error {
	err := godotenv.Load(".env", ".config.env")
	if err != nil {
		return err
	}
	return nil
}
