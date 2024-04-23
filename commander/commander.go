package commander

import (
	"github.com/nsf/termbox-go"
	"gopuby/renderer"
)

const (
	Commmander     = ":commander"
	Open           = ":open"
	List           = ":list"
	Remove         = ":remove"
	Quit           = ":quit"
	ScrollUp       = ":scrollUp"
	ScrollDown     = ":scrollDown"
	NextSection    = ":nextSection"
	PrevSection    = ":prevSection"
	NextChapter    = ":nextChapter"
	PrevChapter    = ":prevChapter"
	ToggleToC      = ":toggleToC"
	ScrollUpToC    = ":scrollUpToC"
	ScrollDownToC  = ":scrollDownToC"
	ToggleMarkRead = ":toggleMarkRead"
	JumpToSection  = ":jumpToSection"
	Find           = ":find"
	FindChapter    = ":findChapter"
)

var SpecialKeyBindings = map[termbox.Key]string{
	termbox.KeySpace:      Commmander,
	termbox.KeyArrowUp:    ScrollUp,
	termbox.KeyArrowDown:  ScrollDown,
	termbox.KeyArrowLeft:  PrevSection,
	termbox.KeyArrowRight: NextSection,
	termbox.KeyPgup:       NextChapter,
	termbox.KeyPgdn:       PrevChapter,
	't':                   ToggleToC,
	'q':                   Quit,
	'r':                   ToggleMarkRead,
	's':                   JumpToSection,
	'/':                   Find,
}

var KeyBindings = map[rune]string{
	't': ToggleToC,
	'q': Quit,
	'r': ToggleMarkRead,
	's': JumpToSection,
	'/': Find,
}

type Commander struct {
	Renderer renderer.Renderer
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
	return c
}

func (c *Commander) executeCommand(command string, arg string) {
	switch command {
	case ScrollUp:
		// c.renderer.ScrollUp(&arg)
	case ScrollDown:
		// c.renderer.ScrollDown(&arg)

	}
}

func (c *Commander) Run() {

eventLoop:
	for {
		ev := termbox.PollEvent()
		if ev.Type == termbox.EventKey {
			if ev.Ch != 0 { // Handle character keys
				if command, ok := KeyBindings[ev.Ch]; ok {
					if command == Quit {
						break eventLoop
					}
					c.executeCommand(command, "")
				}
			} else if command, ok := SpecialKeyBindings[ev.Key]; ok { // Handle special keys
				if command == Quit {
					break eventLoop
				}
				c.executeCommand(command, "")
			}

		}
	}
}
