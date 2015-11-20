package main

import (
	"fmt"
	"github/adchilds/tiles"
	"os"
	"path/filepath"
)

func main() {
	// Render the test template
	tilesDefinitionPath, err := filepath.Abs("TilesGo/src/github/adchilds/tiles/test/tiles-definitions.xml")
	if err != nil {
		fmt.Println(err)
	}

	renderedTemplate, err := tiles.Render("testTemplate", tilesDefinitionPath)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(renderedTemplate)

	os.Exit(0)
}
