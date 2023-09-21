package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Environment   string
	App           string
	Version       string
	DbUser        string
	DbPass        string
	DbUrl         string
	DbPort        string
	Database      string
	Secret        string
	JwtExpiration int
	AccessSecret  []string
	AccessContext []string
	AccessToken   []string
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
	c.Secret = c.ReadConfig("JWT_SECRET")
	c.JwtExpiration = c.ReadNumberConfig("JWT_EXPIRATION")
	c.AccessContext = c.ReadArrayConfig("ACCESS_CONTEXT")
	c.AccessSecret = c.ReadArrayConfig("ACCESS_SECRET")
	c.AccessToken = c.ReadArrayConfig("ACCESS_TOKEN")
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
