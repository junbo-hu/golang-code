package CircleQueue

import "errors"

const QueueSize = 5 //固定大小，最多能存QueueSize-1个数据

type CircleQueue struct {
	data  [QueueSize]interface{}
	front int
	rear  int
}

// 另外一种风格
func InitQueue(q *CircleQueue) { // 初始化，头部尾部重合，为空
	q.front = 0
	q.rear = 0
}

func EnQueue(q *CircleQueue, data interface{}) (err error) {
	if (q.rear+1)%QueueSize == q.front { //空一个位置表示满格
		return errors.New("队列已满了")
	}
	q.data[q.rear] = data
	q.rear = (q.rear + 1) % QueueSize
	return nil
}

func DeQueue(q *CircleQueue) (data interface{}, err error) {
	if q.rear == q.front {
		return nil, errors.New("队列为空")
	}
	res := q.data[q.front]
	q.data[q.front] = 0
	q.front = (q.front + 1) % QueueSize
	return res, nil
}

func QueueLength(q *CircleQueue) int {
	return (q.rear - q.front + QueueSize) % QueueSize
}
