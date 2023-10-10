package ArrayList

import (
	"errors"
	"fmt"
)

// 接口
type List interface {
	Size() int
	Get(index int) (interface{}, error)
	Set(index int, newval interface{}) error
	Insert(index int, newval interface{}) error
	Append(newval interface{})
	Clear()
	Delete(index int) error
	String() string
	Iterator() Iterator
}

// 数据结构
type ArrayList struct {
	dataStore []interface{} // interface{} 代表泛型
	theSize   int
}

func NewArrayList() *ArrayList {
	list := new(ArrayList)
	list.dataStore = make([]interface{}, 0, 10)
	list.theSize = 0
	return list
}

func (list *ArrayList) Size() int {
	return list.theSize
}

func (list *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index >= list.theSize {
		return nil, errors.New("索引越界")
	}
	return list.dataStore[index], nil
}

func (list *ArrayList) Append(newval interface{}) {
	list.dataStore = append(list.dataStore, newval)
	list.theSize++
}

func (list *ArrayList) String() string {
	return fmt.Sprint(list.dataStore) // 打印到字符串
}

func (list *ArrayList) Set(index int, newval interface{}) error {
	if index < 0 || index >= list.theSize {
		return errors.New("索引越界")
	}
	list.dataStore[index] = newval
	return nil
}
func (list *ArrayList) Insert(index int, newval interface{}) error {
	if index < 0 || index >= list.theSize {
		return errors.New("索引越界")
	}
	list.checkIsFull()
	list.dataStore = list.dataStore[:list.theSize+1] // 插入数据，内存移动一位

	for i := list.theSize; i > index; i-- { //从后往前移动
		list.dataStore[i] = list.dataStore[i-1]
	}
	list.dataStore[index] = newval
	list.theSize++
	return nil
}
func (list *ArrayList) Clear() {
	list.dataStore = make([]interface{}, 0, 10)
	list.theSize = 0
}
func (list *ArrayList) Delete(index int) error {
	if index < 0 || index >= list.theSize {
		return errors.New("索引越界")
	}
	list.dataStore = append(list.dataStore[:index], list.dataStore[index+1:]...) //加上'...'就可以追加一个数组， 否则只能追加一个元素
	list.theSize--
	return nil
}

func (list *ArrayList) checkIsFull() {
	if list.theSize >= cap(list.dataStore) {
		newdatastore := make([]interface{}, list.theSize, 2*list.theSize)
		copy(newdatastore, list.dataStore)
		list.dataStore = newdatastore
	}
}
