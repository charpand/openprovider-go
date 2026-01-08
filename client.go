package openprovider

import (
	"net/http"
	"time"
)

const (
	DefaultBaseURL = "https://api.openprovider.eu"
)

type Config struct {
	BaseURL    string
	HTTPClient *http.Client
}

type Client struct {
	baseURL    string
	httpClient *http.Client
}

func NewClient(config Config) *Client {
	baseURL := config.BaseURL
	if baseURL == "" {
		baseURL = DefaultBaseURL
	}

	httpClient := config.HTTPClient
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: time.Second * 30,
		}
	}

	return &Client{
		baseURL:    baseURL,
		httpClient: httpClient,
	}
}
