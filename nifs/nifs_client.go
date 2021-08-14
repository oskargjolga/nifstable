package nifs

import "net/http"

const (
	BaseURL      = "https://api.nifs.no"
	TournamentId = 5
)

type NifsClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewClient() *NifsClient {
	return &NifsClient{
		BaseURL:    BaseURL,
		HTTPClient: http.DefaultClient,
	}
}
