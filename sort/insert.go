package sort

// Insert 插件
func Insert(items ...int) {
	size := len(items)
	for i := 0; i < size; i++ {
		for j := i; j > 0 && items[j] < items[j-1]; j-- {
			items[j], items[j-1] = items[j-1], items[j]
		}
	}
}
