package nifs

import "net/http"

const (
	BaseURL      = "https://api.nifs.no"
	TournamentId = "5"
	StageId      = "682936" // Stage id for eliteserien 2020
)

type NifsClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewNifsClient() *NifsClient {
	return &NifsClient{
		BaseURL:    BaseURL,
		HTTPClient: http.DefaultClient,
	}
}
