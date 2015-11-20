package main

import (
	"fmt"
	"github/adchilds/tiles"
	"os"
	"path/filepath"
)

func main() {
	// Retrieve the full path to the tiles definition
	tilesDefinitionPath, err := filepath.Abs("TilesGo/src/github/adchilds/tiles/test/tiles-definitions.xml")
	if err != nil {
		fmt.Println(err)
	}

	// Render the test template
	renderedTemplate, err := tiles.Render("testTemplate", tilesDefinitionPath)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(renderedTemplate)

	os.Exit(0)
}
