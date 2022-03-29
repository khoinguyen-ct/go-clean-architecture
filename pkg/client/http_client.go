package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"go.elastic.co/apm/module/apmhttp"
	"golang.org/x/net/context/ctxhttp"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

var httpClient *http.Client
var timeout = 5 * time.Second

func init() {
	trans := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   timeout,
			KeepAlive: timeout,
		}).DialContext,
		TLSHandshakeTimeout:   4 * time.Second,
		ResponseHeaderTimeout: 6 * time.Second,
		ExpectContinueTimeout: 4 * time.Second,
		DisableKeepAlives:     false,
		MaxIdleConnsPerHost:   1024,
		MaxConnsPerHost:       2048,
	}
	client := http.Client{
		Transport: trans,
		Timeout:   timeout,
	}
	httpClient = apmhttp.WrapClient(&client)
}

type HTTPClient struct {
	name string
}

func (h *HTTPClient) SendHTTPRequest(ctx context.Context, method, uri string, requestData []byte, res interface{}) error {
	req, err := http.NewRequestWithContext(ctx, method, uri, bytes.NewReader(requestData))
	if err != nil {
		return nil
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := ctxhttp.Do(ctx, httpClient, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, res)
	return err
}

func (h *HTTPClient) SetName(name string) {
	h.name = name
}

func (h *HTTPClient) GetName() string {
	return h.name
}

func (h *HTTPClient) IsDisabled() bool {
	return false
}

func NewHttpClient(name string) *HTTPClient {
	var httpClient = HTTPClient{}
	httpClient.SetName(name)
	return &httpClient
}
