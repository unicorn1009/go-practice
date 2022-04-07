package myutil

import (
	"errors"
)

type QueueNode struct {
	element interface{}
	next    *QueueNode
	prev    *QueueNode
}

type LinkedQueue struct {
	first *QueueNode // 队头，出队端
	last  *QueueNode // 队尾，入队端
	size  int
}

// 入队，将元素添加到队尾，如果成功，则返回true。
func (receiver *LinkedQueue) Offer(e interface{}) bool {
	if err := receiver.addLast(e); err != nil {
		return false
	}
	return true
}

// 入队函数
func (receiver *LinkedQueue) addLast(e interface{}) error {
	if e == nil {
		return errors.New("元素不可为nil")
	}
	oldLast := receiver.last
	newNode := QueueNode{
		element: e,
		next:    nil,
		prev:    oldLast,
	}

	if oldLast == nil {
		receiver.first = &newNode
	} else {
		oldLast.next = &newNode
	}
	receiver.last = &newNode

	receiver.size++
	return nil
}

// 将队首的元素删除，并返回该元素。如果队列为空，返回nil
func (receiver *LinkedQueue) Poll() interface{} {
	if receiver.first == nil {
		return nil
	} else {
		return receiver.pollFirst()
	}
}

// 出队函数
func (receiver *LinkedQueue) pollFirst() interface{} {
	// assert receiver.first != nil
	f := receiver.first
	ele := f.element
	next := f.next
	f.element = nil
	f.next = nil
	receiver.first = next
	if next == nil {
		receiver.last = nil
	} else {
		next.prev = nil
	}
	receiver.size--
	return ele

}

// 查看队头元素
func (receiver *LinkedQueue) Peek() interface{} {
	if receiver.first == nil {
		return nil
	}
	return receiver.first.element
}

// 获取队列中元素数量
func (receiver *LinkedQueue) Size() int {
	return receiver.size
}
