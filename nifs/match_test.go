package nifs

import "testing"

func TestGetResult(t *testing.T) {
	c := NewNifsClient("https://api.nifs.no")
	matches := c.FetchMatches("5", "682936")
	match := matches[2] // Rosenborg - Kristiansund (0-0)
	homeResult := match.GetHomeTeamResult(true)
	awayResult := match.GetAwayTeamResult(true)
	if homeResult.teamName != "Rosenborg" {
		t.Error("Home team name is incorrect")
	}
	if homeResult.goals != 0 {
		t.Error("Home team goals should be 0")
	}

	if awayResult.goals != 0 {
		t.Error("Away team goals should be 0")
	}
	if homeResult.points != 1 {
		t.Error("Home team points should be 1")
	}
	if awayResult.points != 1 {
		t.Error("Away team points should be 1")
	}

}
