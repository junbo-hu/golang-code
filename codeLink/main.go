package main

import "fmt"

func main1() {
	stack := NewStack()
	for i := 1; i < 100; i++ {
		stack.Push(i)
	}

	for data := stack.Pop(); data != nil; data = stack.Pop() {
		fmt.Println(data)
	}
}

func main() {
	queue := NewLinkQueue()
	for i := 0; i < 100; i++ {
		queue.Enqueue(i)
	}

	for data := queue.Dequeue(); data != nil; data = queue.Dequeue() {
		fmt.Println(data)
	}
}
