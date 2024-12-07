package health

import (
	"encoding/json"
	"log"
	"net/http"
)

// HealthCheckHandler godoc
// @Summary     apiのヘルスチェックを行う
// @Description apiのヘルスチェックを行う。ルーティングが正常に登録されているかを確かめる。
// @Success     200 {object} map[string]string "Health check message""
// @Router      /health [get]
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	body := map[string]string{"message": "health check"}
	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Printf("failed to encode to json : %v", err)
	}
}
