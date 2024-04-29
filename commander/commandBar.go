package commander

import (
	"bytes"
	"fmt"

	"github.com/nsf/termbox-go"
)

func formatCommandDescriptionParams(command, params, description string) string {
	return fmt.Sprintf("- %s %s: %s", command, params, description)
}

func formatCommandDescriptionNoParams(command, description string) string {
	return fmt.Sprintf("- %s: %s", command, description)
}

var commandDescriptions = []string{
	formatCommandDescriptionNoParams(CommandHelp, "Show this help screen"),
	formatCommandDescriptionNoParams(CommandToggleCommander, "Toggle the command bar"),
	formatCommandDescriptionParams(CommandOpenFile, "<path>", "Open a file (only .epub is supported)"),
	formatCommandDescriptionNoParams(CommandList, "Opens interactive list of books"),
	formatCommandDescriptionParams(CommandRemove, "<book name>", "Remove a book from the library"),
	formatCommandDescriptionNoParams(CommandQuit, "Quit the application"),
	formatCommandDescriptionNoParams(CommandToggleToC, "Toggle the Table of Contents"),
	formatCommandDescriptionParams(CommandFind, "<search term>", "Search for a term in the book"),
}

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

func (c *Commander) DrawHelpScreen() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	c.DrawCommandBar()
	cols, rows := termbox.Size()

	lines := []string{
		"Welcome to Gopuby!",
		"You can start by loading a book with the 'openFile' command.",
		"",
		"Here are some commands you can use:",
	}
	lines = append(lines, commandDescriptions...)
	lines = append(lines, "", fmt.Sprintf("Use the '%s' command to quit Gopuby.", CommandQuit))

	max := 0
	for _, line := range lines {
		if len(line) > max {
			max = len(line)
		}
	}

	// draw at the center of the screen
	x := cols/2 - max/2
	y := rows / 2

	for i, line := range lines {
		for j, ch := range line {
			termbox.SetCell(x+j, y+i, ch, termbox.ColorWhite, termbox.ColorDefault)
		}
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
