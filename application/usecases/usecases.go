package usecases

import (
	"log/slog"

	rclients "git.n-hub.ru/neosy/npulse-agent/adapter/outbound/rest"
	"git.n-hub.ru/neosy/npulse-agent/application/usecases/watcher"
)

// Usecases represents the business layer of the application.
type Usecases struct {
	Watcher *watcher.Watcher
}

// Dependencies contains external dependencies required by the Usecases.
type Dependencies struct {
	RestClients *rclients.Clients
}

// New returns a new instance of Usecases.
func New(
	logger *slog.Logger,

	// Config
	watcherConfig *watcher.Config,

	// Dependenies
	deps *Dependencies,
) *Usecases {

	return &Usecases{
		Watcher: watcher.NewWatcher(logger, watcherConfig, deps.RestClients.Watcher),
	}
}
