package screen

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/term"
)

type Screen struct {
	rows int
	cols int
}

func New() *Screen {
	fd := os.Stdin.Fd()
	rows, cols, err := term.GetSize(int(fd))
	if err != nil {
		panic(err)
	}

	return &Screen{rows, cols}
}

func (s *Screen) UpdateSize() {
	fd := os.Stdin.Fd()
	rows, cols, err := term.GetSize(int(fd))
	if err != nil {
		return
	}
	s.rows = rows
	s.cols = cols

	fmt.Printf("Rows: %d, Cols: %d\n", s.rows, s.cols)
}

func (s *Screen) StartResizeWatcher() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGWINCH)

	go func() {
		for {
			sig := <-signals
			if sig == syscall.SIGWINCH {
				s.UpdateSize()
			}
		}
	}()

	// Prevent the main goroutine from exiting
	select {}
}
