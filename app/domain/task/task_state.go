package task

import "github.com/kakkky/app/domain/errors"

type State struct {
	value int
}

var (
	todo  = State{value: 0}
	doing = State{value: 1}
	done  = State{value: 2}
)

func newState(s string) (State, error) {
	if s == "todo" {
		return todo, nil
	}
	if s == "doing" {
		return doing, nil
	}
	if s == "done" {
		return done, nil
	}
	return State{}, errors.ErrInvalidTaskState
}

func reconstructState(value int) State {
	if value == 1 {
		return doing
	}
	if value == 2 {
		return done
	}
	return todo
}

func (s State) StrValue() string {
	if s.value == 1 {
		return "doing"
	}
	if s.value == 2 {
		return "done"
	}
	return "todo"
}
func (s State) IntValue() int {
	return s.value
}
