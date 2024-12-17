package health

import (
	"net/http"

	"github.com/kakkky/app/adapter/presentation/presenter"
)

// HealthCheckHandler godoc
// @Summary     apiのヘルスチェックを行う
// @Description apiのヘルスチェックを行う。ルーティングが正常に登録されているかを確かめる。
// @Tags        HealthCheck
// @Success     200 {object} presenter.SuccessResponse[healthResponse] "Health check message""
// @Router      /health [get]
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	resp := healthResponse{
		HealthCheck: "ok",
	}
	presenter.RespondOK(w, resp)
}
