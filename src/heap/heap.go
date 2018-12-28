package heap

import (
	"fmt"
	"strings"
)

type Heap struct {
	values []int
}

func printTree(values []int, index, padding int) {
	if len(values) < 2 {
		return
	}

	val := values[index]

	fmt.Println(strings.Repeat("-", padding), val)

	left  := leftIdx(index)
	right := rightIdx(index)

	if right < len(values) {
		printTree(values, right, padding + 3)
	}

	if left < len(values) {
		printTree(values, left, padding + 3)
	}
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
	start := len(values) / 2
	if len(values) % 2 != 0 {
		start++
	}

	return values[start:]
}

func setValues(values []int) []int {
	result := []int{0}

	for _, val := range values {
		result = append(result, val)
	}

	return result
}

func getValues(values []int) []int {
	return values[1:]
}