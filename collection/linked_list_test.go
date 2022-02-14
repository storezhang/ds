package collection

import (
	`fmt`
	`testing`
)

func TestLinkedList(t *testing.T) {
	_list := NewLinkedList()
	_list.Add(`store`)
	_list.Add(`zhang`)
	_list.Reverse()
	_list.Foreach(func(index int, value interface{}) {
		fmt.Println(index, value)
	})
	fmt.Println(_list.Get(1))
	fmt.Println(_list)
}
