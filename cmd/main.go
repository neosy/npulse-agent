package main

import (
	"fmt"
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

	appName := os.Args[0]

	url := pflag.StringP("url", "a", cfg.ClientWatcher.URL, "Список url адресов сервера nPulseWatcher")
	port := pflag.IntP("port", "p", cfg.ClientWatcher.Port, "Порт сервера nPulseWatcher")

	// Проверяем наличие флага -h или --help до вызова pflag.Parse()
	pflag.Usage = func() {
		// Выводим пример использования
		fmt.Printf("Usage example: %s -a <url1,url2,url...> -p <port>\n", appName)

		// Выводим список параметров
		pflag.PrintDefaults()
	}

	// Разбираем флаги
	pflag.Parse()

	if len(os.Args) == 1 && cfg.ClientWatcher.URL == "" {
		pflag.Usage()
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
