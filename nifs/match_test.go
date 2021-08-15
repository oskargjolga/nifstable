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
		{"Mjøndalen-Molde(1, 0)->(1, 3)", "1847913", 1, 0, 1, 3, 3, 0, 0, 3},
	}

	c := NewNifsClient("https://api.nifs.no")

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			match := c.FetchMatch(test.matchId)
			homeResultHalfTime := match.GetHomeTeamResult(true)
			awayResultHalfTime := match.GetAwayTeamResult(true)
			homeResultFullTime := match.GetHomeTeamResult(false)
			awayResultFullTime := match.GetAwayTeamResult(false)
			if homeResultHalfTime.Goals != test.wantHalfTimeGoalsHome {
				t.Errorf("want %d, got %d", test.wantHalfTimeGoalsHome, homeResultHalfTime.Goals)
			}
			if awayResultHalfTime.Goals != test.wantHalfTimeGoalsAway {
				t.Errorf("want %d, got %d", test.wantHalfTimeGoalsAway, awayResultHalfTime.Goals)
			}
			if homeResultFullTime.Goals != test.wantFullTimeGoalsHome {
				t.Errorf("want %d, got %d", test.wantFullTimeGoalsHome, homeResultFullTime.Goals)
			}
			if awayResultFullTime.Goals != test.wantFullTimeGoalsAway {
				t.Errorf("want %d, got %d", test.wantFullTimeGoalsAway, awayResultFullTime.Goals)
			}
			if homeResultHalfTime.Points != test.wantHalfTimePointsHome {
				t.Errorf("want %d, got %d", test.wantHalfTimePointsHome, homeResultHalfTime.Points)
			}
			if awayResultHalfTime.Points != test.wantHalfTimePointsAway {
				t.Errorf("want %d, got %d", test.wantHalfTimePointsAway, awayResultHalfTime.Points)
			}
			if homeResultFullTime.Points != test.wantFullTimePointsHome {
				t.Errorf("want %d, got %d", test.wantFullTimePointsHome, homeResultFullTime.Points)
			}
			if awayResultFullTime.Points != test.wantFullTimePointsAway {
				t.Errorf("want %d, got %d", test.wantFullTimePointsAway, awayResultFullTime.Points)
			}
		})
	}

}
