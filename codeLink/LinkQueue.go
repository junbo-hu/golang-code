package main

type QueueLink struct {
	rear  *Node
	front *Node
}

type LinkQueue interface {
	length() int
	Enqueue(value interface{})
	Dequeue() interface{}
}

func NewLinkQueue() *QueueLink {
	return &QueueLink{}
}

func (qlk *QueueLink) length() int {
	pnext := qlk.front
	length := 0
	for pnext.pNext != nil {
		pnext = pnext.pNext
		length++
	}
	return length
}
func (qlk *QueueLink) Enqueue(value interface{}) {
	newNode := &Node{value, nil} //尾部追加，头部删除
	if qlk.front == nil {
		qlk.front = newNode
		qlk.rear = newNode
	} else {
		qlk.rear.pNext = newNode
		qlk.rear = qlk.rear.pNext
	}
}
func (qlk *QueueLink) Dequeue() interface{} {
	if qlk.front == nil {
		return nil
	}
	newNode := qlk.front
	if qlk.front == qlk.rear { //只有一个元素
		qlk.front = nil
		qlk.rear = nil
	} else {
		qlk.front = qlk.front.pNext //删除一个
	}
	return newNode.data
}
