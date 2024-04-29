package renderer

import (
	"github.com/nsf/termbox-go"
)

type Renderer struct {
	renderBuffer string
	currentPage  *int
}

func New(page *int) *Renderer {
	r := Renderer{}
	r.currentPage = page
	return &r
}

func RenderEndText(bottomRow bool) {
	cols, rows := termbox.Size()
	str := "Use Left and Right arrow keys to navigate chapters"

	start := (cols - len(str)) / 2
	for i, char := range str {
		termbox.SetCell(start+i, rows-1, char, termbox.ColorWhite, termbox.ColorDefault)
	}

	termbox.Flush()
}

func (r *Renderer) Render(text *string) {
	cols, rows := termbox.Size()
	// Clear the screen
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	// Calculate the indices for the current page
	pageSize := rows * cols
	currentPageStart := (*r.currentPage - 1) * pageSize

	// Make sure we start within the text length
	if currentPageStart >= len(*text) {
		RenderEndText(false)
		termbox.Flush() // Nothing more to display
		return
	}

	currentPageEnd := currentPageStart + pageSize
	if currentPageEnd > len(*text) {
		currentPageEnd = len(*text)
		defer RenderEndText(true)
	}

	currCol, currRow := 0, 0

	r.renderBuffer = (*text)[currentPageStart:currentPageEnd]

	for _, char := range (*text)[currentPageStart:currentPageEnd] {
		if char == '\n' {
			currRow++
			currCol = 0
			if currRow >= rows {
				break
			}
			continue
		}

		if currCol >= cols {
			currRow++
			currCol = 0
			if currRow >= rows {
				break
			}
		}

		termbox.SetCell(currCol, currRow, char, termbox.ColorWhite, termbox.ColorDefault)
		currCol++
	}

	termbox.Flush()
}

func (r *Renderer) GetBuffer() *string {
	return &r.renderBuffer
}
