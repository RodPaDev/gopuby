package commander

import (
	"context"
	"fmt"
	"gopuby/stateMachine"
)

type Command interface {
	Execute(args ...interface{})
}

type ToggleCommander struct {
	Commander *Commander
}

func (cmd *ToggleCommander) Execute(args ...interface{}) {
	print("ToggleCommander")
	if cmd.Commander.StateMachine.CurrentState != stateMachine.CommandMode {
		cmd.Commander.StateMachine.Transition(stateMachine.EnterCommandMode)
		cmd.Commander.DrawCommandBar()
	} else {
		cmd.Commander.StateMachine.Transition(stateMachine.ExitCommandMode)
		cmd.Commander.ClearCommandBar()
	}

}

type OpenFile struct {
	Commander *Commander
}

func (cmd *OpenFile) Execute(args ...interface{}) {
	fmt.Println("OpenFile")
}

type List struct {
	Commander *Commander
}

func (cmd *List) Execute(args ...interface{}) {
	fmt.Println("List")
}

type Remove struct {
	Commander *Commander
}

func (cmd *Remove) Execute(args ...interface{}) {
	fmt.Println("Remove")
}

type Quit struct {
	Commander *Commander
}

func (cmd *Quit) Execute(args ...interface{}) {
	cancel := args[len(args)-1].(context.CancelFunc)
	cancel()
}

type ScrollUp struct {
	Commander *Commander
}

func (cmd *ScrollUp) Execute(args ...interface{}) {
	cmd.Commander.Renderer.ScrollUp(&cmd.Commander.ParsedText)
}

type ScrollDown struct {
	Commander *Commander
}

func (cmd *ScrollDown) Execute(args ...interface{}) {
	cmd.Commander.Renderer.ScrollDown(&cmd.Commander.ParsedText)
}

type NextSection struct {
	Commander *Commander
}

func (cmd *NextSection) Execute(args ...interface{}) {
	fmt.Println("NextSection")
}

type PrevSection struct {
	Commander *Commander
}

func (cmd *PrevSection) Execute(args ...interface{}) {
	fmt.Println("PrevSection")
}

type NextChapter struct {
	Commander *Commander
}

func (cmd *NextChapter) Execute(args ...interface{}) {
	fmt.Println("NextChapter")
}

type PrevChapter struct {
	Commander *Commander
}

func (cmd *PrevChapter) Execute(args ...interface{}) {
	fmt.Println("PrevChapter")
}

type ToggleToC struct {
	Commander *Commander
}

func (cmd *ToggleToC) Execute(args ...interface{}) {
	fmt.Println("ToggleToC")
}

type ToggleMarkRead struct {
	Commander *Commander
}

func (cmd *ToggleMarkRead) Execute(args ...interface{}) {
	fmt.Println("ToggleMarkRead")
}

type JumpToSection struct {
	Commander *Commander
}

func (cmd *JumpToSection) Execute(args ...interface{}) {
	fmt.Println("JumpToSection")
}

type Find struct {
	Commander *Commander
}

func (cmd *Find) Execute(args ...interface{}) {
	fmt.Println("Find")
}

type FindChapter struct {
	Commander *Commander
}

func (cmd *FindChapter) Execute(args ...interface{}) {
	fmt.Println("FindChapter")
}
