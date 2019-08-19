package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
	"os"
	"strings"
)

type Config struct {
	Database struct {
		MongoURL      string `yaml:"mongo_url" envconfig:"DB_HOST"`
		MongoUser     string `yaml:"mongo_username" envconfig:"DB_USERNAME"`
		MongoPassword string `yaml:"mongo_password" envconfig:"DB_PASSWORD"`
		DbName        string `yaml:"mongo_db_name" envconfig:"DB_NAME"`
	} `yaml:"database"`
}

func NewConfig() *Config {
	return GetConfig()
}

var cfg *Config

func GetConfig() *Config {
	cfg = &Config{}

	readFile(cfg)
	readEnv(cfg)

	return cfg
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func getEnv(s string) string {
	v := os.Getenv(s)
	if v == "" {
		return "dev"
	}

	return strings.ToLower(v)
}

func readFile(cfg *Config) {
	f, err := os.Open(getEnv("ENV") + ".yml")
	if err != nil {
		processError(err)
	}

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}

func readEnv(cfg *Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		processError(err)
	}
}
