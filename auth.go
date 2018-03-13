package wechatools

import (
	"errors"
	"fmt"
	"strconv"
)

//GetAccessToken 获取公众号 AccessToken
func GetAccessToken(appid, secretkey string) (string, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", appid, secretkey)
	rs, err := Get(url)
	if err != nil {
		return "", err
	}
	results := JSON2Map(rs)
	if v, ok := results["errcode"]; ok {
		errcode := v.(int)
		str := strconv.Itoa(errcode)
		return "", errors.New(str)
	}
	if v, ok := results["access_token"]; ok {
		if v != "" {
			return v.(string), nil
		}
		return "", errors.New("get access token fail,access_token is empty")
	}
	return "", errors.New("get access token fail,respon key access_token is null")
}

//GetJsapiTicket 获取公众号Jsapi Ticket
func GetJsapiTicket(accessToken string) (string, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi", accessToken)
	rs, err := Get(url)
	if err != nil {
		return "", err
	}
	results := JSON2Map(rs)
	if v, ok := results["ticket"]; ok {
		if v != "" {
			return v.(string), nil
		}
		return "", errors.New("get ticket fail,ticket is empty")
	}
	return "", errors.New("get ticket fail,respon key ticket is null")
}
