package testhelper

import (
	"regexp"
	"testing"
)

// ランダムなULID値を "PLACEHOLDER_ULID" に置換
func NormalizeULID(t *testing.T, b []byte) []byte {
	t.Helper()

	regexp := regexp.MustCompile(`[0-9A-HJKMNP-TV-Z]{26}`)
	s := regexp.ReplaceAllString(string(b), "ulid")
	return []byte(s)
}
