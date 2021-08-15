package nifs

type MatchResult struct {
	HasResult    bool
	TeamId       int
	TeamName     string
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
