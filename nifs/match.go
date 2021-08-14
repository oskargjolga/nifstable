package nifs

import (
	"time"
)

func (match *Match) GetHomeTeamId() int {
	return match.HomeTeam.ID
}

func (match *Match) GetAwayTeamId() int {
	return match.AwayTeam.ID
}

func (match *Match) GetHomeTeamName() string {
	return match.HomeTeam.Name
}

func (match *Match) GetAwayTeamName() string {
	return match.AwayTeam.Name
}

func (match *Match) GetHomeTeamResult(halftime bool) Result {
	result := Result{
		teamName: match.HomeTeam.Name,
		points:   match.getHomeTeamPoints(halftime),
	}
	if halftime {
		result.goals = match.Result.HomeScore45
	} else {
		result.goals = match.Result.HomeScore90
	}
	return result
}

func (match *Match) GetAwayTeamResult(halftime bool) Result {
	result := Result{
		teamName: match.AwayTeam.Name,
		points:   match.getAwayTeamPoints(halftime),
	}
	if halftime {
		result.goals = match.Result.AwayScore45
	} else {
		result.goals = match.Result.AwayScore90
	}
	return result
}

func (match *Match) getAwayTeamPoints(halftime bool) int {
	if halftime {
		if match.Result.AwayScore45 > match.Result.HomeScore45 {
			return 3
		} else if match.Result.AwayScore45 == match.Result.HomeScore45 {
			return 1
		} else {
			return 0
		}
	} else {
		if match.Result.AwayScore90 > match.Result.HomeScore90 {
			return 3
		} else if match.Result.AwayScore90 == match.Result.HomeScore90 {
			return 1
		} else {
			return 0
		}
	}
}

func (match *Match) getHomeTeamPoints(halftime bool) int {
	if halftime {
		if match.Result.HomeScore45 > match.Result.AwayScore45 {
			return 3
		} else if match.Result.HomeScore45 == match.Result.AwayScore45 {
			return 1
		} else {
			return 0
		}
	} else {
		if match.Result.HomeScore90 > match.Result.AwayScore90 {
			return 3
		} else if match.Result.HomeScore90 == match.Result.AwayScore90 {
			return 1
		} else {
			return 0
		}
	}
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
