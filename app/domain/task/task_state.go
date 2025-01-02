package task

import "github.com/kakkky/app/domain/errors"

type State int

const (
	Todo State = iota
	Doing
	Done
)

func newState(value string) (State, error) {
	switch value {
	case "todo":
		return Todo, nil
	case "doing":
		return Doing, nil
	case "done":
		return Done, nil
	default:
		return 0, errors.ErrInvalidTaskState
	}
}

func reconstructState(value int) State {
	switch value {
	case 1:
		return Doing
	case 2:
		return Done
	default:
		return Todo
	}
}

func (s State) StrValue() string {
	switch s {
	case Doing:
		return "doing"
	case Done:
		return "done"
	default:
		return "todo"
	}
}

func (s State) IntValue() int {
	return int(s)
}
