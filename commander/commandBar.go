package commander

import (
	"github.com/nsf/termbox-go"
)

// this is the command bar, it is the main entry point for all commands
// when pressing space the command bar will be shown and the user can enter a command
// the command bar will always be two lines height
// the top most line will be the status bar
// the bottom line will be the input line

// the status line will show the current chapter and section
// for now will only do the input line

// to do the input line I use termbox to draw a line over the last row of the screen
// and then I will draw a ">" character to show the user that he can enter a command
// users press space to be to show the command bar and press space again to hide it

// first lets draw the command bar with termbox
// we go to the last row of the terminal and draw a top line

func (c *Commander) DrawCommandBar() {
	cols, rows := termbox.Size()
	// draw the top line
	for i := 0; i < cols; i++ {
		termbox.SetCell(i, rows-2, '\u2500', termbox.ColorWhite, termbox.ColorDefault)
	}
	for i := 0; i < cols; i++ {
		if i == 0 {
			termbox.SetCell(i, rows-1, '>', termbox.ColorWhite, termbox.ColorDefault)
			continue
		}
		termbox.SetCell(i, rows-1, ' ', termbox.ColorWhite, termbox.ColorDefault)
	}
	termbox.Flush()
}

func (c *Commander) ClearCommandBar() {
	cols, rows := termbox.Size()

	for y := rows - 2; y < rows; y++ {
		for x := 0; x < cols; x++ {
			termbox.SetCell(x, y, ' ', termbox.ColorWhite, termbox.ColorDefault)
		}
	}

	for i := 0; i < cols; i++ {
		termbox.SetCell(i, rows-2, ' ', termbox.ColorWhite, termbox.ColorDefault)
	}
	c.Renderer.Render(&c.ParsedText)
	termbox.Flush()

}

func (c *Commander) ToggleCommandBar() {
	c.isOpen = !c.isOpen
	if c.isOpen {
		c.DrawCommandBar()
	} else {
		c.ClearCommandBar()
	}
}
