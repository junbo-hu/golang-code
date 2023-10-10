package ArrayList

import "errors"

// 建立在ArrayList之上
type IteratorX interface {
	HasNext() bool
	Next() (interface{}, error)
	Remove()
	GetIndex() int
}

type IterableX interface {
	IteratorX() IteratorX // 构造初始化接口
}

type ArrayListIteratorX struct {
	stack        *Stack // 数组指针
	currentIndex int    // 当前索引
}

func (stack *Stack) IteratorX() IteratorX {
	it := new(ArrayListIteratorX) //构造迭代器
	it.currentIndex = 0
	it.stack = stack
	return it
}

func (it *ArrayListIteratorX) HasNext() bool {
	return it.currentIndex < it.stack.myarray.theSize //是否有下一个
}
func (it *ArrayListIteratorX) Next() (interface{}, error) {
	if !it.HasNext() {
		return nil, errors.New("no more")
	}
	value, err := it.stack.myarray.Get(it.currentIndex)
	it.currentIndex++
	return value, err
}
func (it *ArrayListIteratorX) Remove() {
	it.currentIndex--
	it.stack.myarray.Delete(it.currentIndex)
}
func (it *ArrayListIteratorX) GetIndex() int {
	return it.currentIndex
}
