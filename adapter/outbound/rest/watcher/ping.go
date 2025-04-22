package watchercl

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"git.n-hub.ru/neosy/npulse-agent/adapter/outbound/rest/watcher/dto"
	"github.com/valyala/fasthttp"
)

const timeoutSec = 10

// Ping
func (c *WatcherClient) Ping(url string) error {
	reqData := c.mappers.MapToPingRequest()

	reqDataJson, err := json.Marshal(reqData)
	if err != nil {
		return err
	}

	httpReq := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(httpReq)

	httpReq.Header.SetMethod(c.config.Method)
	httpReq.SetRequestURI(strings.TrimRight(url, "/") + c.config.Paths.Ping)
	httpReq.SetBody(reqDataJson)

	// Устанавливаем заголовки
	httpReq.Header.SetContentType("application/json")
	httpReq.Header.Set("Accept", "application/json")
	// TODO Установить, если требуется
	// req.Header.Set("User-Agent", "WatcherClient/1.0")
	// req.Header.Set("Authorization", "Bearer <TOKEN>")

	httpResp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(httpResp)

	err = fasthttp.DoTimeout(httpReq, httpResp, time.Duration(timeoutSec)*time.Second)
	if err != nil {
		return err
	}

	if httpResp.StatusCode() != fasthttp.StatusOK {
		return fmt.Errorf("http status code: %d", httpResp.StatusCode())
	}

	var respDto = &dto.PingResponse{}

	err = json.Unmarshal(httpResp.Body(), respDto)
	if err != nil {
		return err
	}

	return c.mappers.MapPingResponseToError(respDto)
}
