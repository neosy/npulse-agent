package main

import (
	"context"
	"log/slog"
	"os"

	rclients "git.n-hub.ru/neosy/npulse-agent/adapter/outbound/rest"
	watchercl "git.n-hub.ru/neosy/npulse-agent/adapter/outbound/rest/watcher"
	"git.n-hub.ru/neosy/npulse-agent/application/usecases"
	"git.n-hub.ru/neosy/npulse-agent/application/usecases/watcher"
	iconfig "git.n-hub.ru/neosy/npulse-agent/infrastructure/config"
)

func main() {
	cfg := iconfig.New()

	ctx := context.Background()

	// Создаем обработчик с уровнем Info, используя HandlerOptions
	handlerOptions := &slog.HandlerOptions{
		Level: slog.LevelInfo, // Устанавливаем уровень логирования
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, handlerOptions))

	// Usecases
	clientWatcherConfig := &watchercl.Config{
		Method: cfg.WatcherConfig.Method,
		Paths: watchercl.Paths{
			Ping: cfg.WatcherConfig.Paths.Ping,
			Reg:  cfg.WatcherConfig.Paths.Reg,
		},
	}
	ucDeps := &usecases.Dependencies{
		RestClients: rclients.New(clientWatcherConfig),
	}
	watcherConfig := &watcher.Config{
		URLs: cfg.WatcherConfig.ParseURLs(),
		Port: cfg.WatcherConfig.Port,
	}
	uc := usecases.New(logger, watcherConfig, ucDeps)
	// Initialize
	uc.Init(ctx)

	// Run
	err := uc.Watcher.Reg()
	if err != nil {
		logger.Error(err.Error())
	}
}
