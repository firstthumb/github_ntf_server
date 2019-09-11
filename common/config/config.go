package config

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/berglas/pkg/berglas"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
	"os"
	"strings"
)

type Config struct {
	Client struct {
		RepositoryServiceUrl string `yaml:"repositoryService" envconfig:"repositoryService"`
	} `yaml:"client"`
	Secret struct {
		DB string `yaml:"db" envconfig:"DB"`
	} `yaml:"secret"`
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

func GetSecret(obj string) (string, error) {
	// Check environment variables
	val, exists := os.LookupEnv(obj)
	if exists {
		return val, nil
	}

	// Fetch from Google KMS
	ctx := context.Background()

	resp, err := berglas.Access(ctx, &berglas.AccessRequest{
		Bucket: "github-ntf-secrets",
		Object: obj,
	})
	if err != nil {
		return "", err
	}

	return string(resp), nil
}
