// Contains tests for tiles.go
package tiles

import (
	"testing"
	"path/filepath"
	"fmt"
)

// Render the test template
var validDefinitionPath, _ = filepath.Abs("test/tiles-definitions.xml")

// Tests for passing empty strings to tiles.Render(string, string)
func TestRender_empty(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	// Both arguments empty
	renderedTemplate, err := Render("", "")
	assertHasError(err, t)
	assertEmpty(renderedTemplate, t)

	// Template name empty
	renderedTemplate, err = Render("", validDefinitionPath)
	assertHasError(err, t)
	assertEmpty(renderedTemplate, t)

	// TilesDefinition path empty
	renderedTemplate, err = Render("testTemplate", "")
	assertHasError(err, t)
	assertEmpty(renderedTemplate, t)
}

// Tests for passing invalid strings to tiles.Render(string, string)
func TestRender_invalid(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	invalidTemplateName := "invalidTemplateName"
	invalidDefinitionPath := "/Users/test/invalid.xml"

	// Both arguments invalid
	renderedTemplate, err := Render(invalidTemplateName, invalidDefinitionPath)
	assertHasError(err, t)
	assertEmpty(renderedTemplate, t)

	// Template name invalid
	renderedTemplate, err = Render(invalidTemplateName, validDefinitionPath)
	assertHasError(err, t)
	assertEmpty(renderedTemplate, t)

	// TilesDefinition path invalid
	renderedTemplate, err = Render("testTemplate", invalidDefinitionPath)
	assertHasError(err, t)
	assertEmpty(renderedTemplate, t)
}

// Tests for passing valid strings to tiles.Render(string, string)
func TestRender_valid(t *testing.T) {
	renderedTemplate, err := Render("testTemplate", validDefinitionPath)
	assertDoesNotHaveError(err, t)
	assertNotEmpty(renderedTemplate, t)
	assertEquals(expectedData, renderedTemplate, t)

	fmt.Println(renderedTemplate)
}

// Simple wrapper to assert that the given string is empty
func assertEmpty(template string, t *testing.T) {
	if template != "" {
		t.Error("Expected rendered template to be empty.", template)
	}
}

// Simple wrapper to assert that the given string is not empty
func assertNotEmpty(template string, t *testing.T) {
	if template == "" {
		t.Error("Expected rendered template to not be empty", template)
	}
}

// Simple wrapper to assert that the given error is not nil
func assertHasError(err error, t *testing.T) {
	if err == nil {
		t.Error("Expected error but didn't receive one.")
	}
}

// Simple wrapper to assert that the given error is nil
func assertDoesNotHaveError(err error, t *testing.T) {
	if err != nil {
		t.Error("Unexpected error occurred.", err)
	}
}

// Simple wrapper to assert that the given string equals the expected string
func assertEquals(expected string, actual string, t *testing.T) {
	if expected != actual {
		t.Error("Given strings are not equal.")
	}
}

var expectedData = `<html>

    <head>

        <!-- Page title -->
        <title>Golang Tiles Implementation 1</title>

        <!-- metas, styles, scripts -->
        <meta charset="UTF-8">
<meta name="description" content="Apache Tiles port for Golang">
<meta name="keywords" content="Apache,Tiles,XML,Golang">
<meta name="author" content="Adam Childs">
        <link type="text/css" href="fakeStyles.css" />
        <!-- Imported script -->
<script type="text/javascript" src="fakeScript.js"></script>

<!-- Inline script -->
<script type="text/javascript">
    window.location = "http://play.golang.org";
</script>

    </head>

    <body>

        <!-- Body content -->
        <div id="mainContent">

    <p>Hello Golang Tiles!</p>

</div>

    </body>

</html>`