package iconfig

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

func (c *Config) parseFlag() {
	appName := os.Args[0]

	versionFlag := pflag.BoolP("version", "v", false, "Show version")
	urls := pflag.StringP("url", "a", c.WatcherConfig.URLs, "List of server url addresses")
	port := pflag.StringP("port", "p", c.WatcherConfig.Port, "Server port")

	// Проверяем наличие флага -h или --help до вызова pflag.Parse()
	pflag.Usage = func() {
		fmt.Printf("%s agent for monitoring the operation of servers and workstations\n", c.AppName)

		// Выводим пример использования
		fmt.Println()
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

	if versionFlag != nil && *versionFlag {
		fmt.Printf("%s agent for monitoring the operation of servers and workstations\n", c.AppName)
		fmt.Printf("Version %s\n", c.AppConfig.Version)
		os.Exit(0)
	}

	if urls != nil && *urls != "" {
		c.WatcherConfig.URLs = *urls
	}

	if port != nil && *port != "" {
		c.WatcherConfig.Port = *port
	}
}
