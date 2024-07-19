package config

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port string
	}
	GRPC struct {
		URL  string
		Port string
	}
	Postgres struct {
		URL      string
		Port     string
		User     string
		Password string
		Database string
	}
	Redis struct {
		URL      string
		Port     string
		Database string
	}
}

func GetRootDir() string {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)
	return filepath.Join(basePath, "../..")
}

func LoadConfig() Config {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(GetRootDir())

	var cfg Config

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	return cfg
}
