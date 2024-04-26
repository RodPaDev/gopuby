package stateMachine

type State uint

const (
	ReadingMode State = iota
	CommandMode
	ModalMode
)

type Event uint

const (
	EnterCommandMode Event = iota
	ExitCommandMode
	EnterModalMode
	ExitModalMode
)

type Transition struct {
	From  State
	Event Event
}

type StateMachine struct {
	CurrentState State
	Transitions  map[Transition]State
}

func New() *StateMachine {
	s := StateMachine{
		CurrentState: ReadingMode,
		Transitions: map[Transition]State{
			// Reading -> Command | Command -> Reading
			{ReadingMode, EnterCommandMode}: CommandMode,
			{CommandMode, ExitCommandMode}:  ReadingMode,
			// Reading -> Modal | Modal -> Reading
			{ReadingMode, EnterModalMode}: ModalMode,
			{ModalMode, ExitModalMode}:    ReadingMode,
			// Command -> Modal (Modal always exits to Reading mode for simplicity)
			{CommandMode, EnterModalMode}: ModalMode,
		},
	}
	return &s
}

func (s *StateMachine) Transition(event Event) {
	currentTransition := Transition{s.CurrentState, event}
	if newState, ok := s.Transitions[currentTransition]; ok {
		s.CurrentState = newState
	}
}


