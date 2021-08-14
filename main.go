package main

import (
	"fmt"

	"github.com/oskargjolga/nifstable/nifs"
)

const (
	BaseURL      = "https://api.nifs.no"
	TournamentId = "5"      // Tournament id for eliteserien
	StageId      = "682936" // Stage id for eliteserien 2020 season
)

func main() {

	nifsClient := nifs.NewNifsClient(BaseURL)
	matches := nifsClient.FetchMatches(TournamentId, StageId)
	fmt.Println(len(matches))

}
