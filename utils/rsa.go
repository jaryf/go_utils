package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
)

/*
pubKey := `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC6f66pBZ7wS7+eLOGDpILzyRet
Eib53ycQ/roudbhCB4nALvzscm6MyjAZoVf4A/V0PU4LjjIaLQVkQ/sO9oxTUKqj
CxCdq1go8TALGoKZKkPDhno1On/h7mS0eXTlstHGt1NmDRPhT8Ah+Q5xmCRnTHWf
SqHAlW2BkeMfW+GrPQIDAQAB
-----END PUBLIC KEY-----`
*/
func RsaEncode(pubKey, data string) (res string, err error) {
	pubBlock, _ := pem.Decode([]byte(pubKey))
	publicKeyValue, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
	if err != nil {
		return "", err
	}
	pub := publicKeyValue.(*rsa.PublicKey)
	enBytes, err := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(data))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(enBytes), nil
}
