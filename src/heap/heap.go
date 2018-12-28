package heap

import (
	"fmt"
	"strings"
)

//Heap is a base heap structure
type Heap struct {
	values []int
}


//MaxHeap is implementation of max heap
type MaxHeap Heap

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{[]int{0}}
}

func (heap *MaxHeap) Swap(i, j int) {
	swap(heap.values, i, j)
}

func (heap *MaxHeap) GetLeafs() []int {
	return getLeafs(heap.values)
}

func (heap *MaxHeap) Build() {
	leafs := heap.GetLeafs()

	for _, i := range leafs {
		heap.Sweam(i)
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

	heap.Swap(i, chidx)
	heap.Heapify(chidx)
}

func MaxHeapValid(heap *MaxHeap) bool {
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

//MinHeap is implementation of min heap
type MinHeap Heap

func NewMinHeap() *MinHeap {
	return &MinHeap{[]int{0}}
}

func (heap *MinHeap) Swap(i, j int) {
	swap(heap.values, i, j)
}

func (heap *MinHeap) GetLeafs() []int {
	return getLeafs(heap.values)
}

func (heap *MaxHeap) SetValues(values []int) {
	heap.values = []int{}
	heap.values = append(heap.values, 0)

	for _, val := range values {
		heap.values = append(heap.values, val)
	}
}

func (heap *MaxHeap) GetValues() []int {
	return heap.values[1:]
}

func (heap *MaxHeap) Append(val int) {
	heap.values = append(heap.values, val)
	heap.Sweam(len(heap.values) - 1)
}

func (heap *Heap) Print() {
	for i, val := range heap.values {
		if i == 0 { continue }

		parent := parentIdx(i)

		if parent > 0 {
			fmt.Println(val, "parent", heap.values[parent])
		} else {
			fmt.Println(val)
		}
	}
}

func (heap *Heap) PrintTree() {
	printTree(heap, 1, 0)
}

func printTree(heap *Heap, index, padding int) {
	if len(heap.values) < 2 {
		return
	}

	val := heap.values[index]

	fmt.Println(strings.Repeat("-", padding), val)

	left  := leftIdx(index)
	right := rightIdx(index)

	if right < len(heap.values) {
		printTree(heap, right, padding + 3)
	}

	if left < len(heap.values) {
		printTree(heap, left, padding + 3)
	}
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

func (heap *MaxHeap) Sweam(i int) {
	parent := parentIdx(i)

	for i > 1 && heap.values[i] > heap.values[parent] {
		heap.Swap(i, parent)
		i = parent
		parent = parentIdx(i)
	}
}

func (heap *MinHeap) Sweam(i int) {
	parent := parentIdx(i)

	for i > 1 && heap.values[i] > heap.values[parent] {
		heap.Swap(i, parent)
		i = parent
		parent = parentIdx(i)
	}
}

func (heap *Heap) Sweam(i int) {
	panic("needs override")
}

func leftIdx(i int) int {
	return i * 2
}

func rightIdx(i int) int {
	return i * 2 + 1
}

func parentIdx(i int) int {
	return i / 2
}

func leftChild(values []int, i int) int {
	left := leftIdx(i)

	if len(values) > left {
		return values[left]
	}

	return 0
}

func rightChild(values []int, i int) int {
	right := rightIdx(i)

	if len(values) > right {
		return values[right]
	}

	return 0
}

func swap(values []int, i, j int) {
	tmp := values[i]
	values[i] = values[j]
	values[j] = tmp
}

func getLeafs(values []int) []int {
	return values[len(values) / 2:]
}