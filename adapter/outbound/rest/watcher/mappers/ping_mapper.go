package mappers

import (
	"errors"

	"git.n-hub.ru/neosy/npulse-agent/adapter/outbound/rest/watcher/dto"
)

func (m *Mappers) MapToPingRequest() *dto.PingRequest {
	return &dto.PingRequest{
		Text: "Ping",
	}
}

func (m *Mappers) MapPingResponseToError(resp *dto.PingResponse) error {
	if resp.Text != "Pong" {
		return errors.New("incorrect response")
	}

	return nil
}
