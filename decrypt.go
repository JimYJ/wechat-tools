package wechatools

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"log"

	"github.com/tidwall/gjson"
)

//pkcs7UnPadding PKCS7 解填充
func pkcs7UnPadding(src []byte) ([]byte, error) {
	length := len(src)
	if length <= 0 {
		return nil, errors.New("pkcs7: source must not be empty slice")
	}
	padLen := int(src[length-1])
	if padLen > length {
		return nil, errors.New("pkcs7: invalid padding (last byte is larger than total length)")
	}
	origLen := length - padLen
	padding := src[origLen:]
	for i := 0; i < padLen; i++ {
		if padding[i] != byte(padLen) {
			return nil, errors.New("pkcs7: invalid padding (last byte does not match padding)")
		}
	}
	return src[:origLen], nil
}

//WxDecrypt 微信小程序解密
func WxDecrypt(encryptedData, iv, sessionKey, appID string) (map[string]interface{}, error) {
	if len(sessionKey) != 24 || len(iv) != 24 {
		return nil, errors.New("SessionKey or iv invalid")
	}
	sessionkey, err := base64.StdEncoding.DecodeString(sessionKey)
	encrypteddata, err2 := base64.StdEncoding.DecodeString(encryptedData)
	Iv, err3 := base64.StdEncoding.DecodeString(iv)
	if err3 != nil {
		return nil, err3
	} else if err != nil {
		return nil, err
	} else if err2 != nil {
		return nil, err2
	}
	plantText := make([]byte, len(encrypteddata))
	block, err := aes.NewCipher(sessionkey)
	if err != nil {
		return nil, err
	}
	model := cipher.NewCBCDecrypter(block, Iv)
	model.CryptBlocks(plantText, encrypteddata)
	log.Println(string(plantText[:]))
	var errs error
	plantText, errs = pkcs7UnPadding(plantText)
	if errs != nil {
		return nil, errs
	}
	newAppip := gjson.Get(string(plantText[:]), "watermark.appid").String()
	if newAppip != appID {
		return nil, errors.New("Invalid Buffer or sessionKey is expired")
	}
	j := gjson.Parse(string(plantText[:])).Value().(map[string]interface{})
	return j, nil
}
