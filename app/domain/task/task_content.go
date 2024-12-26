package task

import (
	"unicode/utf8"

	"github.com/kakkky/app/domain/errors"
)

type content struct {
	value string
}

func newContent(value string) (content, error) {
	if utf8.RuneCountInString(value) == 0 {
		return content{}, errors.ErrContentEmpty
	}
	return content{value: value}, nil
}

func reconstructContent(value string) content {
	return content{value: value}
}

func (c content) Value() string {
	return c.value
}
