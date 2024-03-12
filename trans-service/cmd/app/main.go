package main

import (
	"github.com/arloz/trans-service/src/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	// 创建echo服务对象
	e := echo.New()

	// 初始化所有的请求处理映射
	initHandlers(e)

	// 在端口8090启动服务
	e.Logger.Fatal(e.Start(":8090"))
}

func initHandlers(e *echo.Echo) {
	// 翻译文本 /api/translate?content=xxx
	e.GET("/api/translate", handler.Translate)
}
