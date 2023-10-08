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

	return mystack.currentSize
}
func (mystack *Stack) Pop() interface{} {

	if !mystack.IsEmpty() {
		last := mystack.dataSource[mystack.currentSize-1]
		mystack.dataSource = mystack.dataSource[:mystack.currentSize-1]
		mystack.currentSize--
		return last
	}
	return nil
}
func (mystack *Stack) Push(data interface{}) {
	if !mystack.IsFull() {
		mystack.dataSource = append(mystack.dataSource, data)
		mystack.currentSize++

	}
}
func (mystack *Stack) IsFull() bool {
	if mystack.currentSize >= mystack.capsize {
		return true
	} else {
		return false
	}
}

func (mystack *Stack) IsEmpty() bool {
	if mystack.currentSize == 0 {
		return true
	} else {
		return false
	}
}
