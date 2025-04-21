package usecases

import (
	"log/slog"

	"git.n-hub.ru/neosy/npulse-agent/application/usecases/watcher"
)

// Usecases represents the business layer of the application.
type Usecases struct {
	Watcher *watcher.Watcher
}

// New returns a new instance of Usecases.
func New(
	logger *slog.Logger,

	// Config
	watcherConfig *watcher.Config,
) *Usecases {

	return &Usecases{
		Watcher: watcher.NewWatcher(watcherConfig),
	}
}
