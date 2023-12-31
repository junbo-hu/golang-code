package ArrayList

type StackArray interface {
	Clear()
	Size() int
	Pop() interface{}
	Push(data interface{})
	IsFull() bool
	IsEmpty() bool
	IteratorX() IteratorX
}

type Stack struct {
	myarray *ArrayList
	capsize int
}

func NewArrayListStack() *Stack {
	mystack := new(Stack)
	mystack.myarray = NewArrayList()
	mystack.capsize = 10
	return mystack
}

func (mystack *Stack) Clear() {
	mystack.myarray.Clear()
	mystack.capsize = 0
}
func (mystack *Stack) Size() int {

	return mystack.myarray.Size()
}
func (mystack *Stack) Pop() interface{} {

	if !mystack.IsEmpty() {
		last := mystack.myarray.dataStore[mystack.myarray.theSize-1]
		mystack.myarray.Delete(mystack.myarray.theSize - 1)
		return last
	}
	return nil
}
func (mystack *Stack) Push(data interface{}) {
	if !mystack.IsFull() {
		mystack.myarray.Append(data)
	}
}
func (mystack *Stack) IsFull() bool {
	if mystack.myarray.Size() >= mystack.capsize {
		return true
	} else {
		return false
	}
}

func (mystack *Stack) IsEmpty() bool {
	if mystack.myarray.Size() == 0 {
		return true
	} else {
		return false
	}
}
