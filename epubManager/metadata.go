package epubManager

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Package struct {
	XMLName  xml.Name `xml:"package"`
	Metadata Metadata `xml:"metadata"`
}

type Metadata struct {
	ID     string `xml:"identifier"`
	Title  string `xml:"title"`
	Author string `xml:"creator"`
	Lang   string `xml:"language"`
	Date   string `xml:"date"`
}

func GetMetadata(dir string) (*Metadata, error) {
	// Open XML file
	xmlFile, err := os.Open(filepath.Join(dir, "EPUB", "content.opf"))
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()

	// Read the opened xmlFile as a byte array.
	byteValue, _ := io.ReadAll(xmlFile)

	// Initialize Package object
	var pkg Package

	err = xml.Unmarshal(byteValue, &pkg)
	if err != nil {
		fmt.Println("Error parsing XML:", err)
		return nil, err
	}

	pkg.Metadata.ID = strings.Replace(pkg.Metadata.ID, "urn:uuid:", "", 1)

	return &pkg.Metadata, nil
}
