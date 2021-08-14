package main

import (
	"fmt"

	"github.com/oskargjolga/nifstable/nifs"
)

func main() {
	nifsClient := nifs.NewClient()
	fmt.Println(nifsClient.BaseURL)

}
