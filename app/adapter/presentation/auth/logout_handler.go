package auth

import (
	"net/http"

	"github.com/kakkky/app/adapter/presentation/middleware"
	"github.com/kakkky/app/adapter/presentation/presenter"
	"github.com/kakkky/app/application/usecase/auth"
)

type LogoutHandler struct {
	logoutUsecase *auth.LogoutUsecase
}

func NewLogoutHandler(logoutUsecase *auth.LogoutUsecase) *LogoutHandler {
	return &LogoutHandler{
		logoutUsecase: logoutUsecase,
	}
}

// @Summary     ユーザーのログアウト
// @Description メールアドレス・パスワードで認証し、署名されたトークンを返す
// @Tags        User/Auth
// @Produce     json
// @Security    BearerAuth
// @Success     204
// @Failure     500 {object} presenter.FailureResponse "内部サーバーエラー"
// @Router      /logout [delete]
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
