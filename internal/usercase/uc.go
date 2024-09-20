package usercase

import clientsREST "git.n-hub.ru/neosy/npulse-agent/internal/transport/rest/clients"

type UseCase struct {
	clients interface{}
}

func New(clients interface{}) *UseCase {
	return &UseCase{
		clients: clients,
	}
}

func (uc *UseCase) WatcherReg() {
	uc.clients.(*clientsREST.Clients).Watcher.Registration()
}
