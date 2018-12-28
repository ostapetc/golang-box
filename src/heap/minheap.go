package heap

//MinHeap is implementation of min heap
type MinHeap Heap

func NewMinHeap() *MinHeap {
	return &MinHeap{[]int{0}}
}

func (heap *MinHeap) GetLeafs() []int {
	return getLeafs(heap.values)
}

func (heap *MinHeap) Build() {
	//for i := 1; i < len(heap.values); i++ {
	//	heap.Swim(i)
	//}

	leafs := heap.GetLeafs()

	for _, i := range leafs {
		heap.Swim(i)
	}
}

func (heap *MinHeap) Heapify(i int) {
	left  := leftChild(heap.values, i)
	right := rightChild(heap.values, i)

	if heap.values[i] <= left && heap.values[i] <= right {
		return
	}

	var chidx int

	if left < right {
		chidx = leftIdx(i)
	} else {
		chidx = rightIdx(i)
	}

	swap(heap.values, i, chidx)
	heap.Heapify(chidx)
}

func (heap *MinHeap) Valid() bool {
	values := heap.values

	for i := 1; i < len(values); i++ {
		val := values[i]

		left := leftChild(heap.values, i)
		if left != 0 && left < val {
			return false
		}

		right := rightChild(heap.values, i)
		if right != 0 && right < val {
			return false
		}
	}

	return true
}

func (heap *MinHeap) SetValues(values []int) {
	heap.values = setValues(values)
}

func (heap *MinHeap) Append(val int) {
	heap.values = append(heap.values, val)
	heap.Swim(len(heap.values) - 1)
}

func (heap *MinHeap) GetValues() []int {
	return getValues(heap.values)
}

func (heap *MinHeap) ExtractMin() int {
	max  := heap.values[1]
	last := heap.values[len(heap.values) - 1]

	heap.values[1] = last
	heap.values = heap.values[:len(heap.values) - 1]

	if len(heap.values) > 1 {
		heap.Heapify(1)
	}

	return max
}

func (heap *MinHeap) Swim(i int) {
	parent := parentIdx(i)

	for i > 1 && heap.values[i] < heap.values[parent] {
		swap(heap.values, i, parent)
		i = parent
		parent = parentIdx(i)
	}
}

func (heap *MinHeap) PrintTree() {
	printTree(heap.values, 1, 0)
}