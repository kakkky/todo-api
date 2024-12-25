package testutil

import "regexp"

// ランダムなULID値を "PLACEHOLDER_ULID" に置換
func NormalizeULID(b []byte) []byte {
	ulidRegex := regexp.MustCompile(`[0-9A-HJKMNP-TV-Z]{26}`)
	s := ulidRegex.ReplaceAllString(string(b), "ulid")
	return []byte(s)
}
