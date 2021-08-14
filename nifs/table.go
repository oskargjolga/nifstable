package nifs

import "fmt"

type Table struct {
	Name         string
	TableEntries []TableEntry
}

func (t *Table) String() string {
	var s string
	for _, entry := range t.TableEntries {
		s += fmt.Sprintf("%v\n", entry)
	}
	return s
}

func NewTable(matches []Match, halftime bool) *Table {
	table := &Table{}
	table.Name = "Eliteserien 2020"
	m := make(map[int]TableEntry)
	for _, match := range matches {
		homeTeamResult := match.GetHomeTeamResult(halftime)
		awayTeamResult := match.GetAwayTeamResult(halftime)
		if val, ok := m[match.GetHomeTeamId()]; ok {
			val.AddResult(homeTeamResult.goals, homeTeamResult.points)
			m[match.GetHomeTeamId()] = val
		} else {
			entry := TableEntry{TeamName: match.GetHomeTeamName()}
			entry.AddResult(homeTeamResult.goals, homeTeamResult.points)
			m[match.GetHomeTeamId()] = entry
		}

		if val, ok := m[match.GetAwayTeamId()]; ok {
			val.AddResult(awayTeamResult.goals, awayTeamResult.points)
			m[match.GetAwayTeamId()] = val
		} else {
			entry := TableEntry{TeamName: match.GetAwayTeamName()}
			entry.AddResult(awayTeamResult.goals, awayTeamResult.points)
			m[match.GetAwayTeamId()] = entry
		}
	}
	for _, v := range m {
		table.TableEntries = append(table.TableEntries, v)
	}

	return table
}

type TableEntry struct {
	TeamName      string
	MatchesPlayed int
	Goals         int
	Points        int
	Wins          int
	Losses        int
	Draws         int
}

func (t *TableEntry) AddResult(goals int, points int) {
	t.MatchesPlayed++
	t.Goals += goals
	t.Points += points
	if points == 3 {
		t.Wins++
	} else if points == 1 {
		t.Draws++
	} else {
		t.Losses++
	}
}

func (t TableEntry) String() string {
	return fmt.Sprintf("%s 		%d", t.TeamName, t.Points)
}