package sort

func Quick(left int, right int, items ...int) {
	if left >= right {
		return
	}

	_pivot := partition(left, right, items...)
	Quick(left, _pivot-1, items...)
	Quick(_pivot+1, right, items...)
}

func partition(left int, right int, items ...int) (pivot int) {
	pivot = items[left]
	for left < right {
		for left < right && pivot <= items[right] {
			right--
		}
		items[left] = items[right]

		for left < right && pivot >= items[left] {
			left++
		}
		items[right] = items[left]
	}
	items[left] = pivot
	pivot = left

	return
}
