package nifs

import (
	"fmt"
	"sort"
	"strings"
	"unicode/utf8"
)

type Table struct {
	Name         string
	TableEntries []TableEntry
}

func (t Table) String() string {
	var s string
	s += "\n"
	s += t.Name + "\n"
	th := t.tableHeader()
	s += strings.Repeat("-", utf8.RuneCountInString(th)) + "\n"
	s += t.tableHeader()
	s += "\n"
	for i, entry := range t.TableEntries {
		s +=

			fmt.Sprintf("#%d	%v\n", i+1, entry)
	}
	return s
}

func (t Table) tableHeader() string {
	s := "Pos\t"
	s += "Club"
	s += strings.Repeat(" ", TEAM_NAME_MAX_WIDTH-utf8.RuneCountInString("Club"))
	s += "P\t\t"
	s += "W\t"
	s += "D\t"
	s += "L\t"
	s += "GF\t"
	s += "GF\t"
	s += "GD\t"
	s += "Pts\t"

	s += "\n"
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
	sort.Sort(ByPointsAndGoalDiff(table.TableEntries))

	return table
}
