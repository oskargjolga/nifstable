package nifs

import (
	"sort"
	"testing"
)

func TestSortByPoints(t *testing.T) {
	var entries = []TableEntry{
		{TeamName: "A", Points: 40},
		{TeamName: "B", Points: 10},
		{TeamName: "C", Points: 20},
		{TeamName: "D", Points: 50},
	}
	var wantOrder = []string{"D", "A", "C", "B"}
	sort.Sort(ByPoints(entries))
	for i, entry := range entries {
		if entry.TeamName != wantOrder[i] {
			t.Errorf("want %s got %s", wantOrder[i], entry.TeamName)
		}
	}

}
