package sort

// Heap 堆
type Heap struct {
	data []int
}

// NewHeap 创建堆
func NewHeap(items ...int) (heap *Heap) {
	heap = &Heap{
		data: make([]int, 0, len(items)),
	}
	for _, item := range items {
		heap.Add(item)
	}

	return
}

func (h *Heap) up(index int) {
	parent := h.parent(index)
	for index > 0 && h.less(index, parent) {
		h.swap(index, parent)
		index = parent
		parent = h.parent(index)
	}
}

func (h *Heap) Add(data int) {
	h.data = append(h.data, data)
	h.up(h.last())
}

func (h *Heap) Pop() (min int) {
	min = -1
	if 0 == len(h.data) {
		return
	}

	// 取出第一个元素，最小堆，一定是最小值
	min = h.data[h.first()]
	// 交换第一个元素和最后一个元素，因为要删除最后一个元素
	h.swap(h.first(), h.last())
	h.data = h.data[:h.last()]
	// 下沉，保证堆的性质（最小堆）
	h.down(h.first())

	return
}

func (h *Heap) down(parent int) {
	size := h.size()
	left := h.left(parent)
	// 判断是不是有子结点，因为左节点一定比右节点小，只要没有左节点那么一定就没有子节点
	for left < size {
		node := left
		right := h.right(parent)
		// 拿到左/右节点中的最小值
		if right < size && h.less(right, left) {
			node = right
		}

		// 判断当前节点是否小于父节点
		if h.less(parent, node) {
			break
		}
		h.swap(node, parent)
		parent = node
		left = h.left(parent)
	}
}

func (h *Heap) swap(left int, right int) {
	h.data[left], h.data[right] = h.data[right], h.data[left]
}

func (h *Heap) less(left int, right int) bool {
	return h.data[left] < h.data[right]
}

func (h *Heap) left(parent int) int {
	return parent*2 + 1
}

func (h *Heap) right(parent int) int {
	return parent*2 + 2
}

func (h *Heap) parent(index int) int {
	return (index - 1) / 2
}

func (h *Heap) first() int {
	return 0
}

func (h *Heap) last() int {
	return len(h.data) - 1
}

func (h *Heap) size() int {
	return len(h.data)
}
