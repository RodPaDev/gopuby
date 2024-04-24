package main

// import "gopuby/epubManager"

import (
	"gopuby/commander"
	"gopuby/htmlParse"

	"github.com/nsf/termbox-go"
)

func main() {

	c := commander.New()
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	const path = "books/tmp/6c088701-fa14-4dc6-b33e-c2749d99d647/EPUB/text/ch003.xhtml"
	str, err := htmlParse.ParseHtml(path)
	if err != nil {
		panic(err)
	}
	c.ParsedText = str
	defer termbox.Close()
	c.Renderer.Render(&c.ParsedText)

	c.Run()
}

// func main() {
// 	path := "books/tmp/6c088701-fa14-4dc6-b33e-c2749d99d647"
// 	e, _ := epubManager.GetMetadata(path)
// 	t, _ := epubManager.GetTableOfContents(path)

// 	println("Title:", e.Title)
// 	println("Author:", e.Author)
// 	println("Language:", e.Lang)
// 	println("Date:", e.Date)
// 	println("ID:", e.ID)

// }
