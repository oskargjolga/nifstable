package nifs

type MatchResult struct {
	TeamId       int
	TeamName     string
	HasResult    bool
	Points       int
	Goals        int
	GoalsAgainst int
}

func NewMatchResult(teamId int, teamName string, points int) MatchResult {
	return MatchResult{
		TeamId:   teamId,
		TeamName: teamName,
		Points:   points,
	}
}
