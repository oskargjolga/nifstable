package nifs

import (
	"sort"
	"testing"
)

func TestSortByPointsAndGoalDiff(t *testing.T) {
	var entries = []TableEntry{
		{TeamName: "A", Points: 40, Goals: 30, GoalsAgainst: 20},
		{TeamName: "B", Points: 10, Goals: 30, GoalsAgainst: 20},
		{TeamName: "C", Points: 20, Goals: 30, GoalsAgainst: 20},
		{TeamName: "D", Points: 50, Goals: 30, GoalsAgainst: 20},
		{TeamName: "E", Points: 20, Goals: 30, GoalsAgainst: 15},
	}
	var wantOrder = []string{"D", "A", "E", "C", "B"}
	sort.Sort(ByPointsAndGoalDiff(entries))
	for i, entry := range entries {
		if entry.TeamName != wantOrder[i] {
			t.Errorf("want %s got %s", wantOrder[i], entry.TeamName)
		}
	}

}
