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
			// args := commandSplit[1:]

			switch command {
			case CommandToggleCommander:
				exitCommandMode(c)
			case CommandScrollUp:
				// command to scroll up
			case CommandScrollDown:
				// command to scroll down
			case CommandPrevChapter:
				// command to go to previous chapter
			case CommandNextChapter:
				// command to go to next chapter
			case CommandQuit:
				quit(cancel)
			default:
				handleUnknownCommand(c, &command)
			}

		case termbox.KeyEsc:
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

func handleUnknownCommand(c *Commander, command *string) {
	c.DrawCommandBar()
	_, rows := termbox.Size()
	termbox.HideCursor()
	err := fmt.Sprintf("Command not found: %s", *command)
	for i, ch := range err {
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
