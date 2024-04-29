package commander

import (
	"context"
	"gopuby/epubManager"
	"gopuby/input"
	"gopuby/renderer"
	"gopuby/stateMachine"
	"gopuby/utils"
	"time"

	"github.com/nsf/termbox-go"
)

const (
	CommandToggleCommander = "toggleCommander"
	CommandOpenFile        = "openFile"    // commander only
	CommandList            = "list"        // commander only
	CommandRemove          = "remove"      // commander only
	CommandQuit            = "quit"        // commander, reading
	CommandScrollUp        = "scrollUp"    // commander, reading, modal
	CommandScrollDown      = "scrollDown"  // commander, reading, modal
	CommandNextChapter     = "nextChapter" // commander, reading
	CommandPrevChapter     = "prevChapter" // commander, reading
	CommandToggleToC       = "toggleToC"   // commander, reading, modal
	CommandFind            = "find"        // commander, possibly a search view in modal
)

type Commander struct {
	// Global
	Renderer     *renderer.Renderer
	StateMachine *stateMachine.StateMachine
	Book         *epubManager.Book

	// Local
	input                 *input.Input
	hasError              bool
	debouncedUpdateDBItem func()
}

func New(
	renderer *renderer.Renderer,
	stateMachine *stateMachine.StateMachine,
	book *epubManager.Book,
) *Commander {
	c := &Commander{
		Renderer:     renderer,
		StateMachine: stateMachine,
		Book:         book,
		input:        &input.Input{},
	}
	c.debouncedUpdateDBItem = utils.Debounce(c.updateDBItem, 500*time.Millisecond)
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	termbox.SetInputMode(termbox.InputEsc)
	termbox.SetOutputMode(termbox.OutputNormal)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	return c
}

func (c *Commander) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer termbox.Close()
	defer cancel()

	// this is cool, it decouples the event polling from the event handling
	eventChan := make(chan termbox.Event)
	go func() {
		for {
			eventChan <- termbox.PollEvent()
		}
	}()

eventLoop:
	for {
		select {
		case <-ctx.Done():
			break eventLoop
		case ev := <-eventChan:
			switch c.StateMachine.CurrentState {
			case stateMachine.ReadingMode:
				c.handleReadingModeInput(ev, cancel)
			case stateMachine.CommandMode:
				c.handleCommandModeInput(ev, cancel)
				// case stateMachine.ModalMode:
				// 	c.handleModalModeInput(ev, cancel)
			}
		}
	}
}
