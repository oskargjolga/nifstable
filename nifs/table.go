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
		if val, ok := m[homeTeamResult.TeamId]; ok {
			val.AddResult(homeTeamResult)
			m[homeTeamResult.TeamId] = val
		} else {
			entry := TableEntry{TeamName: match.GetHomeTeamName()}
			entry.AddResult(homeTeamResult)
			m[homeTeamResult.TeamId] = entry
		}

		if val, ok := m[awayTeamResult.TeamId]; ok {
			val.AddResult(awayTeamResult)
			m[awayTeamResult.TeamId] = val
		} else {
			entry := TableEntry{TeamName: match.GetAwayTeamName()}
			entry.AddResult(awayTeamResult)
			m[awayTeamResult.TeamId] = entry
		}
	}
	for _, v := range m {
		table.TableEntries = append(table.TableEntries, v)
	}

	return table
}
