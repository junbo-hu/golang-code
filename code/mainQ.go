package main

import (
	"code/CircleQueue"
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
func main7() {
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

func main() {
	var queue CircleQueue.CircleQueue
	CircleQueue.InitQueue(&queue)

	CircleQueue.EnQueue(&queue, 1)
	CircleQueue.EnQueue(&queue, 2)
	CircleQueue.EnQueue(&queue, 3)
	CircleQueue.EnQueue(&queue, 4)
	CircleQueue.EnQueue(&queue, 5)

	fmt.Println(CircleQueue.DeQueue(&queue))
	fmt.Println(CircleQueue.DeQueue(&queue))
	fmt.Println(CircleQueue.DeQueue(&queue))
	fmt.Println(CircleQueue.DeQueue(&queue))
	fmt.Println(CircleQueue.DeQueue(&queue))
	fmt.Println(CircleQueue.DeQueue(&queue))
}
