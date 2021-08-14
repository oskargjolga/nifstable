package main

import (
	"fmt"

	"github.com/oskargjolga/nifstable/nifs"
)

func main() {
	nifsClient := nifs.NewNifsClient()
	fmt.Println(nifsClient.BaseURL)

	matches := nifsClient.GetMatches()
	fmt.Println(len(matches))
	fmt.Println(matches)

}
