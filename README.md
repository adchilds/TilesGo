# TilesGo
Apache Tiles port for Golang

This project is a simple port of Apache Tiles to Golang. This project currently only supports minimal Tiles features, 
such as concatenating multiple files by replacing tiles insertAttribute statements with either a constant value or dynamic
file contents.

One requirement for using TilesGo is that the output of Render should not have any formatting restraints. TilesGo does 
not guarantee that spaces will be uniform after inserting dynamic content into the base template.

See tiles.Render(string, string)

# Dependencies
Golang (https://golang.org)

# Example
The following is an example usage:

tiles-definition.xml:

    <tiles-definitions>

        <!-- Test Tiles definition -->
        <definition name="testTemplate" template="base.html">
            <put-attribute name="title" value="Golang Tiles Implementation" />
            <put-attribute name="test" value="test.html" />
        </definition>
    
    </tiles-definitions>

base.html:
	
	<!DOCTYPE html>
	<html>
	
		<head>
			<title><tiles:insertAttribute name="title" /></title>
		</head>
		
		<body>
			
			<!-- Dynamic body content -->
			<tiles:insertAttribute name="test" />
			
		</body>
		
	</html>

test.html:

	<div id="bodyContent">
		<p>Hello Tiles (for Golang)!</p>
	</div>

main.go

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

output:

	<html>
    
        <head>
        
            <title>Golang Tiles Implementation</title>

        </head>
    
        <body>
    
            <!-- Dynamic body content -->
            <div id="bodyContent">
        <p>Hello Tiles (for Golang)!</p>
    </div>
    
        </body>
    
    </html>
