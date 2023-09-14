package config

import (
	"os"
	"strings"
)

type Config struct {
	HttpConfig *Http
}

type Http struct {
	Port string
	Host string
}

func InitConfig() (*Config, error) {
	config := Config{
		HttpConfig: &Http{
			Port: GetEnvAsStr("HTTP_PORT", ":8080"),
			Host: GetEnvAsStr("HTTP_HOST", "localhost"),
		},
	}
	return &config, nil
}

func GetEnvAsStr(name string, defaultValue string) string {
	valStr := os.Getenv(name)
	if strings.TrimSpace(valStr) == "" {
		return defaultValue
	}

	return valStr
}
