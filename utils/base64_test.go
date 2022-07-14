package utils

import "testing"

func TestB64EncodeString(t *testing.T) {
	s := B64EncodeString("123456")
	t.Log(s)
}
