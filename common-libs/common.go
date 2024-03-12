package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	//////////////////////////////////// json 库 ////////////////////////////////////
	jsonFunc()

	//////////////////////////////////// log 库 ////////////////////////////////////
	logFunc()

}

func logFunc() {
	// 创建默认的标准logger
	logger := log.Default()

	// 获取默认logger的flag配置
	flag := logger.Flags()
	fmt.Println("default log flag: ", flag)
	logger.Println("default logger message")

	// 指定logger输出的Writer
	var buf bytes.Buffer
	custLogger := log.New(&buf, "testPrefix", log.Lmicroseconds)
	custLogger.Println("test log message")
	fmt.Println(string(buf.Bytes()))
}

func jsonFunc() {

	type Message struct {
		Name    string `json:"name,omitempty"`
		Type    int
		content string
	}

	// json 序列化结构体对象
	message := Message{Name: "testName", Type: 3, content: "content message"}
	v, err := json.Marshal(message)
	fmt.Println("Message: ", string(v), err)

	// json 反序列化结构体对象
	var decodeMessage Message
	err = json.Unmarshal(v, &decodeMessage)
	fmt.Println("Message json Decode: ", decodeMessage, err)

	// json 序列化map对象
	m := make(map[string]string)
	m["name"] = "testMapName"
	m["values"] = "testValues"
	v, err = json.Marshal(m)
	fmt.Println("Map: ", string(v), err)

	// json 反序列化map对象
	var decodeMap map[string]string
	err = json.Unmarshal(v, &decodeMap)
	fmt.Println("Map json Decode: ", decodeMap, err)
}
