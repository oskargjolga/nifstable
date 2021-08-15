package main

import (
	"flag"

	"github.com/oskargjolga/nifstable/nifs"
)

const (
	BASE_URL                  = "https://api.nifs.no"
	TOURNAMENT_ID_ELITESERIEN = "5"      // Tournament id for eliteserien
	STAGE_ID_2020             = "682936" // Stage id for eliteserien 2020 season
	STAGE_ID_2022             = "684880"
)

func main() {
	halftime := flag.Bool("halftime", false, "compute table based on first 45 minutes of all matches")
	flag.Parse()

	nifsClient := nifs.NewNifsClient(BASE_URL)
	matches := nifsClient.FetchMatches(TOURNAMENT_ID_ELITESERIEN, STAGE_ID_2020)
	table := nifs.NewTable(matches, *halftime)
	table.Render()
}
