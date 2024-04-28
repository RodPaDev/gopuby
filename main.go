package main

import (
	"fmt"
	"gopuby/commander"
	"gopuby/epubManager"
	"gopuby/renderer"
	"gopuby/stateMachine"
	"os"

	"github.com/nsf/termbox-go"
	// "github.com/nsf/termbox-go"
)

type Gopuby struct {
	Book         *epubManager.Book
	Commander    *commander.Commander
	Renderer     *renderer.Renderer
	StateMachine *stateMachine.StateMachine
}

func New() *Gopuby {
	initialPage := 1
	currentChapterIndex := 0
	b := &epubManager.Book{
		CurrentTextPage:     &initialPage,
		CurrentChapterIndex: &currentChapterIndex,
	}
	sm := stateMachine.New()
	r := renderer.New(&initialPage)
	c := commander.New(r, sm, b)

	return &Gopuby{
		Commander:    c,
		Renderer:     r,
		StateMachine: sm,
		Book:         b,
	}
}

func run() error {
	g := New()
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)
	termbox.SetOutputMode(termbox.OutputNormal)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()
	args := os.Args[1:]

	if len(args) > 1 {
		return fmt.Errorf("\nUsage: gopuby [path]")
	} else if len(args) == 0 {
		// transition state to commander
		g.StateMachine.Transition(stateMachine.EnterCommandMode)
		g.Commander.DrawCommandBar()
		// the user then enters command to either list books or open a book from a path
	} else if len(args) == 1 {
		path := args[0]
		if err := g.Book.LoadBook(path); err != nil {
			return err
		}
	}

	g.Renderer.Render(&g.Book.CurrentText)

	g.Commander.Run()

	return nil
}

func main() {

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// const path = "books/tmp/6c088701-fa14-4dc6-b33e-c2749d99d647/EPUB/text/ch003.xhtml"
	// str, err := htmlParse.ParseHtml(path)
	// if err != nil {
	// 	panic(err)
	// }
	// g.Commander.ParsedText = str
	// defer termbox.Close()
	// g.Commander.Renderer.Render(&g.Commander.ParsedText)

	// g.Commander.Run()

}

// 	c := commander.New()
// 	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

// 	const path = "books/tmp/6c088701-fa14-4dc6-b33e-c2749d99d647/EPUB/text/ch003.xhtml"
// 	str, err := htmlParse.ParseHtml(path)
// 	if err != nil {
// 		panic(err)
// 	}
// 	c.ParsedText = str
// 	defer termbox.Close()
// 	c.Renderer.Render(&c.ParsedText)

// 	c.Run()
// }

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
