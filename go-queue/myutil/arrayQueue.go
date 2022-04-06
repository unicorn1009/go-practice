package myutil

import "errors"

// type Person struct {
// 	Name string
// 	City string
// 	Age  int8
// }

type ArrayQueue struct {
	elements []interface{}
	head     int // 队头，出队端
	tail     int // 队尾，入队端
}

// 将元素添加到队尾，如果成功，则返回true。
func (receiver *ArrayQueue) Offer(e interface{}) bool {
	if receiver.elements == nil {
		receiver.elements = make([]interface{}, 16)
	}
	err := receiver.addLast(e)
	return err == nil
}

func (receiver *ArrayQueue) addLast(e interface{}) error {
	if e == nil {
		return errors.New("元素不可为nil")
	}
	// 添加元素
	receiver.elements[receiver.tail] = e

	// 维护 队尾 指针
	receiver.tail = (receiver.tail + 1) & (len(receiver.elements) - 1)
	// TODO 判断是否需要扩容

	return nil
}

// 将队首的元素删除，并返回该元素。如果队列为空，返回nil
func (receiver *ArrayQueue) Poll() interface{} {
	return receiver.pollFirst()
}

func (receiver *ArrayQueue) pollFirst() interface{} {
	h := receiver.head
	result := receiver.elements[h]
	if result == nil {
		return nil
	}
	receiver.elements[h] = nil
	// 维护队头指针
	receiver.head = (h + 1) & (len(receiver.elements) - 1)
	return result
}

func (receiver *ArrayQueue) Peek() interface{} {
	return receiver.elements[receiver.head]
}
