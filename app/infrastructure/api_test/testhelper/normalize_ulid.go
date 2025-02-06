package testhelper

import (
	"regexp"
	"testing"
)

func NormalizeULID(t *testing.T, b []byte) []byte {
	t.Helper()

	regexp := regexp.MustCompile(`"id":\s?"[0-9A-HJKMNP-TV-Z]{26}"`)
	s := regexp.ReplaceAllString(string(b), `"id": "ulid"`)
	return []byte(s)
}
