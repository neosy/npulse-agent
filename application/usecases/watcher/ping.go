package watcher

func (u *Watcher) Ping(url string) error {
	return u.watcherClient.Ping(url)
}
