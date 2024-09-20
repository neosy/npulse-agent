package main

import (
	"os"

	"git.n-hub.ru/neosy/npulse-agent/internal/config"
	clientsREST "git.n-hub.ru/neosy/npulse-agent/internal/transport/rest/clients"
	usecase "git.n-hub.ru/neosy/npulse-agent/internal/usercase"
	"github.com/spf13/pflag"
)

var (
	cfg *config.Config
)

func init() {
	cfg = config.New()
}

func main() {
	var clients interface{}

	url := pflag.StringP("url", "a", cfg.ClientWatcher.URL, "Список url адресов сервера Watcher")

	port := pflag.IntP("port", "p", cfg.ClientWatcher.Port, "Порт сервера Watcher")

	// Разбираем флаги
	pflag.Parse()

	if len(os.Args) == 1 && cfg.ClientWatcher.URL == "" {
		pflag.PrintDefaults() // Выводим список параметров
		os.Exit(0)
	}

	if *url != "" && *port != 0 {
		cfg.ClientWatcher.URL = *url
		cfg.ClientWatcher.URLs = cfg.ClientWatcher.URLsGet()
		cfg.ClientWatcher.Port = *port
	}

	// Инициализация UseCase
	switch cfg.ClientWatcher.Type {
	case "REST":
		clients = clientsREST.New(cfg)
	}
	uc := usecase.New(clients)

	uc.WatcherReg()
}
