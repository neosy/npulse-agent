package watcher

import (
	"errors"
	"fmt"
	"os"

	appdto "git.n-hub.ru/neosy/npulse-agent/application/dto"
	"git.n-hub.ru/neosy/npulse-agent/pkg/nnet"
)

func (u *Watcher) Reg() error {
	var activeURL string

	for _, url := range u.config.URLs {
		if url == "" {
			continue
		}

		url = fmt.Sprintf("%s:%s", url, u.config.Port)

		if u.Ping(url) == nil {
			activeURL = url
			break
		}
	}

	if activeURL == "" {
		return errors.New("servers are not available")
	}

	hostName, _ := os.Hostname()
	hostIP, _ := nnet.HostIP()

	req := &appdto.WatcherRegRequest{
		ServerURL: activeURL,
		IPAddress: hostIP,
		HostName:  hostName,
	}

	return u.watcherClient.Reg(req)
}
