package heap

//MaxHeap is implementation of max heap
type MaxHeap Heap

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{[]int{0}}
}

func (heap *MaxHeap) GetLeafs() []int {
	return getLeafs(heap.values)
}

func (heap *MaxHeap) Build() {
	leafs := heap.GetLeafs()

	for _, i := range leafs {
		heap.Swim(i)
	}
}

func (heap *MaxHeap) Heapify(i int) {
	left  := leftChild(heap.values, i)
	right := rightChild(heap.values, i)

	if heap.values[i] >= left && heap.values[i] >= right {
		return
	}

	var chidx int

	if left > right {
		chidx = leftIdx(i)
	} else {
		chidx = rightIdx(i)
	}

	swap(heap.values, i, chidx)
	heap.Heapify(chidx)
}

func (heap *MaxHeap) Valid() bool {
	values := heap.values

	for i := 1; i < len(values); i++ {
		val := values[i]

		left := leftChild(heap.values, i)
		if left != 0 && left > val {
			return false
		}

		right := rightChild(heap.values, i)
		if right != 0 && right > val {
			return false
		}
	}

	return true
}

func (heap *MaxHeap) SetValues(values []int) {
	heap.values = setValues(values)
}

func (heap *MaxHeap) Append(val int) {
	heap.values = append(heap.values, val)
	heap.Swim(len(heap.values) - 1)
}

func (heap *MaxHeap) GetValues() []int {
	return getValues(heap.values)
}

func (heap *MaxHeap) ExtractMax() int {
	max  := heap.values[1]
	last := heap.values[len(heap.values) - 1]

	heap.values[1] = last
	heap.values = heap.values[:len(heap.values) - 1]

	if len(heap.values) > 1 {
		heap.Heapify(1)
	}

	return max
}

func (heap *MaxHeap) Swim(i int) {
	parent := parentIdx(i)

	for i > 1 && heap.values[i] > heap.values[parent] {
		swap(heap.values, i, parent)
		i = parent
		parent = parentIdx(i)
	}
}

func (heap *MaxHeap) PrintTree() {
	printTree(heap.values, 1, 0)
}