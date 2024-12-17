package user

import (
	"encoding/json"
	"net/http"

	"github.com/kakkky/app/adapter/presentation/presenter"
	"github.com/kakkky/app/application/usecase/user"
	"github.com/kakkky/app/domain/errors"
	"github.com/kakkky/pkg/validation"
)

type PostUserHandler struct {
	registerUsecase *user.RegisterUsecase
}

func NewPostUserHandler(registerUsecase *user.RegisterUsecase) *PostUserHandler {
	return &PostUserHandler{
		registerUsecase: registerUsecase,
	}
}

// @Summary     ユーザーの作成
// @Description 新しいユーザーを作成します。
// @Tags        User
// @Accept      json
// @Produce     json
// @Param       request body     PostUserRequest                             true "ユーザー作成のための情報"
// @Success     201     {object} presenter.SuccessResponse[PostUserResponse] "作成されたユーザーの情報"
// @Failure     400     {object} presenter.FailureResponse                   "不正なリクエスト"
// @Failure     500     {object} presenter.FailureResponse                   "内部サーバーエラー"
// @Router      /user [post]
func (puh *PostUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// jsonをデコード
	var params PostUserRequest
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		presenter.RespondBadRequest(w, err.Error())
		return
	}
	if err := validation.NewValidation().Struct(&params); err != nil {
		presenter.RespondBadRequest(w, err.Error())
		return
	}
	// DTOに詰め替える
	input := user.RegisterUsecaseInputDTO{
		Name:     params.Name,
		Email:    params.Email,
		Password: params.Password,
	}
	// ユースケースに渡して実行
	ctx := r.Context()
	output, err := puh.registerUsecase.Run(ctx, input)
	if (err != nil) && errors.IsDomainErr(err) {
		presenter.RespondBadRequest(w, err.Error())
		return
	}
	if err != nil {
		presenter.RespondInternalServerError(w, err.Error())
		return
	}
	resp := PostUserResponse{
		Email: output.Email,
		Name:  output.Name,
	}
	presenter.RespondCreated(w, resp)
}
