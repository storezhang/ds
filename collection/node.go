package collection

type node struct {
	data interface{}
	prev *node
	next *node
}

func newNode(data interface{}, prev *node, next *node) *node {
	return &node{
		data: data,
		prev: prev,
		next: next,
	}
}

func newEmptyNode() *node {
	return new(node)
}
