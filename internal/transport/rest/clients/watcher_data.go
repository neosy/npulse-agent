package clients

import "git.n-hub.ru/neosy/npulse-agent/internal/models"

type watcherPingRequest struct {
	models.WatcherPingRequest
}

type watcherPingResponse struct {
	models.WatcherPingResponse
}

type watcherRegRequest struct {
	models.WatcherRegRequest
}

type watcherRegResponse struct {
	models.WatcherRegSuccessResponse
}
