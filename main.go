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
	puzzleToken := "your puzzle_token"
	reviewApiUrl := "https://bingdun.apis.show/api/review?puzzle_token=" + puzzleToken
	authID := "your auth_id"
	authSecretKey := "your auth_secret_key"
	timeAt := time.Now().Unix()

	//生成签名
	sign := Sha256(puzzleToken + authID)

	responseData := &Response{}
	response, requestErr := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{"auth_id": authID, "auth_secret_key": authSecretKey, "time_at": timeAt, "sign": sign}).
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
