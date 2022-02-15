package sort_test

import (
	`fmt`
	`testing`

	`github.com/storezhang/ds/sort`
)

func TestHeap(t *testing.T) {
	qsTests := []struct {
		in       []int
		expected []int
	}{
		// {[]int{0}, []int{0}},
		{[]int{2, 8, 1, 89, 32, 89, 354}, []int{1, 2, 8}},
		{[]int{2, 2, 1}, []int{1, 2, 2}},
	}

	for _, test := range qsTests {
		// in := test.in
		heap := sort.NewHeap(test.in...)
		for i := 0; i < 30; i++ {
			fmt.Println(heap.Pop())
		}
	}
}
