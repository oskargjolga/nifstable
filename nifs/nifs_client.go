package nifs

import (
	"encoding/json"
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

func (c *NifsClient) GetMatches(tournamentId string, stageId string) []Match {

	url := c.BaseURL + "/tournaments/" + tournamentId + "/stages/" + stageId + "/matches/"
	client := http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Accept", "application/json")

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.StatusCode != http.StatusOK {
		log.Fatal(res.Status)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var matches []Match
	jsonErr := json.Unmarshal(body, &matches)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return matches

}

// only used for testing
func (c *NifsClient) GetMatch(matchId string) Match {
	url := c.BaseURL + "/matches/" + matchId
	client := http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Accept", "application/json")

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.StatusCode != http.StatusOK {
		log.Fatal(res.Status)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var match Match
	jsonErr := json.Unmarshal(body, &match)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return match
}
