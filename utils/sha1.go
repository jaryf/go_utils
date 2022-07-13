package utils

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"
)

func Sha1EncryptFile(path string) (encrypt string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	h := sha1.New()
	_, err = io.Copy(h, f)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func Sha1EncryptBytes(data []byte) (encrypt string, err error) {
	h := sha1.New()
	_, err = io.Copy(h, bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func Sha1EncryptString(data string) (encrypt string, err error) {
	return Sha1EncryptBytes([]byte(data))
}
