package client

/**
* 定义翻译客户端接口
 */
type TransClient interface {
	// 客户端初始化
	Init() error

	// 执行文本翻译
	DoTrans(content string) (string, error)
}
