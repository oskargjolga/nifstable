package nifs

import (
	"fmt"
	"os"
	"sort"
	"strconv"

	tablewriter "github.com/olekukonko/tablewriter"
)

type Table struct {
	Name         string
	TableEntries []TableEntry
	entryMap     map[int]TableEntry
}

func (t *Table) Render() {
	for _, v := range t.entryMap {
		t.TableEntries = append(t.TableEntries, v)
	}
	sort.Sort(ByPointsAndGoalDiff(t.TableEntries))

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Pos", "Club", "P", "W", "D", "L", "GF", "GA", "GD", "Pts"})
	for i, entry := range t.TableEntries {
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
	table.entryMap = make(map[int]TableEntry)
	for _, match := range matches {
		table.AddMatchResults(match, halftime)
	}
	return table
}

func (t *Table) AddMatchResults(match Match, halftime bool) {
	for _, result := range match.GetResults(halftime) {
		if !result.HasResult {
			return
		}
		if entry, ok := t.entryMap[result.TeamId]; ok {
			entry.AddResult(result)
			t.entryMap[result.TeamId] = entry
		} else {
			entry := TableEntry{TeamName: result.TeamName}
			entry.AddResult(result)
			t.entryMap[result.TeamId] = entry
		}
	}
}
