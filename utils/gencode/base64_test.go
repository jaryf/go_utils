package gencode

import "testing"

func TestB64EncodeString(t *testing.T) {
	s := B64EncodeString("123456")
	t.Log(s)
}

func BenchmarkB64EncodeString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		B64EncodeString("123456")
	}
}
