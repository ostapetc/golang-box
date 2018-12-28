package heap

import (
	"fmt"
	"reflect"
	"testing"
)

func done(name string) {
	if false {
		fmt.Println("Done " + name)
	}
}

func TestMaxHeap_GetLeafs(t *testing.T) {
	type test struct {
		values []int
		expect []int
	}

	tests := []test{
		//	     100
		//	   /     \
		//	  70      40
		//   /  \    /  \
		//  30  25  10   5
		{[]int{100, 70, 40, 30, 25, 10, 5}, []int {30, 25, 10, 5}},
		//	     100
		//	   /     \
		//	  70      40
		//   /  \    /
		//  30  25  10
		{[]int{100, 70, 40, 30, 25, 10}, []int {30, 25, 10}},
	}

	for _, test := range tests {
		heap := NewMaxHeap()
		heap.SetValues(test.values)

		if !reflect.DeepEqual(heap.GetLeafs(), test.expect) {
			t.Error("expected leafs", test.expect, "got", heap.GetLeafs())
		}
	}

	done("TestMaxHeap_GetLeafs")
}

func TestMaxHeap_Build(t *testing.T) {
	size := 10
	vals := []int{}

	for i := 1; i <= size; i++ {
		vals = append(vals, i)
	}

	heap := NewMaxHeap()
	heap.SetValues(vals)

	if heap.Valid() {
		t.Error("Unexpected valid heap")
	}

	heap.Build()

	if !heap.Valid() {
		t.Error("Unexpected invalid heap")
	}

	done("TestMaxHeap_Build")
}

func TestMaxHeap_SetValues(t *testing.T) {
	heap := NewMaxHeap()
	heap.SetValues([]int{1, 2, 3})

	expect := []int{0, 1, 2, 3}

	if !reflect.DeepEqual(expect, heap.values) {
		t.Error("For SetValues {1, 2, 3}, expected {0, 1, 2, 3} but got ", heap.values)
	}

	done("TestMaxHeap_SetValues")
}

func TestMaxHeap_Valid(t *testing.T) {
	type test struct {
		values []int
		expect bool
	}

	tests := []test{
		//	     100
		//	   /     \
		//	  70      40
		//   /  \    /  \
		//  30  25  10   5
		{[]int{100, 70, 40, 30, 25, 10, 5}, true},

		//	     100
		//	   /     \
		//	  100      40
		//   /  \    /  \
		//  30  25  10   5
		{[]int{100, 100, 40, 30, 25, 10, 5}, true},

		//	     45
		//	   /    \
		//	  70     40
		//   /  \    /  \
		//  30  25  10   5
		{[]int{45, 70, 40, 30, 25, 10, 5}, false},
	}

	for _, test := range tests {
		heap := NewMaxHeap()
		heap.SetValues(test.values)

		result := heap.Valid()

		if result != test.expect {
			t.Error("For heap", heap.GetValues(), "expected", test.expect, "but got", result)
		}
	}

	done("TestMaxHeap_Valid")
}

func TestMaxHeap_Append(t *testing.T) {
	heap := NewMaxHeap()

	for i := 0; i < 100; i++ {
		heap.Append(i)

		if !heap.Valid() {
			t.Error("Heap is not valid")
		}
	}

	done("TestMaxHeap_Append")
}

func TestMaxHeap_ExtractMax(t *testing.T) {
	type test struct {
		max    int
		expect []int
	}

	// Init heap values
	//	     100
	//	   /     \
	//	  70      40
	//   /  \    /  \
	//  30  25  10   5

	heap := NewMaxHeap()
	heap.SetValues([]int{100, 70, 40, 30, 25, 10, 5})

	tests := []test{
		{
			100,
			[]int{70, 30, 40, 5, 25, 10},
			//        70
			//	   /     \
			//	  30      40
			//   /  \    /
			//  5  25  10
		},
		{
			70,
			[]int{40, 30, 10, 5, 25},
			//        40
			//	   /     \
			//	  30      10
			//   /  \
			//  5  25
		},
		{
			40,
			[]int{30, 25, 10, 5},
			//        30
			//	   /     \
			//	  25      10
			//   /
			//  5
		},
		{
			30,
			[]int{25, 5, 10},
			//        25
			//	   /     \
			//	  5      10
		},
		{
			25,
			[]int{10, 5},
			//       10
			//	   /
			//	  5
		},
		{
			10,
			[]int{5},
			//5
		},
		{
			5,
			[]int{},
			//empty
		},
	}

	for i, test := range tests {
		max := heap.ExtractMax()

		if max != test.max {
			t.Error(i, "expected", test.max, "but got", max)
		}

		if !heap.Valid() {
			t.Error(i, "Heap is not valid")
		}

		heapValues := heap.GetValues()
		if !reflect.DeepEqual(test.expect, heapValues) {
			t.Error(i, "Expected", test.expect, "got", heapValues)
		}
	}

	done("TestMaxHeap_ExtractMax")
}