package sort

func Quick(start int, end int, items ...int) {
	if start >= end {
		return
	}

	_pivot := partition(start, end, items...)
	Quick(start, _pivot-1, items...)
	Quick(_pivot+1, end, items...)
}

func partition(low int, high int, items ...int) (pivot int) {
	pivot = items[low]
	for low < high {
		for low < high && pivot <= items[high] {
			high--
		}
		items[low] = items[high]

		for low < high && pivot >= items[low] {
			low++
		}
		items[high] = items[low]
	}
	items[low] = pivot
	pivot = low

	return
}
