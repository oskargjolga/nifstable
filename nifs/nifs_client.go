package nifs

import (
	"net/http"
	"time"
)

type NifsClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewNifsClient(baseURL string) *NifsClient {
	return &NifsClient{
		BaseURL:    baseURL,
		HTTPClient: &http.Client{Timeout: 5 * time.Second},
	}
}
