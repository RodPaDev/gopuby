package htmlParse

import (
	"os"
	"strings"

	"golang.org/x/net/html"
)

type RenderedText struct {
	Text *strings.Builder
	Tag  string
}

type HtmlPage struct {
	Title string
	Body  []RenderedText
}

// Todo: links should be rendered as [link](url) and if terminal supports inline links, then it should be rendered as clickable links
// Todo: images should be rendered as [image](url) and if terminal supports inline images, then it should be rendered as images otherwise clickable links

func (hp *HtmlPage) BuildText(n *html.Node, inPreCode bool, currentIndex *int) {
	if n.Type == html.TextNode {
		if n.Parent != nil && (n.Parent.Data == "script" || n.Parent.Data == "style") {
			return
		} else if n.Parent != nil && n.Parent.Data == "title" {
			hp.Title = n.Data
		} else {
			// Append the text to the last element in the body
			renderedText := RenderedText{&strings.Builder{}, n.Parent.Data}
			renderedText.Text.WriteString(n.Data)
			hp.Body = append(hp.Body, renderedText)
			*currentIndex += 1
		}
	}

	if n.Type == html.ElementNode && n.Data == "pre" {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.ElementNode && c.Data == "code" {
				// Find the opening <code> tag and add a newline before it
				r := RenderedText{&strings.Builder{}, c.Data}
				r.Text.WriteString("\n")
				lastElement := hp.Body[len(hp.Body)-1]
				hp.Body = append(hp.Body[:len(hp.Body)-1], r, lastElement)
				inPreCode = true
				break
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		hp.BuildText(c, inPreCode, currentIndex)
	}

	if n.Type == html.ElementNode && n.Data == "pre" && inPreCode {
		// Find the closing </code> tag and add a newline after it
		RenderedText := RenderedText{&strings.Builder{}, n.Data}
		RenderedText.Text.WriteString("\n")
		hp.Body = append(hp.Body, RenderedText)
		inPreCode = false
	}
}

func (hp *HtmlPage) ConvertHTMLToText(htmlNode *html.Node) {
	var htmlPage HtmlPage
	currentIndex := 0
	htmlPage.BuildText(htmlNode, false, &currentIndex)
}

func ParseHtml(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	htmlNode, err := html.Parse(file)
	if err != nil {
		return "", err
	}

	var htmlPage HtmlPage
	htmlPage.ConvertHTMLToText(htmlNode)

	var str string
	for _, v := range htmlPage.Body {
		str += v.Text.String()

	}

	// trim leading and trailing whitespaces
	str = strings.TrimSpace(str)

	return str, nil
}
