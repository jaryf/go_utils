package gencrypt

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"io/ioutil"
)

func Sha1File(path string) (encrypt string, err error) {
	var fileBytes []byte
	fileBytes, err = ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	h := sha1.New()
	h.Write(fileBytes)
	encrypt = hex.EncodeToString(h.Sum(nil))
	return
}

func Sha1Bytes(data []byte) string {
	h := sha1.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

func Sha1String(data string) (encrypt string) {
	return Sha1Bytes([]byte(data))
}

func Sha256String(data string) string {
	return Sha256Bytes([]byte(data))
}

func Sha256Bytes(data []byte) string {
	return hex.EncodeToString(Sha256(data))
}

func Sha256(data []byte) []byte {
	digest := sha256.New()
	digest.Write(data)
	return digest.Sum(nil)
}

func Sha512String(data string) string {
	return Sha512Bytes([]byte(data))
}

func Sha512Bytes(data []byte) string {
	return hex.EncodeToString(Sha512(data))
}

func Sha512(data []byte) []byte {
	digest := sha512.New()
	digest.Write(data)
	return digest.Sum(nil)
}
