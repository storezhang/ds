package search

// Binary 二分
func Binary(target int, items ...int) (searched bool) {
	left := 0
	right := len(items)
	for left < right {
		index := left + (right-left)/2
		middle := items[index]
		if target > middle {
			left = index + 1
		} else if target < middle {
			right = index - 1
		} else {
			searched = true
		}

		if searched {
			return
		}
	}

	return
}
