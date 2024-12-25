package testhelper

import "regexp"

// ランダムなULID値を "PLACEHOLDER_ULID" に置換
func NormalizeULID(b []byte) []byte {
	regexp := regexp.MustCompile(`[0-9A-HJKMNP-TV-Z]{26}`)
	s := regexp.ReplaceAllString(string(b), "ulid")
	return []byte(s)
}
