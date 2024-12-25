package testhelper

import "regexp"

func NomalizeJWT(b []byte) []byte {
	// 正規表現でJWT全体を <REPLACED_TOKEN> に置き換える
	regexp := regexp.MustCompile(`"signed_token":\s?"[^"]+"`)
	s := regexp.ReplaceAllString(string(b), `"signed_token": "jwt"`)
	return []byte(s)
}
