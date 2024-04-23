package main

import (
	"github.com/nsf/termbox-go"
	"gopuby/htmlParse"
	"gopuby/renderer"
)

// import "gopuby/epubManager"

func main() {
	// epubManager.Open("book.epub")
	// const path = "test.xhtml"

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)
	// Set the output mode, which determines how colors and attributes are used
	termbox.SetOutputMode(termbox.OutputNormal)

	// Clear the screen before writing
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	const path = "books/tmp/6c088701-fa14-4dc6-b33e-c2749d99d647/EPUB/text/ch003.xhtml"
	str, err := htmlParse.ParseHtml(path)
	if err != nil {
		panic(err)
	}

	renderer := renderer.New()
	renderer.Render(&str)

eventLoop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Ch == 'q' {
				break eventLoop
			}
			// check for arrow keys
			switch ev.Key {
			case termbox.KeyArrowUp:
				renderer.ScrollUp(&str)
			case termbox.KeyArrowDown:
				renderer.ScrollDown(&str)
			}

		case termbox.EventResize:
			renderer.Render(&str)
		}
	}

}
