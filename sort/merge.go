package sort

// Merge 归并
func Merge(items ...int) (merged []int) {
	size := len(items)
	merged = items
	if 1 == size {
		return
	}

	middle := size / 2
	lefts := Merge(items[:middle]...)
	rights := Merge(items[middle:]...)
	merged = merge(lefts, rights)

	return
}

func merge(lefts []int, rights []int) (merged []int) {
	left := 0
	right := 0
	sizeLeft := len(lefts)
	sizeRight := len(rights)
	merged = make([]int, 0, sizeLeft+sizeRight)
	for left < sizeLeft && right < sizeRight {
		if lefts[left] < rights[right] {
			merged = append(merged, lefts[left])
			left++
		} else {
			merged = append(merged, rights[right])
			right++
		}
	}
	merged = append(merged, lefts[left:]...)
	merged = append(merged, rights[right:]...)

	return
}
