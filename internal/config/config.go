package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/Netflix/go-env"
	"gopkg.in/yaml.v3"
)

var defaultAppConfig []byte

// Global defaults global configuration
var Global = &GlobalConfig{
	ApiServerPort: 8080,
	DatabaseURL:   "postgres://postgres:postgres@127.0.0.1:5432/app?sslmode=disable",
	Env:           "development",
	LogLevel:      "info",
	TimeZone:      "Asia/Taipei",
}

type GlobalConfig struct {
	// global
	ApiServerPort int    `env:"API_SERVER_PORT" yaml:"api_server_port"`
	DatabaseURL   string `env:"DATABASE_URL" yaml:"database_url"`
	Env           string `env:"APP_ENV" yaml:"env"`
	LogLevel      string `env:"APP_LOG_LEVEL" yaml:"log_level"`
	TimeZone      string `env:"TZ" yaml:"time_zone"`
}

func Init(opt ...[]byte) (err error) {
	appConfig := defaultAppConfig
	if len(opt) > 0 {
		appConfig = opt[0]
	}
	if err = yaml.Unmarshal(appConfig, Global); err != nil {
		return err
	}

	if _, err = env.UnmarshalFromEnviron(Global); err != nil {
		return err
	}

	if Global.ApiServerPort < 80 {
		return fmt.Errorf("invalid API servier port: %d", Global.ApiServerPort)
	}

	if len(Global.TimeZone) == 0 {
		tz, err := os.ReadFile("/etc/timezone")
		if err == nil && len(tz) > 0 {
			Global.TimeZone = string(tz)
		}
	}

	return nil
}

func IsProdEnv() bool {
	return Global.Env == "production"
}

func IsTestEnv() bool {
	return Global.Env == "test"
}

func IsDevEnv() bool {
	return Global.Env == "development"
}

func CheckEnv(key string) bool {
	val, ok := os.LookupEnv(key)
	if !ok || len(key) == 0 {
		return false
	}
	val = strings.ToLower(val)
	if val == "0" || val == "false" {
		return false
	}
	return true
}
