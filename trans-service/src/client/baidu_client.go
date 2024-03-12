package client

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"os"

	"github.com/arloz/trans-service/pkg/util/logger"
)

type BaiduTransClient struct {
	appId  string
	appKey string
}

func (client *BaiduTransClient) Init() error {
	// 从系统环境变量中读取百度翻译API的appId和appKey信息
	client.appId = os.Getenv("BAIDU_TRANS_APP_ID")
	client.appKey = os.Getenv("BAIDU_TRANS_APP_KEY")

	logger.Infof("laod baidu translate client success: %v", client.appId)
	return nil
}

func (client *BaiduTransClient) DoTrans(content string) (string, error) {
	if len(client.appId) <= 0 || len(client.appKey) <= 0 {
		return "", errors.New("baidu translate client did not init")
	}

	req := apiRequest{appid: client.appId, q: content, from: "auto", to: "zh", salt: generateSalt()}
	req.sign = client.doSign(&req)

	resp, err := doRequest(&req)
	if err != nil {
		return "", err
	}

	return resp.Dst, nil
}

// 百度API请求信息
type apiRequest struct {
	appid string
	q     string
	from  string
	to    string
	salt  string
	sign  string
}

// 百度API响应结果信息
type apiResult struct {
	From        string          `json:"from"`
	To          string          `json:"to"`
	TransResult []apiResultItem `json:"trans_result"`
}

// 百度API结果对象
type apiResultItem struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
}

// 执行API调用请求
func doRequest(transReq *apiRequest) (*apiResultItem, error) {
	// 请求参数
	params := url.Values{}
	params.Add("appid", transReq.appid)
	params.Add("q", transReq.q)
	params.Add("from", transReq.from)
	params.Add("to", transReq.to)
	params.Add("salt", transReq.salt)
	params.Add("sign", transReq.sign)

	// 创建一个新的 HTTP 请求
	url := "https://fanyi-api.baidu.com/api/trans/vip/translate?" + params.Encode()
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("request translate api fail: " + resp.Status)
	}

	// 读取响应 body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 解析结果
	var result apiResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	if result.TransResult == nil {
		return nil, errors.New("translate empty")
	}

	return &(result.TransResult[0]), nil
}

// 加签计算
func (client *BaiduTransClient) doSign(req *apiRequest) string {
	signContent := client.appId + req.q + req.salt + client.appKey

	// 计算 MD5 哈希值, 将哈希值转换为十六进制字符串
	h := md5.New()
	h.Write([]byte(signContent))
	sum := h.Sum(nil)
	md5Str := fmt.Sprintf("%x", sum)

	return md5Str
}

func generateSalt() string {
	// 生成一个随机字符串
	str := make([]byte, 10)
	for i := range str {
		str[i] = byte(rand.Intn(26)) + 'a'
	}

	return string(str)
}
