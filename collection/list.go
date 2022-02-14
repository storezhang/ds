package collection

type list interface {
	collection

	// Add 最前添加
	Add(item interface{})

	// Append 最后添加
	Append(item interface{})

	// Reverse 反转
	Reverse()

	// Foreach 遍历数据
	Foreach(fun foreachList)
}
