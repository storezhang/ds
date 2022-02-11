package sort

// Merge 归并
func Merge(items ...int) (merged []int) {
	size := len(items)
	merged = items
	if 1 == size {
		return
	}

	middle := size / 2
	left := Merge(items[:middle]...)
	right := Merge(items[middle:]...)
	merged = merge(left, right)

	return
}

func merge(left []int, right []int) (merged []int) {
	i := 0
	j := 0
	sizeLeft := len(left)
	sizeRight := len(right)
	merged = make([]int, 0, sizeLeft+sizeRight)
	for i < sizeLeft && j < sizeRight {
		if left[i] < right[j] {
			merged = append(merged, left[i])
			i++
		} else {
			merged = append(merged, right[j])
			j++
		}
	}
	merged = append(merged, left[i:]...)
	merged = append(merged, right[j:]...)

	return
}
