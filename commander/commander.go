package commander

import (
	"context"
	"fmt"
	"gopuby/input"
	"gopuby/renderer"
	"gopuby/stateMachine"
	"strings"

	"github.com/nsf/termbox-go"
)

const (
	CommandToggleCommander = "toggleCommander"
	CommandOpenFile        = "openFile"
	CommandList            = "list"
	CommandRemove          = "remove"
	CommandQuit            = "quit"
	CommandScrollUp        = "scrollUp"
	CommandScrollDown      = "scrollDown"
	CommandNextSection     = "nextSection"
	CommandPrevSection     = "prevSection"
	CommandNextChapter     = "nextChapter"
	CommandPrevChapter     = "prevChapter"
	CommandToggleToC       = "toggleToC"
	CommandToggleMarkRead  = "toggleMarkRead"
	CommandJumpToSection   = "jumpToSection"
	CommandFind            = "find"
	CommandFindChapter     = "findChapter"
)

var SpecialKeyBindings = map[termbox.Key]string{
	termbox.KeySpace:      CommandToggleCommander,
	termbox.KeyArrowUp:    CommandScrollUp,
	termbox.KeyArrowDown:  CommandScrollDown,
	termbox.KeyArrowLeft:  CommandPrevSection,
	termbox.KeyArrowRight: CommandNextSection,
	termbox.KeyPgup:       CommandNextChapter,
	termbox.KeyPgdn:       CommandPrevChapter,
}

var KeyBindings = map[rune]string{
	't': CommandToggleToC,
	'q': CommandQuit,
	'r': CommandToggleMarkRead,
	's': CommandJumpToSection,
	'/': CommandFind,
}

type Commander struct {
	Renderer     renderer.Renderer
	ParsedText   string
	StateMachine stateMachine.StateMachine
	Commands     map[string]Command

	input    input.Input
	hasError bool
}

func New() *Commander {
	c := &Commander{}
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	termbox.SetInputMode(termbox.InputEsc)
	termbox.SetOutputMode(termbox.OutputNormal)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	c.Renderer = *renderer.New()
	c.StateMachine = *stateMachine.New()
	c.Commands = map[string]Command{}
	c.mapCommands()
	return c
}

func (c *Commander) mapCommands() {
	c.Commands[CommandToggleCommander] = &ToggleCommander{Commander: c}
	c.Commands[CommandOpenFile] = &OpenFile{Commander: c}
	c.Commands[CommandList] = &List{Commander: c}
	c.Commands[CommandRemove] = &Remove{Commander: c}
	c.Commands[CommandQuit] = &Quit{Commander: c}
	c.Commands[CommandScrollUp] = &ScrollUp{Commander: c}
	c.Commands[CommandScrollDown] = &ScrollDown{Commander: c}
	c.Commands[CommandNextSection] = &NextSection{Commander: c}
	c.Commands[CommandPrevSection] = &PrevSection{Commander: c}
	c.Commands[CommandNextChapter] = &NextChapter{Commander: c}
	c.Commands[CommandPrevChapter] = &PrevChapter{Commander: c}
	c.Commands[CommandToggleToC] = &ToggleToC{Commander: c}
	c.Commands[CommandToggleMarkRead] = &ToggleMarkRead{Commander: c}
	c.Commands[CommandJumpToSection] = &JumpToSection{Commander: c}
	c.Commands[CommandFind] = &Find{Commander: c}
}

func (c *Commander) handleReadingModeInput(ev termbox.Event, cancel context.CancelFunc) {

	if ev.Type == termbox.EventKey {
		commandKey, err := getCommandKey(ev)
		if err != nil {
			return
		}
		c.executeCommand(commandKey, cancel)
	}
}

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
			// split input into command and args
			args := []interface{}{}

			args = append(args, strings.Split(*input, " "), cancel)

			termbox.Flush()
			c.executeCommand(*input, args...)
		case termbox.KeyEsc:
			c.executeCommand(CommandToggleCommander)
		case termbox.KeyArrowUp, termbox.KeyArrowDown, termbox.KeyArrowLeft, termbox.KeyArrowRight:
			// no. you get one line and if you typo, you backspace. git gud
		case termbox.KeyBackspace, termbox.KeyBackspace2:
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
		case termbox.KeySpace:
			c.input.AppendToBuffer(' ')
			termbox.SetCursor(x+c.input.GetRuneCount()+1, y)
			termbox.SetCell(x+c.input.GetRuneCount(), y, ev.Ch, termbox.ColorWhite, termbox.ColorDefault)
			termbox.Flush()
		case termbox.KeyTab:
			c.input.AppendToBuffer('\t')
			termbox.SetCursor(x+c.input.GetRuneCount()+1, y)
			termbox.SetCell(x+c.input.GetRuneCount(), y, ev.Ch, termbox.ColorWhite, termbox.ColorDefault)
			termbox.Flush()
		default:
			if ev.Ch != 0 {
				// write the character to the screen
				c.input.AppendToBuffer(ev.Ch)
				termbox.SetCursor(x+c.input.GetRuneCount()+1, y)
				termbox.SetCell(x+c.input.GetRuneCount(), y, ev.Ch, termbox.ColorWhite, termbox.ColorDefault)
				termbox.Flush()
			}
		}
	}
}

func (c *Commander) handleModalModeInput(ev termbox.Event, cancel context.CancelFunc) {
	if ev.Type == termbox.EventKey {
		commandKey, err := getCommandKey(ev)
		if err != nil {
			// Handle the case where no command is found for the key
			return
		}
		c.executeCommand(commandKey, cancel)
	}
}

func (c *Commander) executeCommand(commandKey string, args ...interface{}) {
	if command, ok := c.Commands[commandKey]; ok {
		command.Execute(args...)
	} else {
		c.DrawCommandBar()
		_, rows := termbox.Size()
		termbox.HideCursor()
		err := fmt.Sprintf("Command not found: %s", commandKey)
		for i, ch := range err {
			termbox.SetCell(i+2, rows-1, ch, termbox.ColorRed, termbox.ColorDefault)
		}
		termbox.Flush()
		c.hasError = true
	}
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
			case stateMachine.ModalMode:
				c.handleModalModeInput(ev, cancel)
			}
		}
	}
}
