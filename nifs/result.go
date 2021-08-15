package nifs

type MatchResult struct {
	TeamName     string
	Points       int
	Goals        int
	GoalsAgainst int
}

func (r *MatchResult) GoalDiff() int {
	return r.Goals - r.GoalsAgainst
}
