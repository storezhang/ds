package collection

import (
	`sync`
)

var _ list = (*LinkedList)(nil)

// LinkedList 链表
type LinkedList struct {
	head *node
	tail *node

	mutex *sync.RWMutex
	size  int
}

// NewLinkedList 创建链表
func NewLinkedList() *LinkedList {
	return &LinkedList{
		head: nil,
		tail: nil,

		mutex: new(sync.RWMutex),
		size:  0,
	}
}

func (ll *LinkedList) Add(item interface{}) {
	_node := newNode(item)
	if 0 == ll.size {
		ll.head = _node
		ll.tail = _node
	} else {
		ll.head.prev = _node
		_node.next = ll.head
		_node.prev = nil
		ll.head = _node
	}

	ll.size++
}

func (ll *LinkedList) Append(item interface{}) {
	_node := newNode(item)
	if 0 == ll.size {
		ll.head = _node
		ll.tail = _node
	} else {
		ll.tail.next = _node
		_node.prev = ll.tail
		_node.next = nil
		ll.tail = _node
	}
}

func (ll *LinkedList) Get(index int) (item interface{}) {
	ll.mutex.RLock()
	defer ll.mutex.RUnlock()

	// 循环遍历取数据
	_node := ll.head
	for i := 0; i < index && nil != _node; i++ {
		_node = _node.next
	}
	item = _node.data

	return
}

func (ll *LinkedList) Foreach(fun foreachList) {
	_node := ll.head
	for index := 0; index < ll.size && nil != _node; index++ {
		fun(index, _node.data)
		_node = _node.next
	}
}

func (ll *LinkedList) Size() int {
	ll.mutex.RLock()
	defer ll.mutex.RUnlock()

	return ll.size
}
