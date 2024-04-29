package commander

import (
	"context"
	"time"

	"github.com/rodpadev/gopuby/epubManager"
	"github.com/rodpadev/gopuby/input"
	"github.com/rodpadev/gopuby/renderer"
	"github.com/rodpadev/gopuby/stateMachine"
	"github.com/rodpadev/gopuby/utils"

	"github.com/nsf/termbox-go"
)

const (
	CommandHelp            = "help"
	CommandToggleCommander = "toggleCommander"
	CommandOpenFile        = "openFile"
	// CommandList            = "list"
	// CommandRemove          = "remove"
	CommandQuit        = "quit"
	CommandScrollUp    = "scrollUp"
	CommandScrollDown  = "scrollDown"
	CommandNextChapter = "nextChapter"
	CommandPrevChapter = "prevChapter"
	// CommandToggleToC       = "toggleToC"
	CommandFind = "find"
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
