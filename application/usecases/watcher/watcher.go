package watcher

type Config struct {
	URLs []string
	Port string
}

type Watcher struct {
	config *Config
}

func NewWatcher(config *Config) *Watcher {
	return &Watcher{
		config: config,
	}
}
