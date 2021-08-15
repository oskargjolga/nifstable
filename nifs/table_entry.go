package nifs

import "fmt"

type TableEntry struct {
	TeamName      string
	MatchesPlayed int
	Goals         int
	GoalsAgainst  int
	Points        int
	Wins          int
	Losses        int
	Draws         int
}

func (t *TableEntry) AddResult(result MatchResult) {
	t.MatchesPlayed++
	t.Goals += result.Goals
	t.GoalsAgainst += result.GoalsAgainst
	t.Points += result.Points
	if result.Points == 3 {
		t.Wins++
	} else if result.Points == 1 {
		t.Draws++
	} else {
		t.Losses++
	}
}

func (t *TableEntry) GoalDiff() int {
	return t.Goals - t.GoalsAgainst
}

func (t TableEntry) String() string {
	return fmt.Sprintf("%s 		%d", t.TeamName, t.Points)
}
