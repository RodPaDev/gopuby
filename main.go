package main

import "gopuby/screen"

// import "gopuby/epubManager"

func main() {
	// epubManager.Open("book.epub")
	// const path = "test.xhtml"
	// const path = "books/tmp/6c088701-fa14-4dc6-b33e-c2749d99d647/EPUB/text/ch003.xhtml"
	// str, err := htmlParse.ParseHtml(path)
	screen := screen.New()
	screen.StartResizeWatcher()

}
