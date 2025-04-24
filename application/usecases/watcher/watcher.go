package watcher

import (
	"log/slog"

	watchercl "git.n-hub.ru/neosy/npulse-agent/adapter/outbound/rest/watcher"
)

type Config struct {
	URLs []string
	Port string
}

type Watcher struct {
	logger *slog.Logger
	config *Config

	// REST clients
	watcherClient *watchercl.WatcherClient
}

func NewWatcher(
	logger *slog.Logger,
	config *Config,

	// REST clients
	restWatchrClient *watchercl.WatcherClient,
) *Watcher {
	return &Watcher{
		logger: logger,
		config: config,

		// REST clients
		watcherClient: restWatchrClient,
	}
}
