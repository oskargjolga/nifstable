package nifs

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type NifsClient struct {
	BaseURL string
}

func NewNifsClient(baseURL string) *NifsClient {
	return &NifsClient{
		BaseURL: baseURL,
	}
}

func (c *NifsClient) PerformRequest(req *http.Request) ([]byte, error) {
	client := http.Client{
		Timeout: time.Second * 10,
	}
	req.Header.Set("Accept", "application/json")
	res, getErr := client.Do(req)
	if getErr != nil {
		return nil, getErr
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, readErr
	}

	return body, nil
}

func (c *NifsClient) FetchMatches(tournamentId string, stageId string) []Match {

	url := c.BaseURL + "/tournaments/" + tournamentId + "/stages/" + stageId + "/matches/"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	body, err := c.PerformRequest(req)
	if err != nil {
		log.Fatal(err)
	}

	var matches []Match
	jsonErr := json.Unmarshal(body, &matches)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return matches

}

// used for testing
func (c *NifsClient) FetchMatch(matchId string) Match {
	url := c.BaseURL + "/matches/" + matchId

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Accept", "application/json")

	body, err := c.PerformRequest(req)
	if err != nil {
		log.Fatal(err)
	}
	var match Match
	jsonErr := json.Unmarshal(body, &match)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return match
}
