package nifs

type MatchResult struct {
	TeamId       int
	TeamName     string
	Points       int
	Goals        int
	GoalsAgainst int
}

func NewMatchResult(teamId int, teamName string, points int, goals int, goalsAgainst int) MatchResult {
	return MatchResult{
		TeamId:       teamId,
		TeamName:     teamName,
		Points:       points,
		Goals:        goals,
		GoalsAgainst: goalsAgainst,
	}
}

func (r *MatchResult) GoalDiff() int {
	return r.Goals - r.GoalsAgainst
}