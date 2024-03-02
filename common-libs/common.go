package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Name    string `json:"name,omitempty"`
	Type    int
	content string
}

func main() {
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
