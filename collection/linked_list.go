package collection

var _ = (*LinkedList)(nil)

// LinkedList 链表
type LinkedList struct {
	header *node
	tail   *node
	size   int
}

// NewLinkedList 创建链表
func NewLinkedList() *LinkedList {
	return &LinkedList{
		header: newEmptyNode(),
		tail:   newEmptyNode(),
		size:   0,
	}
}

func (ll *LinkedList) Add(data interface{}) {
	defer func() {
		ll.size++
	}()

	_node := newNode(data, ll.header, ll.tail)
	ll.header.next = _node
}

func (ll *LinkedList) Size() int {
	return ll.size
}
