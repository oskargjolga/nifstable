package nifs

import (
	"fmt"
	"os"
	"sort"
	"strconv"

	tablewriter "github.com/olekukonko/tablewriter"
)

type Table struct {
	Name     string
	EntryMap map[int]TableEntry
}

func (t *Table) TableEntries() []TableEntry {
	var tableEntries = []TableEntry{}
	for _, v := range t.EntryMap {
		tableEntries = append(tableEntries, v)
	}
	sort.Sort(ByPointsAndGoalDiff(tableEntries))
	return tableEntries
}

func (t *Table) Render() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Pos", "Club", "P", "W", "D", "L", "GF", "GA", "GD", "Pts"})
	for i, entry := range t.TableEntries() {
		table.Append([]string{
			fmt.Sprintf("#%d", i+1),
			entry.TeamName,
			strconv.Itoa(entry.MatchesPlayed),
			strconv.Itoa(entry.Wins),
			strconv.Itoa(entry.Draws),
			strconv.Itoa(entry.Losses),
			strconv.Itoa(entry.Goals),
			strconv.Itoa(entry.GoalsAgainst),
			strconv.Itoa(entry.GoalDiff()),
			strconv.Itoa(entry.Points),
		})
	}
	table.Render()
}

func NewTable(matches []Match, halftime bool) *Table {
	table := &Table{}
	table.Name = "Eliteserien 2020"
	table.EntryMap = make(map[int]TableEntry)
	for _, match := range matches {
		if match.MatchStatusID != 2 { // Not started (Also "result unknown")
			table.AddMatchResults(match, halftime)
		}
	}
	return table
}

func (t *Table) AddMatchResults(match Match, halftime bool) {
	for _, result := range match.GetResults(halftime) {
		if entry, ok := t.EntryMap[result.TeamId]; ok {
			entry.AddResult(result)
			t.EntryMap[result.TeamId] = entry
		} else {
			entry := TableEntry{TeamName: result.TeamName}
			entry.AddResult(result)
			t.EntryMap[result.TeamId] = entry
		}
	}
}
