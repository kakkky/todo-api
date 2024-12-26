package task

import "github.com/kakkky/app/domain/errors"

type state struct {
	value int
}

var (
	todo  = state{value: 0}
	doing = state{value: 1}
	done  = state{value: 2}
)

func newState(s string) (state, error) {
	if s == "todo" {
		return todo, nil
	}
	if s == "doing" {
		return doing, nil
	}
	if s == "done" {
		return done, nil
	}
	return state{}, errors.ErrInvalidTaskState
}

func reconstructState(value int) state {
	if value == 1 {
		return doing
	}
	if value == 2 {
		return done
	}
	return todo
}

func (s state) Value() string {
	if s.value == 1 {
		return "doing"
	}
	if s.value == 2 {
		return "done"
	}
	return "todo"
}
