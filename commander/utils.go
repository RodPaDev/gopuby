package commander

import (
	"errors"

	"github.com/nsf/termbox-go"
)

func getCommandKey(ev termbox.Event) (string, error) {
	if ev.Ch != 0 {
		if command, ok := KeyBindings[ev.Ch]; ok {
			return command, nil
		}
	} else if command, ok := SpecialKeyBindings[ev.Key]; ok {
		return command, nil
	}

	return "", errors.New("no command found for key")
}
