package main

import (
	"flag"

	"github.com/oskargjolga/nifstable/nifs"
)

const (
	BaseURL      = "https://api.nifs.no"
	TournamentId = "5"      // Tournament id for eliteserien
	StageId      = "682936" // Stage id for eliteserien 2020 season
)

func main() {
	halftime := flag.Bool("halftime", false, "compute table based on first 45 minutes of all matches")
	flag.Parse()

	nifsClient := nifs.NewNifsClient(BaseURL)
	matches := nifsClient.FetchMatches(TournamentId, StageId)
	table := nifs.NewTable(matches, *halftime)
	table.Render()
}
