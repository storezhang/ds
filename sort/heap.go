package sort

func Heap(items ...int) {

}

type heap []int

func (h heap) swap(left int, right int) {
	h[left], h[right] = h[right], h[left]
}

func (h heap) less(left int, right int) bool {
	return h[left] < h[right]
}

func (h heap) up(index int) {
	for {
		parent := (index - 1) / 2
		if parent == index || h.less(index, parent) {
			break
		}
		h.swap(index, parent)
		index = parent
	}
}

func (h *heap) add(data int) {
	*h = append(*h, data)
	h.up(len(*h) - 1)
}
