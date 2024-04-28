package commander

import (
	"context"
	"gopuby/stateMachine"

	"github.com/nsf/termbox-go"
)

var normal = map[rune]string{
	't': CommandToggleToC,
	'q': CommandQuit,
}

var special = map[termbox.Key]string{
	termbox.KeyEsc:        CommandQuit,
	termbox.KeySpace:      CommandToggleCommander,
	termbox.KeyArrowUp:    CommandScrollUp,
	termbox.KeyArrowDown:  CommandScrollDown,
	termbox.KeyArrowLeft:  CommandPrevChapter,
	termbox.KeyArrowRight: CommandNextChapter,
}

func (c *Commander) executeSpecial(commandName string, cancel context.CancelFunc) {
	switch commandName {
	case CommandToggleCommander:
		c.StateMachine.Transition(stateMachine.EnterCommandMode)
		c.DrawCommandBar()
	case CommandScrollUp:
		c.Renderer.ScrollUp(&c.Book.CurrentText)
	case CommandScrollDown:
		c.Renderer.ScrollDown(&c.Book.CurrentText)
	case CommandPrevChapter:
		if err := c.Book.MoveChapter(-1); err != nil {
			panic(err)
		}
		c.Renderer.Render(&c.Book.CurrentText)
	case CommandNextChapter:
		if err := c.Book.MoveChapter(1); err != nil {
			panic(err)
		}
		c.Renderer.Render(&c.Book.CurrentText)
	case CommandQuit:
		quit(cancel)
	}
}

func (c *Commander) executeNormal(commandName string, cancel context.CancelFunc) {
	switch commandName {
	case CommandToggleToC:
		// toggleTOC(cancel)
	case CommandQuit:
		quit(cancel)

	}
}

func (c *Commander) handleReadingModeInput(ev termbox.Event, cancel context.CancelFunc) {
	if ev.Type != termbox.EventKey {
		return
	}
	if commandName, exists := special[ev.Key]; exists {
		c.executeSpecial(commandName, cancel)
	} else if commandName, exists := normal[ev.Ch]; exists {
		c.executeNormal(commandName, cancel)
	}
}

func quit(cancel context.CancelFunc) {
	cancel()
}
