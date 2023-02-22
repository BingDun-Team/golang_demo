package main

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/go-resty/resty/v2"
)

type Response struct {
	Code int
	Em   string
	Et   string
	Data *ResponseData
}

type ResponseData struct {
	PuzzleToken string
}

func main() {

	client := resty.New()

	//接口信息
	puzzleToken := "8bb72d21ee65e325c548da0e04bdd3eb"
	reviewApiUrl := "https://bingdun.apis.show/api/review?puzzle_token=" + puzzleToken
	authID := "6e565a15d7da27b5d1c949357761a8e4"
	authSecretKey := "912097369277ecb5dac3d1bd7ab00d2e"
	timeAt := "1675750472"

	//生成签名
	sign := Sha256(puzzleToken + authID)

	responseData := &Response{}
	response, requestErr := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"auth_id":"` + authID + `", "auth_secret_key":"` + authSecretKey + `", "time_at":"` + timeAt + `", "sign":"` + sign + `"}`).
		SetResult(responseData).
		Post(reviewApiUrl)
	if requestErr != nil {
		//请求失败
	}

	if response.StatusCode() == 200 {
		//冰盾服务响应成功
		if responseData.Code == 0 {
			//校验成功 进行自己的业务逻辑

		} else {
			//校验异常 参照Code码进行更正

		}
	} else {
		//冰盾服务响应失败

	}

}

// Sha256加密
func Sha256(str string) string {
	m := sha256.New()
	m.Write([]byte(str))
	res := hex.EncodeToString(m.Sum(nil))
	return res
}
