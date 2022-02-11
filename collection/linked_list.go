package collection

var _ list = (*LinkedList)(nil)

// LinkedList 链表
type LinkedList struct {
	head *node
	tail *node
	size int
}

// NewLinkedList 创建链表
func NewLinkedList() *LinkedList {
	head := newEmptyNode()
	tail := newEmptyNode()
	head.next = tail
	tail.prev = head

	return &LinkedList{
		head: head,
		tail: tail,
		size: 0,
	}
}

func (ll *LinkedList) Add(data interface{}) {
	defer func() {
		ll.size++
	}()

	_node := newNode(data)
	tail := ll.tail
	tail.next = _node
	_node.prev = _node
	ll.tail = tail
}

func (ll *LinkedList) Size() int {
	return ll.size
}
