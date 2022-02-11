package collection

import (
	`fmt`
	`testing`
)

func TestLinkedList(t *testing.T) {
	_list := NewLinkedList()
	_list.Add(`store`)
	_list.Add(`zhang`)
	fmt.Println(_list)
}
