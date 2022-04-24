package c2api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	WorldURL   = "https://api.collective2.com/world/apiv3/"
	PrivateURL = "https://api.collective2.com/platform/apiv3/" // don't forget the platform id
)

// The interface maybe used for supplying the alternative httpClient
type httpClient interface {
	Do(request *http.Request) (response *http.Response, err error)
}

type Client struct {
	client  httpClient
	baseURL *url.URL
	apiKey  string

	// public API of the client
	Strategy  *StrategyAccessService
	Signal    *SignalEntryService
	AutoTrade *AutoTradeManagementService
}

func NewClient(httpClient httpClient, baseURL string, apiKey string) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	if !strings.HasSuffix(baseURL, "/") {
		baseURL += "/"
	}

	parsedBaseURL, err := url.Parse(baseURL)

	if err != nil {
		return nil, err
	}

	c := &Client{
		client:  httpClient,
		baseURL: parsedBaseURL,
		apiKey:  apiKey,
	}
	c.AutoTrade = &AutoTradeManagementService{client: c}
	c.Strategy = &StrategyAccessService{client: c}
	c.Signal = &SignalEntryService{client: c}

	return c, nil
}

func NewDefaultClient(apiKey string) (*Client, error) {
	return NewClient(nil, WorldURL, apiKey)
}

func (c *Client) Do(request *http.Request, v interface{}) (response *http.Response, err error) {
	httpResp, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}

	code := httpResp.StatusCode
	if code < 200 || code > 299 {
		return httpResp, fmt.Errorf("request failed with status code=%d", code)
	}

	if v != nil {
		defer httpResp.Body.Close()
		err = json.NewDecoder(httpResp.Body).Decode(v)
	}

	return httpResp, err
}

func (c *Client) NewPostRequest(relativeUrl string, v interface{}) (req *http.Request, err error) {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(v)

	url, _ := c.baseURL.Parse(relativeUrl)
	req, err = http.NewRequest("POST", url.String(), b)

	if err != nil {
		return req, err
	}

	req.Header.Set("Content-Type", "application/json")
	return req, nil
}
