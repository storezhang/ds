package sort_test

import (
	`reflect`
	`testing`

	`github.com/storezhang/ds/sort`
)

func TestInsert(t *testing.T) {
	qsTests := []struct {
		in       []int
		expected []int
	}{
		{[]int{0}, []int{0}},
		{[]int{2, 8, 1}, []int{1, 2, 8}},
		{[]int{2, 2, 1}, []int{1, 2, 2}},
	}

	for _, test := range qsTests {
		in := test.in
		sort.Insert(test.in...)
		if !reflect.DeepEqual(test.in, test.expected) {
			t.Errorf("qs(%v) = %v；期望：%v", in, test.in, test.expected)
		}
	}
}

func BenchmarkInsert(b *testing.B) {
	for n := 0; n < b.N; n++ {
		items := []int{3, 432, 543, 89, 231312, 89, 431248, 2, 90, 7, 5, 9}
		sort.Insert(items...)
	}
}
