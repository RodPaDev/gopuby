package main

import (
	"fmt"
	"gopuby/commander"
	"gopuby/db"
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
	db.New("gopuby.db")
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
		g.Commander.DrawHelpScreen()
		// the user then enters command to either list books or open a book from a path
	} else if len(args) == 1 {
		path := args[0]
		if err := g.Book.LoadBook(path); err != nil {
			return err
		}
		g.Renderer.Render(&g.Book.CurrentText)
	}

	g.Commander.Run()

	return nil
}

func main() {

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

}
