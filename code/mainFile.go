package main

import (
	"code/StackArray"
	"errors"
	"fmt"
	"io/ioutil"
)

// 递归遍历文件夹
func GetAll(path string, files []string, level int) ([]string, error) {
	levelStr := ""
	newlevel := level
	if newlevel == 1 {
		levelStr = ""
	} else {
		for ; newlevel > 1; newlevel-- {
			levelStr += "|--"
		}
	}

	read, err := ioutil.ReadDir(path)
	if err != nil {
		return files, errors.New("文件夹不可读")
	}

	for _, fi := range read {
		if fi.IsDir() {
			fulldir := path + "\\" + fi.Name()
			files = append(files, levelStr+fulldir) //追加路径
			files, _ = GetAll(fulldir, files, level+1)
		} else {
			fulldir := path + "\\" + fi.Name()
			files = append(files, levelStr+fulldir)
		}
	}

	return files, nil
}

func main() {
	path := "E:\\学习\\go-project"
	files := []string{}
	files, err := GetAll(path, files, 1)

	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}
}

// 用栈模拟实现递归遍历文件夹
func main2x() {
	path := "E:\\学习\\go-project"
	files := []string{}
	mystack := StackArray.NewStack()
	mystack.Push(path)
	for !mystack.IsEmpty() {
		path := mystack.Pop().(string)
		files = append(files, path)
		read, _ := ioutil.ReadDir(path)
		for _, fi := range read {
			if fi.IsDir() {
				fulldir := path + "\\" + fi.Name()
				files = append(files, fulldir)
				mystack.Push(fulldir)
			} else {
				fulldir := path + "\\" + fi.Name()
				files = append(files, fulldir)
			}
		}
	}

	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}
}
