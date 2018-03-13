package wechatools

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//Get http Get请求简单封装
func Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

//JSON2Map []byte JSON转map
func JSON2Map(b []byte) map[string]interface{} {
	var f interface{}
	json.Unmarshal(b, &f)
	return f.(map[string]interface{})
}
