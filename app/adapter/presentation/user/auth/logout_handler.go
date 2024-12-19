package auth

import (
	"net/http"

	"github.com/kakkky/app/adapter/presentation/middleware"
	"github.com/kakkky/app/adapter/presentation/presenter"
	"github.com/kakkky/app/application/usecase/user/auth"
)

type LogoutHandler struct {
	logoutUsecase *auth.LogoutUsecase
}

func NewLogoutHandler(logoutUsecase *auth.LogoutUsecase) *LogoutHandler {
	return &LogoutHandler{
		logoutUsecase: logoutUsecase,
	}
}

func (lh *LogoutHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// 認可制御のミドルウェアでコンテキストに値が付加されている
	// リクエストスコープのコンテキストからuserIDを取得する
	userID := r.Context().Value(middleware.UserIDKey{}).(string)
	ctx := r.Context()
	if err := lh.logoutUsecase.Run(ctx, auth.LogoutUsecaseInputDTO{ID: userID}); err != nil {
		presenter.RespondInternalServerError(rw, err.Error())
	}
	presenter.RespondNoContent(rw)
}
