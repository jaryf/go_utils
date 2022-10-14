package gencrypt

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
)

func Md5Bytes(data []byte) (res string, err error) {
	h := md5.New()
	if _, err = h.Write(data); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func Md5String(data string) (res string, err error) {
	return Md5Bytes([]byte(data))
}

func Md5File(path string) (res string, err error) {
	var fileBytes []byte
	fileBytes, err = ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return Md5Bytes(fileBytes)
}
