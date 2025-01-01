package task

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kakkky/app/domain/errors"
)

func TestNewTask(t *testing.T) {
	t.Parallel()
	type args struct {
		userId  string
		content string
		state   string
	}
	tests := []struct {
		name    string
		args    args
		want    *Task
		errType error
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				userId:  "1",
				content: "This is a task",
				state:   "todo",
			},
			want: &Task{
				userId:  "1",
				content: Content{value: "This is a task"},
				state:   Todo,
			},
			wantErr: false,
		},
		{
			name: "準正常系：contentが空",
			args: args{
				content: "",
				state:   "todo",
			},
			errType: errors.ErrContentEmpty,
			wantErr: true,
		},
		{
			name: "準正常系：stateが不正",
			args: args{
				content: "This is a task",
				state:   "invalid",
			},
			errType: errors.ErrInvalidTaskState,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := NewTask(tt.args.userId, tt.args.content, tt.args.state)
			if (err != nil) != tt.wantErr && errors.Is(err, tt.errType) {
				t.Fatalf("NewTask() error=%v, but wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(got, tt.want, cmpopts.IgnoreFields(Task{}, "id"), cmp.AllowUnexported(Task{}, Content{})); diff != "" {
				t.Errorf("NewTask() -got,+want :%v", diff)
			}
		})
	}
}

func TestTask_UpdateState(t *testing.T) {
	// 初期タスクを作成
	task, _ := NewTask("1", "This is a task", "todo")

	t.Parallel()
	type args struct {
		state string
	}
	tests := []struct {
		name    string
		args    args
		want    *Task
		errType error
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				state: "done",
			},
			want: &Task{
				id:      task.id,
				userId:  task.userId,
				content: task.content,
				state:   Done,
			},
			wantErr: false,
		},
		{
			name: "準正常系：不正なstate",
			args: args{
				state: "invalid",
			},
			errType: errors.ErrInvalidTaskState,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := task.UpdateState(tt.args.state)
			if (err != nil) != tt.wantErr && errors.Is(err, tt.errType) {
				t.Fatalf("UpdateState() error=%v, wantErr=%v", err, tt.wantErr)
			}
			if err != nil && !errors.Is(err, tt.errType) {
				t.Fatalf("UpdateState() error type=%v, wantErr type=%v", err, tt.errType)
			}
			if diff := cmp.Diff(got, tt.want, cmpopts.IgnoreFields(Task{}, "id"), cmp.AllowUnexported(Task{}, Content{})); diff != "" {
				t.Errorf("UpdateState() -got,+want :%v", diff)
			}
		})
	}
}
