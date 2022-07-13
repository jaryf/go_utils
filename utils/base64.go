package utils

import (
	"encoding/base64"
	"io/ioutil"
)

// B64Encode encodes bytes with BASE64 algorithm.
func B64Encode(src []byte) []byte {
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(src)))
	base64.StdEncoding.Encode(dst, src)
	return dst
}

// B64EncodeString encodes string with BASE64 algorithm.
func B64EncodeString(src string) string {
	return B64EncodeToString([]byte(src))
}

// B64EncodeToString encodes bytes to string with BASE64 algorithm.
func B64EncodeToString(src []byte) string {
	return string(B64Encode(src))
}

// B64EncodeFile encodes file content of `path` using BASE64 algorithms.
func B64EncodeFile(path string) ([]byte, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return B64Encode(content), nil
}

// EncodeFileToString encodes file content of `path` to string using BASE64 algorithms.
func EncodeFileToString(path string) (string, error) {
	content, err := B64EncodeFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// B64Decode decodes bytes with BASE64 algorithm.
func B64Decode(data []byte) ([]byte, error) {
	var (
		src    = make([]byte, base64.StdEncoding.DecodedLen(len(data)))
		n, err = base64.StdEncoding.Decode(src, data)
	)
	return src[:n], err
}

// B64DecodeString decodes string with BASE64 algorithm.
func B64DecodeString(data string) ([]byte, error) {
	return B64Decode([]byte(data))
}

// B64DecodeToString decodes string with BASE64 algorithm.
func B64DecodeToString(data string) (string, error) {
	b, err := B64DecodeString(data)
	return string(b), err
}
