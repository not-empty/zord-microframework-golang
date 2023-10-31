package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
}

func NewConfig() *Config {
	config := Config{}
	err := config.loadEnvs()
	if err != nil {
		panic(err)
	}
	return &config
}

func (c *Config) ReadConfig(Key string) string {
	env := os.Getenv(Key)
	if env == "" {
		panic(fmt.Sprintf("[Config] Unable to read environment variable: %v", Key))
	}
	return env
}

func (c *Config) ReadNumberConfig(Key string) int {
	env := os.Getenv(Key)
	if env == "" {
		panic(fmt.Sprintf("[Config] Unable to read environment variable: %v", Key))
	}
	config, err := strconv.Atoi(env)
	if err != nil {
		panic(fmt.Sprintf("[Config] Jwt Expiration (minutes) must be integer: %v", Key))
	}
	return config
}

func (c *Config) ReadArrayConfig(Key string) []string {
	env := os.Getenv(Key)
	if env == "" {
		panic(fmt.Sprintf("[Config] Unable to read environment variable: %v", Key))
	}
	config := strings.Split(env, ",")
	return config
}

func (c *Config) loadEnvs() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}
	return nil
}
