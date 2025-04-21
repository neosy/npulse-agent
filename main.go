package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"git.n-hub.ru/neosy/npulse-agent/application/usecases"
	"git.n-hub.ru/neosy/npulse-agent/application/usecases/watcher"
	iconfig "git.n-hub.ru/neosy/npulse-agent/infrastructure/config"
)

func main() {
	cfg := iconfig.New()

	ctx, cancel := context.WithCancel(context.Background())

	// Создаем обработчик с уровнем Info, используя HandlerOptions
	handlerOptions := &slog.HandlerOptions{
		Level: slog.LevelInfo, // Устанавливаем уровень логирования
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, handlerOptions))

	// Usecases
	watcherConfig := &watcher.Config{
		URLs: cfg.WatcherConfig.ParseURLs(),
		Port: cfg.WatcherConfig.Port,
	}
	uc := usecases.New(logger, watcherConfig)
	// Initialize
	uc.Init(ctx)

	// Захват сигналов завершения (Ctrl+C, SIGTERM)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Ждем сигнал завершения или отмены контекста
	select {
	case <-ctx.Done():
		logger.ErrorContext(ctx, "Context complete, shutting down services...")
	case sig := <-sigChan:
		logger.ErrorContext(ctx, fmt.Sprintf("Signal received: %v, shutting down...", sig))
		cancel()
	}
}
