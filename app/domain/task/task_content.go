package task

import (
	"unicode/utf8"

	"github.com/kakkky/app/domain/errors"
)

type Content struct {
	value string
}

func newContent(value string) (Content, error) {
	if utf8.RuneCountInString(value) == 0 {
		return Content{}, errors.ErrContentEmpty
	}
	return Content{value: value}, nil
}

func reconstructContent(value string) Content {
	return Content{value: value}
}

func (c Content) Value() string {
	return c.value
}
