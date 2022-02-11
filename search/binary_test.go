package search_test

import (
	`testing`

	`github.com/storezhang/ds/search`
)

func TestBinary(t *testing.T) {
	qsTests := []struct {
		in       []int
		target   int
		expected bool
	}{
		{[]int{0}, 0, true},
		{[]int{1, 2, 8}, 8, true},
		{[]int{1, 3, 5, 9, 17}, 10, false},
	}

	for _, test := range qsTests {
		actual := search.Binary(test.target, test.in...)
		if actual != test.expected {
			t.Errorf("qs(%v) = %v；期望：%v", test.in, actual, test.expected)
		}
	}
}
