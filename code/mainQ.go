package main

import (
	"code/Queue"
	"fmt"
	"io/ioutil"
)

func main11() {
	myq := Queue.NewQueue()
	for i := 0; i < 100; i++ {
		myq.EnQueue(1)
		myq.EnQueue(2)
		myq.EnQueue(3)
		myq.EnQueue(4)
	}

	fmt.Println(myq.Size())

	fmt.Println(myq.DeQueue())
	fmt.Println(myq.DeQueue())
	fmt.Println(myq.DeQueue())
	fmt.Println(myq.DeQueue())
	fmt.Println(myq.DeQueue())
}

// 用队列遍历文件夹，广度优先
func main() {
	path := "E:\\学习\\go-project"
	files := []string{}
	mystack := Queue.NewQueue()
	mystack.EnQueue(path)
	for !mystack.IsEmpty() {
		path := mystack.DeQueue().(string)
		files = append(files, path)
		read, _ := ioutil.ReadDir(path)
		for _, fi := range read {
			if fi.IsDir() {
				fulldir := path + "\\" + fi.Name()
				files = append(files, fulldir)
				mystack.EnQueue(fulldir)
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
