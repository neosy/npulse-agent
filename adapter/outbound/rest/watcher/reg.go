package watchercl

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"git.n-hub.ru/neosy/npulse-agent/adapter/outbound/rest/watcher/dto"
	appdto "git.n-hub.ru/neosy/npulse-agent/application/dto"
	"github.com/valyala/fasthttp"
)

// Reg
func (c *WatcherClient) Reg(req *appdto.WatcherRegRequest) error {
	if req == nil {
		return errors.New("function parameter is a null pointer")
	}

	reqData := c.mappers.MapWatcherRegRequestToRegRequest(req)

	reqDataJson, err := json.Marshal(reqData)
	if err != nil {
		return err
	}

	httpReq := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(httpReq)

	httpReq.Header.SetMethod(c.config.Method)
	httpReq.SetRequestURI(strings.TrimRight(req.ServerURL, "/") + c.config.Paths.Reg)
	httpReq.SetBody(reqDataJson)

	// Устанавливаем заголовки
	httpReq.Header.SetContentType("application/json")
	httpReq.Header.Set("Accept", "application/json")
	// TODO Установить, если требуется
	// httpReq.Header.Set("User-Agent", "WatcherClient/1.0")
	// httpReq.Header.Set("Authorization", "Bearer <TOKEN>")

	httpResp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(httpResp)

	err = fasthttp.DoTimeout(httpReq, httpResp, time.Duration(timeoutSec)*time.Second)
	if err != nil {
		return err
	}

	if httpResp.StatusCode() != fasthttp.StatusOK {
		return fmt.Errorf("http status code: %d", httpResp.StatusCode())
	}

	var respDto = &dto.RegResponse{}

	err = json.Unmarshal(httpResp.Body(), respDto)
	if err != nil {
		return err
	}

	return c.mappers.MapRegResponseToError(respDto)
}
