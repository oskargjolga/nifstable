package nifs

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

const (
	TEAM_NAME_WIDTH = 20
)

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

func (t *TableEntry) separator() string {
	return strings.Repeat(" ", TEAM_NAME_WIDTH-utf8.RuneCountInString(t.TeamName))
}

func (t TableEntry) String() string {
	return fmt.Sprintf("%s%s%d", t.TeamName, t.separator(), t.Points)
}
