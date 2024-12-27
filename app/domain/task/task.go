package task

import (
	"github.com/kakkky/app/domain/errors"
	"github.com/kakkky/pkg/ulid"
)

type Task struct {
	id      string
	userId  string
	content Content
	state   State
}

func NewTask(
	userId string,
	content string,
	state string,
) (*Task, error) {
	validatedContent, err := newContent(content)
	if err != nil {
		return nil, err
	}
	validatedState, err := newState(state)
	if err != nil {
		return nil, err
	}
	return &Task{
		id:      ulid.NewUlid(),
		userId:  userId,
		content: validatedContent,
		state:   validatedState,
	}, nil
}

// リポジトリから使用
func ReconstructTask(
	id string,
	userId string,
	content string,
	state int, //DBにはint型でタスク状態を保存している
) *Task {
	return &Task{
		id:      id,
		userId:  userId,
		content: reconstructContent(content),
		state:   reconstructState(state),
	}
}

func (t *Task) UpdateState(
	state string,
) (*Task, error) {
	validatedState, err := newState(state)
	if err != nil {
		return nil, err
	}
	return &Task{
		id:      t.id,
		userId:  t.userId,
		content: t.content,
		state:   validatedState,
	}, nil
}

func (t *Task) IsOperableBy(userId string) error {
	if t.userId != userId {
		return errors.ErrForbiddenTaskOperation
	}
	return nil
}

func (t *Task) GetID() string {
	return t.id
}
func (t *Task) GetUserId() string {
	return t.userId
}
func (t *Task) GetContent() Content {
	return t.content
}
func (t *Task) GetState() State {
	return t.state
}
