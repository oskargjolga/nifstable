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
}

func (t *Table) Render() {
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
	m := make(map[int]TableEntry)
	for _, match := range matches {
		homeTeamResult := match.GetHomeTeamResult(halftime)
		if !homeTeamResult.HasResult {
			continue
		}
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
