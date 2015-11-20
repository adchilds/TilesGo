// Provides basic Apache Tiles support for Golang. See https://tiles.apache.org for more information on Tiles. Currently
// the only supported functionality is specifying a base template with insertAttribute statements that can be replaced
// with a corresponding value or file contents.
package tiles

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Represents a collection of Tiles definitions
type TilesDefinitions struct {
	XMLName xml.Name `xml:"tiles-definitions"`
	Definitions []TilesDefinition `xml:"definition"`
}

// Represents an single Tiles definition as outlined in the tiles-definition XML file
type TilesDefinition struct {
	XMLName xml.Name `xml:"definition"`
	Name string `xml:"name,attr"`
	Template string `xml:"template,attr"`
	Attributes []TilesAttribute `xml:"put-attribute"`
}

// Represents a single attribute for a Tiles definition
type TilesAttribute struct {
	XMLName xml.Name `xml:"put-attribute"`
	Name string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

// Renders the template with the given name
func Render(templateName string, tilesDefinitionPath string) (string, error) {

	// Attempt to open the given tiles-definition
	tilesXmlFile, err := openResource(tilesDefinitionPath)
	if err != nil {
		return "", errors.New("Could not parse the given Tiles definition [" + tilesDefinitionPath + "]")
	}

	// Parse the definitions
	tilesDefinitions, err := getTilesDefinitions(tilesXmlFile)
	if err != nil {
		return "", errors.New("Could not parse the given Tiles definition [" + tilesDefinitionPath + "]")
	}

	// Find the definition for the given templateName
	definition, err := getTilesDefinition(templateName, tilesDefinitions)
	if err != nil {
		return "", errors.New("Could not find the given definition [" + templateName + "]")
	}

	// Determine the base path of the tiles-definition (this is where all resources should exist)
	resourcesBasePath, err := filepath.Abs(filepath.Dir(tilesXmlFile.Name()))
	if err != nil {
		return "", errors.New("Could not find base path of resources. Resources base path must be the location of the tiles definition.")
	}

	// Grab the base template file contents
	renderedTemplate, err := getResourceContents(definition.Template, resourcesBasePath)
	if err != nil {
		return "", errors.New("Could not find definition base template definition [" + definition.Name + "]")
	}

	// Replace each '<tiles:insertAttribute name="..." />' with the correct template
	renderedTemplate = populateBaseTemplate(renderedTemplate, definition.Attributes, resourcesBasePath)

	return renderedTemplate, nil
}

// Parses the Tiles definition XML file at the given file path, marshalling the XML to TilesDefinitions
func getTilesDefinitions(xmlFile *os.File) (*TilesDefinitions, error) {
	// Decode/un-marshal the tiles-definition
	tilesDefinitions := &TilesDefinitions{}
	xml.NewDecoder(xmlFile).Decode(&tilesDefinitions)

	return tilesDefinitions, nil
}

// Returns the Tiles definition with the given templateName
func getTilesDefinition(templateName string, definitions *TilesDefinitions) (TilesDefinition, error) {

	for _, element := range definitions.Definitions {

		if element.Name == templateName {
			return element, nil
		}
	}

	return TilesDefinition{}, errors.New("Could not find tiles definition [" + templateName + "]")
}

// Opens a new os.File at the given path
func openResource(tilesDefinitionPath string) (*os.File, error) {
	// Attempt to open the given tiles-definition
	tilesXmlFile, err := os.Open(tilesDefinitionPath)

	// Exit early if we can't open the file
	if err != nil {
		return nil, errors.New("Could not parse the given tiles definition [" + tilesDefinitionPath + "]")
	}

	return tilesXmlFile, nil
}

// Returns the textual contents of the file at the given path
func getResourceContents(filename string, resourceBasePath string) (string, error) {
	var resourcePath = resourceBasePath + string(os.PathSeparator) + filename

	// Open the given resource
	resourceFile, err := openResource(resourcePath)
	if err != nil {
		return "", errors.New("Could not find resource [" + resourcePath + "]")
	}

	// Read the contents of the file
	resourceContents, err := ioutil.ReadAll(resourceFile)
	if err != nil {
		return "", errors.New("Could not read resource contents [" + resourcePath + "]")
	}

	return string(resourceContents), nil
}

// Populates the given base template by replacing all '<tiles:insertAttribute name="..." />' with the contents of the name
// attribute. If the name attribute is not a valid filepath, the string value of name will be used; otherwise, the file
// contents will be used.
func populateBaseTemplate(baseTemplate string, attributes []TilesAttribute, resourcesPath string) string {

	var renderedTemplate string = baseTemplate

	// Replace all tiles:insertAttribute statements with the correct values
	for _, attribute := range attributes {
		var tilesAttributeDef = "<tiles:insertAttribute name=\"" + attribute.Name + "\" />"

		// If the template contains the given attribute, replace the attribute with the corresponding value
		if strings.Contains(baseTemplate, tilesAttributeDef) {
			tilesAttributeContent, err := getResourceContents(attribute.Value, resourcesPath)

			if err != nil {
				// This means the given attribute is not a file, so just use the string value
				renderedTemplate = strings.Replace(renderedTemplate, tilesAttributeDef, attribute.Value, -1)
			} else {
				// The attribute is a filepath, use the file contents
				renderedTemplate = strings.Replace(renderedTemplate, tilesAttributeDef, tilesAttributeContent, -1)
			}
		}
	}

	return renderedTemplate
}