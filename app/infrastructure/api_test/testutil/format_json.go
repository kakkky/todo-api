package testutil

import (
	"bytes"
	"encoding/json"
	"testing"
)

func FormatJSON(t *testing.T, b []byte) []byte {
	t.Helper()

	var out bytes.Buffer
	// bodyをインデント整形し、バッファoutに書き込む
	err := json.Indent(&out, b, "", "  ")
	if err != nil {
		t.Fatal(err)
	}

	return out.Bytes()
}
