package grand

import (
	"testing"
)

func TestIntn(t *testing.T) {
	s := Intn(500)
	t.Log(s)
}

func BenchmarkIntn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Intn(500)
	}
}

func TestRandBytes(t *testing.T) {
	b := RandBytes(10)
	t.Log(b)
}

func BenchmarkRandBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandBytes(10)
	}
}

func TestRandRangeIntN(t *testing.T) {
	n := RandRangeIntN(1, 4)
	t.Log(n)
}

func TestS(t *testing.T) {
	s := S(50, false)
	t.Log(s)
}

func TestRandStr(t *testing.T) {
	randStr := RandStr("01234你耗阿萨德绿卡结束勒肯定就看", 5)
	t.Log(randStr)
}

func BenchmarkRandStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStr("01234你耗阿萨德绿卡结束勒肯定就看", 5)
	}
}

func TestRandDigits(t *testing.T) {
	s := RandDigits(5)
	t.Log(s)
}

func TestRandLetters(t *testing.T) {
	s := RandLetters(5)
	t.Log(s)
}

func TestRandSymbols(t *testing.T) {
	s := RandSymbols(5)
	t.Log(s)
}

func TestRandPerm(t *testing.T) {
	s := RandPerm(5)
	t.Log(s)
}
