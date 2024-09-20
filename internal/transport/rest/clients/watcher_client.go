package clients

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"git.n-hub.ru/neosy/npulse-agent/internal/config"
	"git.n-hub.ru/neosy/npulse-agent/internal/pkg/nnet"
	"github.com/neosy/gofw/nbasic"
	"github.com/neosy/gofw/nfasthttp"
)

type WatcherClient struct {
	nfasthttp.Client

	method    string
	endpoints WatcherClientEndpoints
}

type WatcherClientEndpoints struct {
	ping string
	reg  string
}

func NewWatcherClient(config config.ClientWatcherConfig) *WatcherClient {
	endpoints := WatcherClientEndpoints{
		ping: config.Endpoints.Ping,
		reg:  config.Endpoints.Reg,
	}

	client := &WatcherClient{
		method:    config.Method,
		endpoints: endpoints,
	}

	var url string
	for _, value := range config.URLs {
		if value != "" && clientPing(value, config.Port, config.Method, config.Endpoints.Ping) {
			url = value
			break
		}
	}

	if url != "" {
		client.Init(url, config.Port)
	} else {
		log.Printf("Watcher (urls: %s, port: %v) is not available\n", config.URL, config.Port)
		os.Exit(1)
	}

	return client
}

func (client *WatcherClient) Registration() error {
	client.Init(client.Address, client.Port)

	hostName, _ := os.Hostname()
	hostIP, _ := nnet.HostIP()

	reqData := &watcherRegRequest{}
	reqData.ServerIP = hostIP
	reqData.ServerName = hostName

	respData := &watcherRegResponse{}

	// Отправка запроса
	err := client.SendRequest(client.CreateURI(client.endpoints.reg), client.method, reqData)
	defer client.Release()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	resp := client.Response()
	respDataJSON := resp.Body()

	if err = json.Unmarshal(respDataJSON, &respData); err != nil {
		log.Println(nbasic.ErrJSONParsing.Error())
		return err
	}

	if !respData.Success {
		return fmt.Errorf("the client returned an error")
	}

	return nil
}
