package mappers

import (
	"errors"

	"git.n-hub.ru/neosy/npulse-agent/adapter/outbound/rest/watcher/dto"
	appdto "git.n-hub.ru/neosy/npulse-agent/application/dto"
)

func (m *Mappers) MapWatcherRegRequestToRegRequest(req *appdto.WatcherRegRequest) *dto.RegRequest {
	return &dto.RegRequest{
		IPAddress: req.IPAddress,
		HostName:  req.HostName,
	}
}

func (m *Mappers) MapRegResponseToError(resp *dto.RegResponse) error {
	if resp.Status != "success" {
		return errors.New("incorrect response")
	}

	return nil
}
