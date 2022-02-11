package sort

// Bubble 冒泡
func Bubble(items ...int) {
	size := len(items)
	for i := 0; i < size; i++ {
		for j := 0; j < size-i-1; j++ {
			if items[j] > items[j+1] {
				items[j], items[j+1] = items[j+1], items[j]
			}
		}
	}
}
