package iconfig

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

func (c *Config) parseFlag() {
	appName := os.Args[0]

	urls := pflag.StringP("url", "a", c.WatcherConfig.URLs, "Список url адресов сервера nPulseWatcher")
	port := pflag.StringP("port", "p", c.WatcherConfig.Port, "Порт сервера nPulseWatcher")

	// Проверяем наличие флага -h или --help до вызова pflag.Parse()
	pflag.Usage = func() {
		// Выводим пример использования
		fmt.Printf("Usage example: %s -a <url1 url2 url3...> -p <port>\n", appName)

		// Выводим список параметров
		pflag.PrintDefaults()
	}

	// Разбираем флаги
	pflag.Parse()

	if len(os.Args) == 1 && c.WatcherConfig.URLs == "" {
		pflag.Usage()
		os.Exit(0)
	}

	if urls != nil && *urls != "" {
		c.WatcherConfig.URLs = *urls
	}

	if port != nil && *port != "" {
		c.WatcherConfig.Port = *port
	}
}
