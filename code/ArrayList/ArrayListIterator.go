package ArrayList

import "errors"

// 建立在ArrayList之上
type Iterator interface {
	HasNext() bool
	Next() (interface{}, error)
	Remove()
	GetIndex() int
}

type Iterable interface {
	Iterator() Iterator // 构造初始化接口
}

type ArrayListIterator struct {
	list         *ArrayList // 数组指针
	currentIndex int        // 当前索引
}

func (list *ArrayList) Iterator() Iterator {
	it := new(ArrayListIterator) //构造迭代器
	it.currentIndex = 0
	it.list = list
	return it
}

func (it *ArrayListIterator) HasNext() bool {
	return it.currentIndex < it.list.theSize //是否有下一个
}
func (it *ArrayListIterator) Next() (interface{}, error) {
	if !it.HasNext() {
		return nil, errors.New("no more")
	}
	value, err := it.list.Get(it.currentIndex)
	it.currentIndex++
	return value, err
}
func (it *ArrayListIterator) Remove() {
	it.currentIndex--
	it.list.Delete(it.currentIndex)
}
func (it *ArrayListIterator) GetIndex() int {
	return it.currentIndex
}
