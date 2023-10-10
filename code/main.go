package main

import (
	"code/ArrayList"
	"code/StackArray"
	"fmt"
)

func main1() {
	mystack := StackArray.NewStack()
	mystack.Push(1)
	mystack.Push(2)
	mystack.Push(3)
	mystack.Push(4)

	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
}

func main2() {
	// list := ArrayList.NewArrayList()
	// 定义接口对象，赋值的对象必须实现接口的所有方法
	var list ArrayList.List = ArrayList.NewArrayList()
	list.Append("a")
	list.Append("b")
	list.Append("c")
	list.Append("d")

	fmt.Println(list)
	fmt.Println(list.Size())
}

func main3() {
	// list := ArrayList.NewArrayList()
	// 定义接口对象，赋值的对象必须实现接口的所有方法
	var list ArrayList.List = ArrayList.NewArrayList()
	list.Append("a")
	list.Append("b")
	list.Append("c")
	list.Append("d")
	for i := 0; i < 10; i++ {
		list.Insert(1, "z")
		fmt.Println(list)
	}
	for i := 0; i < 10; i++ {
		list.Delete(1)
		fmt.Println(list)
	}
	//fmt.Println(list)
	//fmt.Println(list.Size())
}

func main4() {
	var list ArrayList.List = ArrayList.NewArrayList()
	list.Append("1")
	list.Append("2")
	list.Append("3")
	list.Append("4")
	list.Append("5")
	list.Append("6")

	it := list.Iterator()
	for it.HasNext() {
		item, _ := it.Next()
		fmt.Println(item)
	}
	item, err := it.Next()
	fmt.Println(item)
	fmt.Println(err)
}

func main() {
	mystack := ArrayList.NewArrayListStack()
	mystack.Push(1)
	mystack.Push(2)
	mystack.Push(3)
	mystack.Push(4)

	//fmt.Println(mystack.Pop())
	//fmt.Println(mystack.Pop())
	//fmt.Println(mystack.Pop())
	//fmt.Println(mystack.Pop())
	for it := mystack.IteratorX(); it.HasNext(); {
		item, _ := it.Next()
		fmt.Println(item)
	}
}
