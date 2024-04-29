package commander

import (
	"context"
	"fmt"
	"gopuby/stateMachine"
	"strings"

	"github.com/nsf/termbox-go"
)

func (c *Commander) handleCommandModeInput(ev termbox.Event, cancel context.CancelFunc) {
	if ev.Type == termbox.EventKey {
		cols, rows := GetCommandBarInputPosition()
		x, y := cols+1, rows

		if c.hasError {
			c.DrawCommandBar()
			c.hasError = false
		}

		switch ev.Key {
		case termbox.KeyEnter:
			input := c.input.ArchiveAndClear()
			if len(*input) == 0 {
				c.DrawCommandBar()
				return
			}
			commandSplit := strings.Split(*input, " ")
			command := commandSplit[0]

			switch command {
			case CommandToggleCommander:
				if c.Book.Metadata.ID == "" {
					handleNoBookLoaded(c)
					return
				}
				exitCommandMode(c)
			case CommandOpenFile:
				if len(commandSplit) != 2 {
					handleUnknownCommand(c, &command)
					return
				}
				c.Book.LoadBook(commandSplit[1])
				exitCommandMode(c)
			case CommandQuit:
				quit(cancel)
			case CommandHelp:
				c.DrawHelpScreen()
			default:
				handleUnknownCommand(c, &command)
			}

		case termbox.KeyEsc:
			if c.Book.Metadata.ID == "" {
				handleNoBookLoaded(c)
				return
			}
			exitCommandMode(c)
		case termbox.KeyArrowUp, termbox.KeyArrowDown:
			// no. you get one line and if you typo, you backspace. git gud
		case termbox.KeyArrowLeft, termbox.KeyArrowRight:
			// arrow up and down will be to traverse the command history
		case termbox.KeyBackspace, termbox.KeyBackspace2:
			handleBackspace(c, x, y)
		case termbox.KeySpace:
			insertCharacter(c, x, y, ' ')
		case termbox.KeyTab:
			insertCharacter(c, x, y, '\t')
		default:
			if ev.Ch != 0 {
				insertCharacter(c, x, y, ev.Ch)
			}
		}
	}
}

func insertCharacter(c *Commander, x, y int, ch rune) {
	c.input.AppendToBuffer(ch)
	termbox.SetCursor(x+c.input.GetRuneCount()+1, y)
	termbox.SetCell(x+c.input.GetRuneCount(), y, ch, termbox.ColorWhite, termbox.ColorDefault)
	termbox.Flush()
}

func handleBackspace(c *Commander, x, y int) {
	newRow := 0
	if c.input.GetRuneCount() > 0 {
		c.input.RemoveLastChar()
		newRow = x + c.input.GetRuneCount() + 1
	} else {
		newRow = x + 1
	}
	if newRow > 1 {
		termbox.SetCell(newRow, y, ' ', termbox.ColorWhite, termbox.ColorDefault)
		termbox.SetCursor(newRow, y)
		termbox.Flush()
	}
}

func handleNoBookLoaded(c *Commander) {
	errorMsg := fmt.Sprintf("No book loaded. Use '%s' to load a book or '%s' to list books", CommandOpenFile, CommandList)
	drawError(c, errorMsg)
}

func handleUnknownCommand(c *Commander, command *string) {
	errorMsg := fmt.Sprintf("Unknown command: %s", *command)
	drawError(c, errorMsg)
}

func drawError(c *Commander, errorMsg string) {
	c.DrawCommandBar()
	_, rows := termbox.Size()
	termbox.HideCursor()
	for i, ch := range errorMsg {
		termbox.SetCell(i+2, rows-1, ch, termbox.ColorRed, termbox.ColorDefault)
	}
	termbox.Flush()
	c.hasError = true
}

func exitCommandMode(c *Commander) {
	c.StateMachine.Transition(stateMachine.ExitCommandMode)
	c.ClearCommandBar()
	c.input.ClearBuffer()
}
