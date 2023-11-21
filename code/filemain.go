package main

import (
	"fmt"
	"os"
)

func main() {

	// 读取文件内容
	content, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("文件读取失败：", err)
		return
	}
	fmt.Println("文件内容为：", string(content))

	// 写入文件内容
	err = os.WriteFile("example.txt", []byte("Hello, world!"), 0644)
	if err != nil {
		fmt.Println("文件写入失败：", err)
		return
	}
	fmt.Println("文件写入成功！")
}
