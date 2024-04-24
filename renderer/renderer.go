package renderer

import (
	"github.com/nsf/termbox-go"
)

type Renderer struct {
	renderBuffer string
	currentPage  int
}

func New() *Renderer {
	return &Renderer{"", 1}
}

func (r *Renderer) Render(text *string) {
	cols, rows := termbox.Size()
	// Clear the screen
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	// Calculate the indices for the current page
	pageSize := rows * cols
	currentPageStart := (r.currentPage - 1) * pageSize

	// Make sure we start within the text length
	if currentPageStart >= len(*text) {
		termbox.Flush() // Nothing more to display
		return
	}

	currentPageEnd := currentPageStart + pageSize
	if currentPageEnd > len(*text) {
		currentPageEnd = len(*text)
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

func (r *Renderer) ScrollUp(text *string) {
	r.currentPage--
	if r.currentPage < 1 {
		r.currentPage = 1
	}
	r.Render(text)
}

func (r *Renderer) ScrollDown(text *string) {
	r.currentPage++
	r.Render(text)
}

func (r *Renderer) GetBuffer() *string {
    return &r.renderBuffer
}
