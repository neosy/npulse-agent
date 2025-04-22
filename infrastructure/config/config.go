package iconfig

import (
	"fmt"
	"log"
	"strings"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

// Основные настройки
type Config struct {
	Name string `env:"APP_NAME" envDefault:"nPulseAgent"`

	// Config
	WatcherConfig WatcherConfigDefault
}

type WatcherConfigDefault struct {
	URLs   string `env:"WATCHER_URLS" envDefault:"http://192.168.23.11,http://192.168.23.24"`
	Port   string `env:"WATCHER_PORT" envDefault:"8081"`
	Method string `enc:"WATCHER_METHOD" envDefault:"GET"`
	Paths  WatcherRestPaths
}

// REST Paths
type WatcherRestPaths struct {
	Ping string `env:"WATCHER_PATH_PING" envDefault:"/watcher/ping"`
	Reg  string `env:"WATCHER_PATH_REG" envDefault:"/watcher/reg"`
}

func (c *WatcherConfigDefault) ParseURLs() []string {
	return strings.Split(c.URLs, " ")
}

// Создание объекта Config
func New() *Config {
	c := &Config{}

	c.load()

	c.parseFlag()

	return c
}

// Load config from environment variables
func (config *Config) load() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file, proceeding with environment variables only")
	}
	if err := env.Parse(config); err != nil {
		log.Fatalf("Config load(). Read configuration error: %s\n", err)
	}
}
