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
	currentPageEnd := currentPageStart + pageSize

	currCol, currRow := 0, 0

	for i := currentPageStart; i < len(*text) && i < currentPageEnd; i++ {
		char := rune((*text)[i])

		if char == '\n' {
			currRow++
			currCol = 0
			if currRow >= rows {
				break
			}
			continue
		}

		termbox.SetCell(currCol, currRow, char, termbox.ColorWhite, termbox.ColorDefault)
		currCol++

		if currCol >= cols {
			currRow++
			currCol = 0
			if currRow >= rows {
				break
			}
		}
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

func (r *Renderer) GetBuffer() string {
	return r.renderBuffer
}
