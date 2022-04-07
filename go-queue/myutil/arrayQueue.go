package myutil

import (
	"errors"
)

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

// 入队函数
func (receiver *ArrayQueue) addLast(e interface{}) error {
	if e == nil {
		return errors.New("元素不可为nil")
	}
	// 添加元素
	receiver.elements[receiver.tail] = e

	// 维护 队尾 指针
	receiver.tail = (receiver.tail + 1) & (len(receiver.elements) - 1)
	// TODO 判断是否需要扩容
	if receiver.tail == receiver.head {
		receiver.doubleCapacity()
	}

	return nil
}

// 将队首的元素删除，并返回该元素。如果队列为空，返回nil
func (receiver *ArrayQueue) Poll() interface{} {
	return receiver.pollFirst()
}

// 出队函数
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

// 查看队头元素
func (receiver *ArrayQueue) Peek() interface{} {
	return receiver.elements[receiver.head]
}

// 队列扩容函数
func (receiver *ArrayQueue) doubleCapacity() {
	oldCapacity := len(receiver.elements)
	newCapacity := oldCapacity << 1
	if newCapacity < 0 {
		panic("空间过大")
	}
	// 迁移元素
	// 迁移方法一
	newArr := make([]interface{}, newCapacity)
	for i := 0; i < oldCapacity; i++ {
		newArr[i] = receiver.elements[(i+receiver.head)&(oldCapacity-1)]
	}

	// 迁移方法二
	// arr := make([]interface{}, oldCapacity)
	// newArr := append(receiver.elements[receiver.head:], receiver.elements[:receiver.tail]...)
	// newArr = append(newArr, arr...)

	receiver.elements = newArr
	receiver.head = 0
	receiver.tail = oldCapacity
}

// 获取队列中元素数量
func (receiver *ArrayQueue) Size() int {
	return (receiver.tail - receiver.head) & (len(receiver.elements) - 1)
}
