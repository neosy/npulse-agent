package clients

import (
	"encoding/json"
	"log"

	"git.n-hub.ru/neosy/npulse-agent/internal/config"
	"github.com/neosy/gofw/nbasic"
	"github.com/neosy/gofw/nfasthttp"
)

type Clients struct {
	Watcher *WatcherClient
}

func New(config *config.Config) *Clients {
	return &Clients{
		Watcher: NewWatcherClient(config.ClientWatcher),
	}
}

func clientPing(url string, port int, method string, endpoint string) bool {
	var client nfasthttp.Client

	client.Init(url, port)

	reqData := &watcherPingRequest{}
	reqData.Text = "Ping"

	respData := &watcherPingResponse{}

	// Отправка запроса
	err := client.SendRequest(client.CreateURI(endpoint), method, reqData)
	defer client.Release()
	if err != nil {
		log.Println(err.Error())
		return false
	}

	resp := client.Response()
	respDataJSON := resp.Body()

	if err = json.Unmarshal(respDataJSON, &respData); err != nil {
		log.Println(nbasic.ErrJSONParsing.Error())
		return false
	}

	if respData.Text != "Pong" {
		return false
	}

	return true
}
