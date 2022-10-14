package grand

import (
	"crypto/rand"
	"encoding/binary"
	"time"
)

const (
	bufferChanSize = 10000
)

var (
	letters    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" // 52
	symbols    = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"                   // 32
	digits     = "0123456789"                                           // 10
	characters = letters + digits + symbols                             // 94
	bufferChan = make(chan []byte, bufferChanSize)
)

func init() {
	go asyncProducingRandomBufferBytesLoop()
}

func asyncProducingRandomBufferBytesLoop() {
	var step int
	for {
		buffer := make([]byte, 1024)
		if n, err := rand.Read(buffer); err != nil {
			panic(`error reading random buffer from system`)
		} else {
			for _, step = range []int{4} {
				for i := 0; i <= n-4; i += step {
					bufferChan <- buffer[i : i+4]
				}
			}
		}
	}
}

// Intn 返回一个随机数字
func Intn(max int) int {
	if max <= 0 {
		return max
	}
	n := int(binary.LittleEndian.Uint32(<-bufferChan)) % max
	if (max > 0 && n < 0) || (max < 0 && n > 0) {
		return -n
	}
	return n
}

// RandBytes 返回随机长度的byte数组
func RandBytes(n int) []byte {
	if n <= 0 {
		return nil
	}
	i := 0
	b := make([]byte, n)
	for {
		copy(b[i:], <-bufferChan)
		i += 4
		if i >= n {
			break
		}
	}
	return b
}

// RandRangeIntN 返回min到max之间的随机整数,支持负数,包含边界
func RandRangeIntN(min, max int) int {
	if min >= max {
		return min
	}
	if min >= 0 {
		return Intn(max-min+1) + min
	}
	return Intn(max+(0-min)+1) - (0 - min)
}

// 返回指定长度的数字、字符，第二个参数symbols用于指定知否返回的随机字符串中包含特殊字符
func S(n int, symbols ...bool) string {
	if n <= 0 {
		return ""
	}
	var (
		b           = make([]byte, n)
		numberBytes = RandBytes(n)
	)
	for i := range b {
		if len(symbols) > 0 && symbols[0] {
			b[i] = characters[numberBytes[i]%94]
		} else {
			b[i] = characters[numberBytes[i]%62]
		}
	}
	return string(b)
}

func D(min, max time.Duration) time.Duration {
	multiple := int64(1)
	if min != 0 {
		for min%10 == 0 {
			multiple *= 10
			min /= 10
			max /= 10
		}
	}
	n := int64(RandRangeIntN(int(min), int(max)))
	return time.Duration(n * multiple)
}

// Str 用于从给定的字符列表中选择指定长度的随机字符串返回
func RandStr(s string, n int) string {
	if n <= 0 {
		return ""
	}
	var (
		b     = make([]rune, n)
		runes = []rune(s)
	)
	if len(runes) <= 255 {
		numberBytes := RandBytes(n)
		for i := range b {
			b[i] = runes[int(numberBytes[i])%len(runes)]
		}
	} else {
		for i := range b {
			b[i] = runes[Intn(len(runes))]
		}
	}
	return string(b)
}

// RandDigits 返回指定长度的随机数字字符串
func RandDigits(n int) string {
	if n <= 0 {
		return ""
	}
	var (
		b           = make([]byte, n)
		numberBytes = RandBytes(n)
	)
	for i := range b {
		b[i] = digits[numberBytes[i]%10]
	}
	return string(b)
}

// RandLetters 返回指定长度的随机英文字符串
func RandLetters(n int) string {
	if n <= 0 {
		return ""
	}
	var (
		b           = make([]byte, n)
		numberBytes = RandBytes(n)
	)
	for i := range b {
		b[i] = letters[numberBytes[i]%52]
	}
	return string(b)
}

// RandSymbols 返回指定长度的随机特殊字符串
func RandSymbols(n int) string {
	if n <= 0 {
		return ""
	}
	var (
		b           = make([]byte, n)
		numberBytes = RandBytes(n)
	)
	for i := range b {
		b[i] = symbols[numberBytes[i]%32]
	}
	return string(b)
}

// RandPerm 返回指定长度的随机int数组
func RandPerm(n int) []int {
	m := make([]int, n)
	for i := 0; i < n; i++ {
		j := Intn(i + 1)
		m[i] = m[j]
		m[j] = i
	}
	return m
}
