package testhelper

import (
	"regexp"
	"testing"
)

func NomalizeJWT(t *testing.T, b []byte) []byte {
	t.Helper()

	// 正規表現でJWT全体を <REPLACED_TOKEN> に置き換える
	regexp := regexp.MustCompile(`"signed_token":\s?"[^"]+"`)
	s := regexp.ReplaceAllString(string(b), `"signed_token": "jwt"`)
	return []byte(s)
}
