package epubManager

import (
	"encoding/xml"
	"io"
	"os"
	"path/filepath"
)

type NavPoint struct {
	Title    string     `xml:"navLabel>text"`
	Content  Content    `xml:"content"`
	Children []NavPoint `xml:"navPoint"`
}

type Content struct {
	Link string `xml:"src,attr"`
}
type Toc struct {
	NavPoints []NavPoint `xml:"navMap>navPoint"` // Assuming navMap is the parent node
}

func GetTableOfContents(dir string) (*Toc, error) {
	// parse the toc.ncx file
	t := Toc{}
	tocXml, err := os.Open(filepath.Join(dir, "EPUB", "toc.ncx"))
	if err != nil {
		return nil, err
	}
	defer tocXml.Close()

	byteValue, err := io.ReadAll(tocXml)
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(byteValue, &t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
