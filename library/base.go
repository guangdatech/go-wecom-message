package library

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// Post 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
// 错误信息
func Post(url string, data interface{}, contentType string) (string,error) {

	// 超时时间：5秒
	client := &http.Client{Timeout: 20 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result),nil
}
