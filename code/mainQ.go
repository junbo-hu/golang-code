package main

import (
	"code/Queue"
	"fmt"
)

func main() {
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
