package nifs

type Result struct {
	TeamName     string
	Points       int
	Goals        int
	GoalsAgainst int
}

func (r *Result) GoalDiff() int {
	return r.Goals - r.GoalsAgainst
}
