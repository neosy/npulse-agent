package config

import (
	"log"
	"strings"

	"github.com/caarlos0/env/v11"
)

// Основные настройки
type Config struct {
	Name          string `env:"APP_NAME" envDefault:"nPulseAgent"`
	ClientWatcher ClientWatcherConfig
}

// Настройки клиента
type ClientWatcherConfig struct {
	Type      string `env:"CLIENT_WATCHER_TYPE" envDefault:"REST"`
	URL       string `env:"CLIENT_WATCHER_URL" envDefault:""`
	URLs      []string
	Port      int    `env:"CLIENT_WATCHER_PORT" envDefault:"8080"`
	Method    string `enc:"CLIENT_WATCHER_METHOD" envDefault:"GET"`
	Endpoints ClientWatcherEndpoints
}

// Endpoints
type ClientWatcherEndpoints struct {
	Ping string `env:"CLIENT_WATCHER_CMD_PING" envDefault:"/watcher/ping"`
	Reg  string `env:"CLIENT_WATCHER_CMD_REG" envDefault:"/watcher/reg"`
}

// Создание объекта Config
func New() *Config {
	c := &Config{}

	c.load()

	c.ClientWatcher.URLs = c.ClientWatcher.URLsGet()

	return c
}

func (config *ClientWatcherConfig) URLsGet() []string {
	return strings.Split(config.URL, ",")
}

// Load config from environment variables
func (config *Config) load() {
	if err := env.Parse(config); err != nil {
		log.Fatalf("Config load(). Read configuration error: %s\n", err)
	}
}
