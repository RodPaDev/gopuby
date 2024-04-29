package commander

import (
	"context"

	"github.com/rodpadev/gopuby/db"
	"github.com/rodpadev/gopuby/stateMachine"

	"github.com/nsf/termbox-go"
)

var normal = map[rune]string{
	// 't': CommandToggleToC,
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
		if err := c.Book.MovePage(-1); err != nil {
			panic(err)
		}
		c.Renderer.Render(&c.Book.CurrentText)
		c.updateDBItem()
	case CommandScrollDown:
		if err := c.Book.MovePage(1); err != nil {
			panic(err)
		}
		c.Renderer.Render(&c.Book.CurrentText)
		c.updateDBItem()
	case CommandPrevChapter:
		if err := c.Book.MoveChapter(-1); err != nil {
			panic(err)
		}
		c.Renderer.Render(&c.Book.CurrentText)
		c.updateDBItem()
	case CommandNextChapter:
		if err := c.Book.MoveChapter(1); err != nil {
			panic(err)
		}
		c.Renderer.Render(&c.Book.CurrentText)
		c.updateDBItem()
	case CommandQuit:
		quit(cancel)
	}
}

func (c *Commander) executeNormal(commandName string, cancel context.CancelFunc) {
	switch commandName {
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

func (c *Commander) updateDBItem() {
	db.GlobalDB.Book.CurrentPage = *c.Book.CurrentTextPage
	db.GlobalDB.Book.CurrentChapter = *c.Book.CurrentChapterIndex
	db.GlobalDB.UpdateBook(*db.GlobalDB.Book)
}
