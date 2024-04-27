package input

import (
	"bytes"
	"sync"
)

type InputBuffer struct {
	runeCount int
	buffer    bytes.Buffer
	lock      sync.Mutex
}

type Input struct {
	inputBuffer InputBuffer
	history     []string
}

func (i *Input) AppendToBuffer(ch rune) {
	i.inputBuffer.lock.Lock()
	defer i.inputBuffer.lock.Unlock()
	i.inputBuffer.buffer.WriteRune(ch)
	i.inputBuffer.runeCount += 1
}

func (i *Input) RemoveLastChar() {
	i.inputBuffer.lock.Lock()
	defer i.inputBuffer.lock.Unlock()
	buf := i.inputBuffer.buffer.Bytes()
	if len(buf) > 0 {
		runeStart := len(buf) - 1
		for runeStart > 0 && (buf[runeStart]&0xC0) == 0x80 {
			runeStart -= 1
		}
		i.inputBuffer.buffer.Truncate(runeStart)
	}
	i.inputBuffer.runeCount -= 1
	if i.inputBuffer.runeCount < 0 {
		i.inputBuffer.runeCount = 0
	}
}

func (i *Input) ArchiveAndClear() *string {
	i.inputBuffer.lock.Lock()
	defer i.inputBuffer.lock.Unlock()
	input := i.inputBuffer.buffer.String()
	i.history = append(i.history, input)
	i.inputBuffer.buffer.Reset()
	i.inputBuffer.runeCount = 0
	return &input
}

func (i *Input) GetInputAtPos(pos int) string {
	if pos < 0 || pos > len(i.history) {
		return ""
	}
	return i.history[pos]
}

func (i *Input) GetRuneCount() int {
	return i.inputBuffer.runeCount
}
