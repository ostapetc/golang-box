package heap

import (
	"reflect"
	"testing"
)

//func TestMinHeap_GetLeafs(t *testing.T) {
//	//	     1
//	//	   /   \
//	//	  2     3
//	//   / \   / \
//	//  4  5  6   7
//
//	values := []int{1, 2, 3, 4, 5, 6, 7}
//	leafs  := []int {4, 5, 6, 7}
//
//	heap := NewMinHeap()
//	heap.SetValues(values)
//
//	fmt.Println("leafs", heap.GetLeafs())
//
//	if !reflect.DeepEqual(heap.GetLeafs(), leafs) {
//		t.Error("expected leafs", leafs, "got", heap.GetLeafs())
//	}
//
//	done("TestMinHeap_GetLeafs");
//}

//func TestMinHeap_Build(t *testing.T) {
//	size := 10
//	vals := []int{}
//
//	for i := size; i > 0; i-- {
//		vals = append(vals, i)
//	}
//
//	heap := NewMinHeap()
//	heap.SetValues(vals)
//
//	if heap.Valid() {
//		t.Error("Unexpected valid heap")
//	}
//
//	heap.Build()
//	heap.PrintTree()
//	fmt.Println(heap.values)
//
//	if !heap.Valid() {
//		t.Error("Unexpected invalid heap")
//	}
//}

func TestMinHeap_SetValues(t *testing.T) {
	heap := NewMinHeap()
	heap.SetValues([]int{1, 2, 3})

	expect := []int{0, 1, 2, 3}

	if !reflect.DeepEqual(expect, heap.values) {
		t.Error("For SetValues {1, 2, 3}, expected {0, 1, 2, 3} but got ", heap.values)
	}

	done("TestMinHeap_SetValues")
}

func TestMinHeap_Valid(t *testing.T) {
	type test struct {
		values []int
		expect bool
	}

	tests := []test{
		//	     1
		//	   /  \
		//	  2    3
		//  /  \   / \
		// 4   5  6   7
		{[]int{1, 2, 3, 4, 5, 6, 7}, true},
		//	     1
		//	   /  \
		//	  1    3
		//  /  \   / \
		// 4   5  6   7
		{[]int{1, 2, 3, 4, 5, 6, 7}, true},
		//	     2
		//	   /  \
		//	  1    3
		//  /  \   / \
		// 4   5  6   7
		{[]int{2, 1, 3, 4, 5, 6, 7}, false},
	}

	for _, test := range tests {
		heap := NewMinHeap()
		heap.SetValues(test.values)

		result := heap.Valid()

		if result != test.expect {
			t.Error("For heap", heap.GetValues(), "expected", test.expect, "but got", result)
		}
	}

	done("TestMinHeap_Valid")
}

func TestMinHeap_Append(t *testing.T) {
	heap := NewMinHeap()

	for i := 0; i < 100; i++ {
		heap.Append(i)

		if !heap.Valid() {
			t.Error("Heap is not valid")
		}
	}

	done("TestMinHeap_Append")
}
//
//func TestMinHeap_ExtractMax(t *testing.T) {
//	type test struct {
//		max    int
//		expect []int
//	}
//
//	// Init heap values
//	//	     100
//	//	   /     \
//	//	  70      40
//	//   /  \    /  \
//	//  30  25  10   5
//
//	heap := NewMinHeap()
//	heap.SetValues([]int{100, 70, 40, 30, 25, 10, 5})
//
//	tests := []test{
//		{
//			100,
//			[]int{70, 30, 40, 5, 25, 10},
//			//        70
//			//	   /     \
//			//	  30      40
//			//   /  \    /
//			//  5  25  10
//		},
//		{
//			70,
//			[]int{40, 30, 10, 5, 25},
//			//        40
//			//	   /     \
//			//	  30      10
//			//   /  \
//			//  5  25
//		},
//		{
//			40,
//			[]int{30, 25, 10, 5},
//			//        30
//			//	   /     \
//			//	  25      10
//			//   /
//			//  5
//		},
//		{
//			30,
//			[]int{25, 5, 10},
//			//        25
//			//	   /     \
//			//	  5      10
//		},
//		{
//			25,
//			[]int{10, 5},
//			//       10
//			//	   /
//			//	  5
//		},
//		{
//			10,
//			[]int{5},
//			//5
//		},
//		{
//			5,
//			[]int{},
//			//empty
//		},
//	}
//
//	for i, test := range tests {
//		max := heap.ExtractMin()
//
//		if max != test.max {
//			t.Error(i, "expected", test.max, "but got", max)
//		}
//
//		if !heap.Valid() {
//			t.Error(i, "Heap is not valid")
//		}
//
//		heapValues := heap.GetValues()
//		if !reflect.DeepEqual(test.expect, heapValues) {
//			t.Error(i, "Expected", test.expect, "got", heapValues)
//		}
//	}
//
//	done("TestMinHeap_ExtractMax")
//}