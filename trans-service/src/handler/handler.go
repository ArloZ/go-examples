package handler

import (
	"github.com/arloz/trans-service/src/client"

	"net/http"

	"github.com/labstack/echo/v4"
)

/**
 * http 请求响应对象
 */
type Response struct {
	Success bool   `json:"success,omitempty"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

const (
	QUERY_KEY_CONTENT = "content"
)

// 百度翻译客户端，后续可优化为配置化
var transClient client.TransClient

// 初始化
func init() {
	transClient = &client.BaiduTransClient{}
	transClient.Init()
}

func Translate(c echo.Context) error {
	// 获取翻译参数
	content := c.QueryParam(QUERY_KEY_CONTENT)
	if content == "" {
		return BadRequest(c, http.StatusBadRequest, "translate content is blank")
	}

	result, err := transClient.DoTrans(content)
	if err != nil {
		// 翻译失败
		return BadRequest(c, http.StatusInternalServerError, err.Error())
	}

	// 翻译成功，返回对应翻译结果
	return OK(c, result)
}

func OK(c echo.Context, data any) error {
	resp := Response{Success: true, Data: data}
	return c.JSON(http.StatusOK, resp)
}

func BadRequest(c echo.Context, code int, message string) error {
	resp := Response{Success: false, Message: message}
	return c.JSON(code, resp)
}
