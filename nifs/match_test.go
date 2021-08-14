package nifs

import "testing"

func TestGetResult(t *testing.T) {
	c := NewNifsClient("https://api.nifs.no")
	match := c.FetchMatch("1800723") // Rosenborg - Kristiansund (0-0)
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

	match2 := c.FetchMatch("1805724") // Strømsgodset - Sandefjord (1, 2) -> (3, 4)
	homeResult2 := match2.GetHomeTeamResult(true)
	awayResult2 := match2.GetAwayTeamResult(true)
	if homeResult2.teamName != "Strømsgodset" {
		t.Error("Home team name is incorrect")
	}
	if homeResult2.goals != 1 {
		t.Error("Home team goals should be 1")
	}
	if awayResult2.goals != 2 {
		t.Error("Away team goals should be 2")
	}
	if homeResult2.points != 0 {
		t.Error("Home team points should be 0")
	}
	if awayResult2.points != 3 {
		t.Error("Away team points should be 3")
	}

}
