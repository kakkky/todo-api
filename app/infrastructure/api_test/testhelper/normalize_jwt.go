package testhelper

import (
	"regexp"
	"testing"
)

func NormalizeJWT(t *testing.T, b []byte) []byte {
	t.Helper()

	regexp := regexp.MustCompile(`"jwt_token":\s?"[^"]+"`)
	s := regexp.ReplaceAllString(string(b), `"jwt_token": "jwt"`)
	return []byte(s)
}
