package c2api

import (
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

	return c, nil
}

func NewDefaultClient(apiKey string) (*Client, error) {
	return NewClient(nil, WorldURL, apiKey)
}
