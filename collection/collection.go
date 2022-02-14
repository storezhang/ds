package collection

type collection interface {
	// Get 取数据
	Get(index int) interface{}

	// Size 大小
	Size() int
}
