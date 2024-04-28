package commander

import (
	"bytes"
	"github.com/nsf/termbox-go"
)

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

	x, y := 2, rows-1
	termbox.SetCursor(x+c.input.GetRuneCount(), y)
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
	c.Renderer.Render(&c.Book.CurrentText)
	termbox.SetCursor(0, 0)
	termbox.HideCursor()
	termbox.Flush()
}

func GetCommandBarInputPosition() (int, int) {
	_, rows := termbox.Size()
	return 0, rows - 1
}

func DrawBuffer(buf *bytes.Buffer, x, y int) {
	str := buf.String()
	for i, ch := range str {
		termbox.SetCell(x+i, y, rune(ch), termbox.ColorWhite, termbox.ColorDefault)
	}
	termbox.Flush()
}
