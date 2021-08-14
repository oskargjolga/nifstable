package nifs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func (c *NifsClient) GetMatches() []Match {
	response, err := http.Get(c.BaseURL + "/tournaments/" + TournamentId + "/stages/" + StageId + "/matches/")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var matches []Match
	json.Unmarshal(responseData, &matches)

	return matches

}

// Generated using https://mholt.github.io/json-to-go/
type Match struct {
	Timestamp time.Time `json:"timestamp"`
	Name      string    `json:"name"`
	Result    struct {
		HomeScore45 int    `json:"homeScore45"`
		HomeScore90 int    `json:"homeScore90"`
		AwayScore45 int    `json:"awayScore45"`
		AwayScore90 int    `json:"awayScore90"`
		Type        string `json:"type"`
		UID         string `json:"uid"`
	} `json:"result"`
	HomeTeam struct {
		Logo struct {
			URL         string `json:"url"`
			ImageTypeID int    `json:"imageTypeId"`
			Type        string `json:"type"`
			UID         string `json:"uid"`
			ID          int    `json:"id"`
		} `json:"logo"`
		MatchStatistics interface{} `json:"matchStatistics"`
		Name            string      `json:"name"`
		Names           interface{} `json:"names"`
		Clubs           []struct {
			Name           string `json:"name"`
			PrimaryColor   string `json:"primaryColor"`
			SecondaryColor string `json:"secondaryColor"`
			TextColor      string `json:"textColor"`
			Type           string `json:"type"`
			UID            string `json:"uid"`
			ID             int    `json:"id"`
			SportID        int    `json:"sportId"`
		} `json:"clubs"`
		ShortName   string `json:"shortName"`
		Type        string `json:"type"`
		UID         string `json:"uid"`
		ID          int    `json:"id"`
		ExternalIds struct {
			Tv2  []int `json:"tv2"`
			Fiks []int `json:"fiks"`
		} `json:"externalIds"`
	} `json:"homeTeam"`
	AwayTeam struct {
		Logo struct {
			URL         string `json:"url"`
			ImageTypeID int    `json:"imageTypeId"`
			Type        string `json:"type"`
			UID         string `json:"uid"`
			ID          int    `json:"id"`
		} `json:"logo"`
		MatchStatistics interface{} `json:"matchStatistics"`
		Name            string      `json:"name"`
		Names           interface{} `json:"names"`
		Clubs           []struct {
			Name           string `json:"name"`
			PrimaryColor   string `json:"primaryColor"`
			SecondaryColor string `json:"secondaryColor"`
			TextColor      string `json:"textColor"`
			Type           string `json:"type"`
			UID            string `json:"uid"`
			ID             int    `json:"id"`
			SportID        int    `json:"sportId"`
		} `json:"clubs"`
		Type        string `json:"type"`
		UID         string `json:"uid"`
		ID          int    `json:"id"`
		ExternalIds struct {
			Tv2  []int `json:"tv2"`
			Fiks []int `json:"fiks"`
		} `json:"externalIds"`
	} `json:"awayTeam"`
	MatchStatusID int         `json:"matchStatusId"`
	MatchTypeID   interface{} `json:"matchTypeId"`
	Stadium       struct {
		Name    string `json:"name"`
		Type    string `json:"type"`
		UID     string `json:"uid"`
		ID      int    `json:"id"`
		SportID int    `json:"sportId"`
	} `json:"stadium"`
	Attendance int `json:"attendance"`
	Round      int `json:"round"`
	LiveFeeds  []struct {
		CustomerID  interface{} `json:"customerId"`
		MatchID     int         `json:"matchId"`
		Prioritized bool        `json:"prioritized"`
		Ended       bool        `json:"ended"`
		User        struct {
			Name          string      `json:"name"`
			TwitterHandle interface{} `json:"twitterHandle"`
			Type          string      `json:"type"`
			UID           string      `json:"uid"`
			ID            int         `json:"id"`
		} `json:"user"`
		Type        string      `json:"type"`
		UID         string      `json:"uid"`
		ID          int         `json:"id"`
		ExternalIds interface{} `json:"externalIds"`
		SportID     int         `json:"sportId"`
	} `json:"liveFeeds"`
	Comment     string `json:"comment"`
	Tv2Ids      []int  `json:"tv2Ids"`
	CoveredLive bool   `json:"coveredLive"`
	StageID     int    `json:"stageId"`
	TvChannels  []struct {
		Name        string      `json:"name"`
		ProgID      interface{} `json:"progId"`
		Type        string      `json:"type"`
		UID         string      `json:"uid"`
		ID          int         `json:"id"`
		ExternalIds interface{} `json:"externalIds"`
		SportID     int         `json:"sportId"`
	} `json:"tvChannels"`
	Odds []struct {
		NorskTippingMatchID            int       `json:"norskTippingMatchId"`
		NorskTippingBetObjectID        int       `json:"norskTippingBetObjectId"`
		NorskTippingBetObjectNumber    int       `json:"norskTippingBetObjectNumber"`
		OddsHome                       float64   `json:"oddsHome"`
		OddsDraw                       float64   `json:"oddsDraw"`
		OddsAway                       float64   `json:"oddsAway"`
		MinimumNumberOfMatchesOnCoupon int       `json:"minimumNumberOfMatchesOnCoupon"`
		MaximumNumberOfMatchesOnCoupon int       `json:"maximumNumberOfMatchesOnCoupon"`
		SaleStartTime                  time.Time `json:"saleStartTime"`
		SaleStopTime                   time.Time `json:"saleStopTime"`
		Type                           string    `json:"type"`
		UID                            string    `json:"uid"`
		ID                             int       `json:"id"`
		SportID                        int       `json:"sportId"`
	} `json:"odds"`
	LastUpdated string `json:"lastUpdated"`
	MatchLength string `json:"matchLength"`
	Type        string `json:"type"`
	UID         string `json:"uid"`
	ID          int    `json:"id"`
	ExternalIds struct {
		Tv2             []int `json:"tv2"`
		Fiks            int   `json:"fiks"`
		FiksMatchNumber int64 `json:"fiksMatchNumber"`
	} `json:"externalIds"`
	SportID int `json:"sportId"`
}
