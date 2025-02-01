//go:build integration

package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kakkky/app/adapter/presentation/handler/task"
	"github.com/kakkky/app/infrastructure/api_test/testhelper"
	dbTesthelper "github.com/kakkky/app/infrastructure/db/testhelper"
	"github.com/sebdah/goldie/v2"
)

func TestTask_GetTask(t *testing.T) {
	tests := []struct {
		name      string
		pathParam string
		wantCode  int
		gfName    string
		isLogin   bool //ログインさせるか
	}{
		{
			name:      "正常系",
			wantCode:  http.StatusOK,
			pathParam: "1",
			gfName:    "get_task_nomal",
			isLogin:   true,
		},
		{
			name:      "準正常系：未ログイン",
			wantCode:  http.StatusUnauthorized,
			pathParam: "1",
			gfName:    "get_task_seminomal_not_loggedin",
			isLogin:   false,
		},
		{
			name:      "準正常系：存在しないタスク",
			pathParam: "0", //存在しないタスクのid
			wantCode:  http.StatusBadRequest,
			gfName:    "get_task_seminomal_not_found",
			isLogin:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbTesthelper.SetupFixtures("../testdata/fixtures/users.yml", "../testdata/fixtures/tasks.yml")
			// リクエストボディをマーシャル（→json）
			r := httptest.NewRequest(http.MethodGet, "/tasks/"+tt.pathParam, nil)
			rw := httptest.NewRecorder()
			// ログイン状態をセットアップ
			// Authorizationヘッダーを付加する
			if tt.isLogin {
				jwtToken := testhelper.SetupLogin("1")
				defer testhelper.CleanupLogin("1")
				r.Header.Set("Authorization", "Bearer "+jwtToken)
			}
			// リクエストを送信
			mux.ServeHTTP(rw, r)
			// ステータスコードを検証
			if rw.Code != tt.wantCode {
				t.Errorf("got %d , but want Code %d", rw.Code, tt.wantCode)
			}
			resp := testhelper.FormatJSON(
				t,
				rw.Body.Bytes(),
			)
			g := goldie.New(
				t,
				goldie.WithFixtureDir("../testdata/golden/task"),
				goldie.WithNameSuffix(".golden.json"))
			g.Assert(t, tt.gfName, resp)
		})
	}
}

func TestTask_GetTasks(t *testing.T) {
	tests := []struct {
		name     string
		wantCode int
		gfName   string
		isLogin  bool //ログインさせるか
	}{
		{
			name:     "正常系",
			wantCode: http.StatusOK,
			gfName:   "get_tasks_nomal",
			isLogin:  true,
		},
		{
			name:     "準正常系：未ログイン",
			wantCode: http.StatusUnauthorized,
			gfName:   "get_tasks_seminomal_not_loggedin",
			isLogin:  false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			dbTesthelper.SetupFixtures("../testdata/fixtures/users.yml", "../testdata/fixtures/tasks.yml")
			r := httptest.NewRequest(http.MethodGet, "/tasks", nil)
			rw := httptest.NewRecorder()
			// ログイン状態をセットアップ
			// Authorizationヘッダーを付加する
			if tt.isLogin {
				jwtToken := testhelper.SetupLogin("1")
				defer testhelper.CleanupLogin("1")
				r.Header.Set("Authorization", "Bearer "+jwtToken)
			}
			mux.ServeHTTP(rw, r)
			if rw.Code != tt.wantCode {
				t.Errorf("got %d , but want Code %d", rw.Code, tt.wantCode)
			}
			resp := testhelper.FormatJSON(
				t,
				rw.Body.Bytes(),
			)
			g := goldie.New(
				t,
				goldie.WithFixtureDir("../testdata/golden/task"),
				goldie.WithNameSuffix(".golden.json"))
			g.Assert(t, tt.gfName, resp)
		})

	}
}

func TestTask_GetUserTasks(t *testing.T) {
	tests := []struct {
		name     string
		wantCode int
		gfName   string
		isLogin  bool //ログインさせるか
	}{
		{
			name:     "正常系",
			wantCode: http.StatusOK,
			gfName:   "get_user_tasks_nomal",
			isLogin:  true,
		},
		{
			name:     "準正常系：未ログイン",
			wantCode: http.StatusUnauthorized,
			gfName:   "get_user_tasks_seminomal_not_loggedin",
			isLogin:  false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			dbTesthelper.SetupFixtures("../testdata/fixtures/users.yml", "../testdata/fixtures/tasks.yml")
			r := httptest.NewRequest(http.MethodGet, "/user/tasks", nil)
			rw := httptest.NewRecorder()
			// ログイン状態をセットアップ
			// Authorizationヘッダーを付加する
			if tt.isLogin {
				jwtToken := testhelper.SetupLogin("1")
				defer testhelper.CleanupLogin("1")
				r.Header.Set("Authorization", "Bearer "+jwtToken)
			}
			mux.ServeHTTP(rw, r)
			if rw.Code != tt.wantCode {
				t.Errorf("got %d , but want Code %d", rw.Code, tt.wantCode)
			}
			resp := testhelper.FormatJSON(
				t,
				rw.Body.Bytes(),
			)
			g := goldie.New(
				t,
				goldie.WithFixtureDir("../testdata/golden/task"),
				goldie.WithNameSuffix(".golden.json"))
			g.Assert(t, tt.gfName, resp)
		})
	}
}

func TestTask_PostTask(t *testing.T) {
	tests := []struct {
		name     string
		req      task.PostTaskRequest
		wantCode int
		gfName   string
		isLogin  bool //ログインさせるか
	}{
		{
			name:     "正常系",
			wantCode: http.StatusCreated,
			req: task.PostTaskRequest{
				Content: "content",
				State:   "todo",
			},
			gfName:  "post_task_nomal",
			isLogin: true,
		},
		{
			name:     "準正常系：未ログイン",
			wantCode: http.StatusUnauthorized,
			req: task.PostTaskRequest{
				Content: "content",
				State:   "todo",
			},
			gfName:  "post_task_seminomal_not_loggedin",
			isLogin: false,
		},
		{
			name:     "準正常系：タスクの内容が空",
			wantCode: http.StatusBadRequest,
			req: task.PostTaskRequest{
				Content: "",
				State:   "todo",
			},
			gfName:  "post_task_seminomal_content_empty",
			isLogin: true,
		},
		{
			name:     "準正常系：タスク状態の値が不正",
			wantCode: http.StatusBadRequest,
			req: task.PostTaskRequest{
				Content: "content",
				State:   "invalid",
			},
			gfName:  "post_task_seminomal_invalid_state",
			isLogin: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbTesthelper.SetupFixtures("../testdata/fixtures/users.yml", "../testdata/fixtures/tasks.yml")
			// リクエストボディをマーシャル（→json）
			b, _ := json.Marshal(tt.req)
			r := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(b))
			rw := httptest.NewRecorder()
			// ログイン状態をセットアップ
			// Authorizationヘッダーを付加する
			if tt.isLogin {
				jwtToken := testhelper.SetupLogin("1")
				defer testhelper.CleanupLogin("1")
				r.Header.Set("Authorization", "Bearer "+jwtToken)
			}
			mux.ServeHTTP(rw, r)
			if rw.Code != tt.wantCode {
				t.Errorf("got %d , but want Code %d", rw.Code, tt.wantCode)
			}
			resp := testhelper.FormatJSON(
				t,
				testhelper.NormalizeULID(rw.Body.Bytes()),
			)
			g := goldie.New(
				t,
				goldie.WithFixtureDir("../testdata/golden/task"),
				goldie.WithNameSuffix(".golden.json"))
			g.Assert(t, tt.gfName, resp)
		})
	}
}

func TestTask_UpdateTaskState(t *testing.T) {
	tests := []struct {
		name      string
		req       task.UpdateTaskStateRequest
		pathParam string
		wantCode  int
		gfName    string
		isLogin   bool //ログインさせるか
	}{
		{
			name:      "正常系",
			wantCode:  http.StatusOK,
			pathParam: "1",
			req: task.UpdateTaskStateRequest{
				State: "doing",
			},
			gfName:  "update_task_nomal",
			isLogin: true,
		},
		{
			name:      "準正常系：未ログイン",
			wantCode:  http.StatusUnauthorized,
			pathParam: "1",
			req: task.UpdateTaskStateRequest{
				State: "done",
			},
			gfName:  "update_task_seminomal_not_loggedin",
			isLogin: false,
		},
		{
			name:      "準正常系：タスク状態の値が不正",
			wantCode:  http.StatusBadRequest,
			pathParam: "1",
			req: task.UpdateTaskStateRequest{
				State: "invalid",
			},
			gfName:  "update_task_seminomal_invalid_state",
			isLogin: true,
		},
		{
			name:      "準正常系：他ユーザーのタスクは操作できない",
			wantCode:  http.StatusForbidden,
			pathParam: "3", //user2のタスク
			req: task.UpdateTaskStateRequest{
				State: "done",
			},
			gfName:  "update_task_seminomal_forbidden_operate_others_task",
			isLogin: true,
		},
		{
			name:      "準正常系：存在しないタスクはエラーが返る",
			wantCode:  http.StatusBadRequest,
			pathParam: "0", //存在しない
			req: task.UpdateTaskStateRequest{
				State: "done",
			},
			gfName:  "update_task_seminomal_not_found",
			isLogin: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbTesthelper.SetupFixtures("../testdata/fixtures/users.yml", "../testdata/fixtures/tasks.yml")
			// リクエストボディをマーシャル（→json）
			b, _ := json.Marshal(tt.req)
			r := httptest.NewRequest(http.MethodPatch, "/tasks/"+tt.pathParam+"/state", bytes.NewBuffer(b))
			rw := httptest.NewRecorder()
			// ログイン状態をセットアップ
			// Authorizationヘッダーを付加する
			if tt.isLogin {
				jwtToken := testhelper.SetupLogin("1")
				defer testhelper.CleanupLogin("1")
				r.Header.Set("Authorization", "Bearer "+jwtToken)
			}
			mux.ServeHTTP(rw, r)
			if rw.Code != tt.wantCode {
				t.Errorf("got %d , but want Code %d", rw.Code, tt.wantCode)
			}
			resp := testhelper.FormatJSON(
				t,
				rw.Body.Bytes(),
			)
			g := goldie.New(
				t,
				goldie.WithFixtureDir("../testdata/golden/task"),
				goldie.WithNameSuffix(".golden.json"))
			g.Assert(t, tt.gfName, resp)
		})
	}
}

func TestTask_DeleteTask(t *testing.T) {
	tests := []struct {
		name      string
		pathParam string
		wantCode  int
		gfName    string
		isLogin   bool //ログインさせるか
	}{
		{
			name:      "正常系",
			wantCode:  http.StatusNoContent,
			pathParam: "1",
			gfName:    "delete_task_nomal",
			isLogin:   true,
		},
		{
			name:      "準正常系：未ログイン",
			wantCode:  http.StatusUnauthorized,
			pathParam: "1",
			gfName:    "delete_task_seminomal_not_loggedin",
			isLogin:   false,
		},
		{
			name:      "準正常系：他ユーザーのタスクは操作できない",
			wantCode:  http.StatusForbidden,
			pathParam: "3",
			gfName:    "delete_task_seminomal_forbidden_operate_others_task",
			isLogin:   true,
		},
		{
			name:      "準正常系：存在しないタスクはエラーが返る",
			wantCode:  http.StatusBadRequest,
			pathParam: "0",
			gfName:    "delete_task_seminomal_not_found",
			isLogin:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbTesthelper.SetupFixtures("../testdata/fixtures/users.yml", "../testdata/fixtures/tasks.yml")

			r := httptest.NewRequest(http.MethodDelete, "/tasks/"+tt.pathParam, nil)
			rw := httptest.NewRecorder()
			// ログイン状態をセットアップ
			// Authorizationヘッダーを付加する
			if tt.isLogin {
				jwtToken := testhelper.SetupLogin("1")
				defer testhelper.CleanupLogin("1")
				r.Header.Set("Authorization", "Bearer "+jwtToken)
			}
			mux.ServeHTTP(rw, r)
			if rw.Code != tt.wantCode {
				t.Errorf("got %d , but want Code %d", rw.Code, tt.wantCode)
			}
			// 204ステータスコードなら、JSONは返らないのでここで終了
			if rw.Code == http.StatusNoContent {
				return
			}
			resp := testhelper.FormatJSON(
				t,
				rw.Body.Bytes(),
			)
			g := goldie.New(
				t,
				goldie.WithFixtureDir("../testdata/golden/task"),
				goldie.WithNameSuffix(".golden.json"))
			g.Assert(t, tt.gfName, resp)
		})
	}
}
