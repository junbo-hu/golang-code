package StackArray

type StackArray interface {
	Clear()
	Size() int
	Pop() interface{}
	Push(data interface{})
	IsFull() bool
	IsEmpty() bool
}

type Stack struct {
	dataSource  []interface{}
	capsize     int
	currentSize int
}

func NewStack() *Stack {
	mystack := new(Stack)
	mystack.dataSource = make([]interface{}, 0, 10)
	mystack.capsize = 10
	mystack.currentSize = 0
	return mystack
}

func (mystack *Stack) Clear() {

}
func (mystack *Stack) Size() int {

}
func (mystack *Stack) Pop() interface{} {

}
func (mystack *Stack) Push(data interface{}) {

}
func (mystack *Stack) IsFull() bool {}

func (mystack *Stack) IsEmpty() bool {
	
}
