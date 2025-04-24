package rclients

import watchercl "git.n-hub.ru/neosy/npulse-agent/adapter/outbound/rest/watcher"

// Clients the implementation of REST clients
type Clients struct {
	Watcher *watchercl.WatcherClient
}

// New returns a new instance of REST clients
func New(watcherConfig *watchercl.Config) *Clients {
	return &Clients{
		Watcher: watchercl.NewWatcherClient(watcherConfig),
	}
}
