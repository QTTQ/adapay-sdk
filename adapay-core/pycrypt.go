package adapayCore

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"strings"
)

func RsaLongEncrypt(content string, msc *MerchSysConfig) (result string) {

	return ""
}

func RsaLongDecrypt(content string, msc *MerchSysConfig) (result string) {

	return ""
}

func RsaSign(content string, msc *MerchSysConfig) (result string, err error) {
	privateKey, err := getPrivateKey(msc)
	if err != nil {
		return "", err
	} else if privateKey == nil {
		return "", errors.New("Error in get private key")
	}

	h := sha1.New()
	h.Write([]byte([]byte(content)))
	hash := h.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey.(*rsa.PrivateKey), crypto.SHA1, hash[:])
	if err != nil {
		Println("Sign Error: ", err)
		return "", err
	}

	out := base64.StdEncoding.EncodeToString(signature)

	return out, nil
}

func RsaSignVerify(cipher string, content string, msc MerchSysConfig) (result bool) {

	return false
}

func getPrivateKey(msc *MerchSysConfig) (interface{}, error) {
	priKeyString := msc.RspPriKey
	if priKeyString == "" {
		return nil, errors.New("Unknow private key data")
	}


	if !strings.Contains(priKeyString, "-----BEGIN PRIVATE KEY-----") {
		priKeyString = "-----BEGIN PRIVATE KEY-----\n" + priKeyString
	}
	if !strings.Contains(priKeyString, "-----END PRIVATE KEY-----") {
		priKeyString = priKeyString + "\n-----END PRIVATE KEY-----"
	}



	defer func() {
		err := recover()
		if err != nil {
			Println("Error in get private key: ", err)
		}
	}()
	keyByts, _ := pem.Decode([]byte(priKeyString))

	privateKey, err := x509.ParsePKCS8PrivateKey(keyByts.Bytes)
	if err != nil {
		return nil, err
	}


	return privateKey, nil
}
