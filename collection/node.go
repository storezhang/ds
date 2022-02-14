package collection

type node struct {
	prev *node
	data interface{}
	next *node
}

func newNode(data interface{}) *node {
	return &node{
		data: data,
	}
}
