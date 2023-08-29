package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Environment string
	App         string
	Version     string
	DbUser      string
	DbPass      string
	DbUrl       string
	DbPort      string
	Database    string
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
	c.Environment = c.ReadConfig("ENVIRONMENT")
	c.Version = c.ReadConfig("VERSION")
	c.App = c.ReadConfig("APP")
	c.DbUser = c.ReadConfig("DB_USER")
	c.DbPass = c.ReadConfig("DB_PASS")
	c.DbUrl = c.ReadConfig("DB_URL")
	c.DbPort = c.ReadConfig("DB_PORT")
	c.Database = c.ReadConfig("DB_DATABASE")
}

func (c *Config) ReadConfig(Key string) string {
	env := os.Getenv(Key)
	if env == "" {
		panic(fmt.Sprintf("[Config] Unable to read environment variable: %v", Key))
	}
	return env
}

func (c *Config) loadEnvs() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}
	return nil
}
