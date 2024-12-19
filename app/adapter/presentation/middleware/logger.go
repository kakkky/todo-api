package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// HTTP リクエストとレスポンスをログに記録するミドルウェア
func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		// リクエストの処理
		requestLog(r)

		// レスポンス用ラッパーを作成
		buf := &bytes.Buffer{} //レスポンスボディを書き込むためのバッファ
		rww := newResponseWrapper(rw, buf)

		// ハンドラ呼び出し
		// ラッパーをResponseWriterとして渡し、バッファにもレスポンスを書き込むようにする
		h.ServeHTTP(rww, r)

		// レスポンスの処理
		responseLog(rww, buf)
	})
}

// リクエストボディを取得しログに出力
func requestLog(r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	strReqBody := string(reqBody)
	// リクエストボディをログに記録
	fmt.Printf("-----------------------------------------------------------------------------------------------------\n"+
		"[REQUEST]\n"+
		"Timestamp  : %s\n"+
		"Method     : %s\n"+
		"URL        : %s\n"+
		"Header     : %s\n"+
		"Body       : \n%s\n",
		time.Now().Format(time.RFC3339),
		r.Method,
		r.URL.Path,
		r.Header.Get("Authorization"),
		strReqBody)

	// リクエストボディを再代入
	// リクエストボディはストリームなので再代入する必要がある
	r.Body = io.NopCloser(bytes.NewBuffer(reqBody))
}

// 　レスポンスボディを整形しログに出力
func responseLog(rww *rwWrapper, buf *bytes.Buffer) {
	var formattedRespBody bytes.Buffer
	json.Indent(&formattedRespBody, buf.Bytes(), "", "    ")
	status := rww.statusCode
	fmt.Printf(
		"\n[RESPONSE]\n"+
			"Timestamp  : %s\n"+
			"Status     : %d\n"+
			"Body       :\n%s"+
			"-----------------------------------------------------------------------------------------------------\n",
		time.Now().Format(time.RFC3339),
		status,
		formattedRespBody.String(),
	)
}

// レスポンスのログ記録用にレスポンスをラップ
type rwWrapper struct {
	rw          http.ResponseWriter
	multiWriter io.Writer
	statusCode  int
}

func newResponseWrapper(rw http.ResponseWriter, buf io.Writer) *rwWrapper {
	return &rwWrapper{
		rw:          rw,
		multiWriter: io.MultiWriter(rw, buf),
	}
}

// http.ResponseWriterインターフェースを満たすためのメソッド
func (rww *rwWrapper) Header() http.Header {
	return rww.rw.Header()
}
func (rww *rwWrapper) Write(b []byte) (int, error) {
	return rww.multiWriter.Write(b)
}
func (rww *rwWrapper) WriteHeader(statusCode int) {
	rww.statusCode = statusCode
	rww.rw.WriteHeader(statusCode)
}
