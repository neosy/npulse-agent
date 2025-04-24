package watchercl

import "git.n-hub.ru/neosy/npulse-agent/adapter/outbound/rest/watcher/mappers"

type Config struct {
	Method string
	Paths  Paths
}

type Paths struct {
	Ping string
	Reg  string
}

// WatcherClient represents a client
type WatcherClient struct {
	// mappers provides data mapping functionality
	mappers *mappers.Mappers
	config  *Config
}

// NewWatcherClient returns a new instance of WatcherClient
func NewWatcherClient(config *Config) (client *WatcherClient) {
	client = &WatcherClient{
		mappers: mappers.NewMappers(),
		config:  config,
	}

	return
}
