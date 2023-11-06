package _defer

import (
	"fmt"
	"io"
	"log"
	"os"
)

func FileReadCase() {
	file, err := os.Open("go学习笔记.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := recover()
		if err != nil {
			log.Println(err)
		}
	}()

	defer func() {
		fmt.Println("文件被关闭！")
		file.Close()
	}()

	defer func() {
		fmt.Println("位置1")
	}()

	defer func() {
		fmt.Println("位置2")
	}()
	panic("抛出异常")
	buf := make([]byte, 10)
	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if n == 0 {
			break
		}
		fmt.Println(buf[:n])

	}
}
