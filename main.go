package main

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
	defer termbox.Close()
	c.Renderer.Render(&str)
	c.Run()

}
