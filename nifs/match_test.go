package nifs

import "testing"

func TestGetResult(t *testing.T) {
	var tests = []struct {
		name    string
		matchId string

		wantHalfTimeGoalsHome int
		wantHalfTimeGoalsAway int
		wantFullTimeGoalsHome int
		wantFullTimeGoalsAway int

		wantHalfTimePointsHome int
		wantHalfTimePointsAway int
		wantFullTimePointsHome int
		wantFullTimePointsAway int
	}{
		{"Rosenborg-Kristiansund(0,0)", "1800723", 0, 0, 0, 0, 1, 1, 1, 1},
		{"Strømsgodset-Sandefjord(1, 2)->(3, 4)", "1805724", 1, 2, 3, 4, 0, 3, 0, 3},
	}

	c := NewNifsClient("https://api.nifs.no")

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			match := c.FetchMatch(test.matchId)
			homeResultHalfTime := match.GetHomeTeamResult(true)
			awayResultHalfTime := match.GetAwayTeamResult(true)
			homeResultFullTime := match.GetHomeTeamResult(false)
			awayResultFullTime := match.GetAwayTeamResult(false)
			if homeResultHalfTime.goals != test.wantHalfTimeGoalsHome {
				t.Errorf("want %d, got %d", test.wantHalfTimeGoalsHome, homeResultHalfTime.goals)
			}
			if awayResultHalfTime.goals != test.wantHalfTimeGoalsAway {
				t.Errorf("want %d, got %d", test.wantHalfTimeGoalsAway, awayResultHalfTime.goals)
			}
			if homeResultFullTime.goals != test.wantFullTimeGoalsHome {
				t.Errorf("want %d, got %d", test.wantFullTimeGoalsHome, homeResultFullTime.goals)
			}
			if awayResultFullTime.goals != test.wantFullTimeGoalsAway {
				t.Errorf("want %d, got %d", test.wantFullTimeGoalsAway, awayResultFullTime.goals)
			}
			if homeResultHalfTime.points != test.wantHalfTimePointsHome {
				t.Errorf("want %d, got %d", test.wantHalfTimePointsHome, homeResultHalfTime.points)
			}
			if awayResultHalfTime.points != test.wantHalfTimePointsAway {
				t.Errorf("want %d, got %d", test.wantHalfTimePointsAway, awayResultHalfTime.points)
			}
			if homeResultFullTime.points != test.wantFullTimePointsHome {
				t.Errorf("want %d, got %d", test.wantFullTimePointsHome, homeResultFullTime.points)
			}
			if awayResultFullTime.points != test.wantFullTimePointsAway {
				t.Errorf("want %d, got %d", test.wantFullTimePointsAway, awayResultFullTime.points)
			}
		})
	}

}
