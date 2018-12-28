package maxheap

import (
	"fmt"
	"strings"
)

type MaxHeap struct {
	values []int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{[]int{0}}
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

func (heap *MaxHeap) Print() {
	for i, val := range heap.values {
		if i == 0 { continue }

		parent := heap.ParentIdx(i)

		if parent > 0 {
			fmt.Println(val, "parent", heap.values[parent])
		} else {
			fmt.Println(val)
		}
	}
}

func (heap *MaxHeap) PrintTree() {
	printTree(heap, 1, 0)
}

func printTree(heap *MaxHeap,index, padding int) {
	if len(heap.values) < 2 {
		return
	}

	val := heap.values[index]

	fmt.Println(strings.Repeat("-", padding), val)

	left  := heap.LeftIdx(index)
	right := heap.RightIdx(index)

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

func (heap *MaxHeap) Heapify(i int) {
	left  := heap.LeftChild(i)
	right := heap.RightChild(i)

	if heap.values[i] >= left && heap.values[i] >= right {
		return
	}

	var chidx int

	if left > right {
		chidx = heap.LeftIdx(i)
	} else {
		chidx = heap.RightIdx(i)
	}

	heap.Swap(i, chidx)
	heap.Heapify(chidx)
}

func (heap *MaxHeap) Sweam(i int) {
	parent := heap.ParentIdx(i)

	for i > 1 && heap.values[i] > heap.values[parent] {
		heap.Swap(i, parent)
		i = parent
		parent = heap.ParentIdx(i)
	}
}

func (heap *MaxHeap) Swap(i, j int) {
	tmp := heap.values[i]
	heap.values[i] = heap.values[j]
	heap.values[j] = tmp
}

func (heap *MaxHeap) LeftChild(i int) int {
	left := heap.LeftIdx(i)

	if len(heap.values) > left {
		return heap.values[left]
	}

	return 0
}

func (heap *MaxHeap) RightChild(i int) int {
	right := heap.RightIdx(i)

	if len(heap.values) > right {
		return heap.values[right]
	}

	return 0
}

func (heap *MaxHeap) LeftIdx(i int) int {
	return i * 2
}

func (heap *MaxHeap) RightIdx(i int) int {
	return i * 2 + 1
}

func (heap *MaxHeap) ParentIdx(i int) int {
	return i / 2
}

