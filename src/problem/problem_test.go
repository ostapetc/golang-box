package fairrations

import (
	"fmt"
	"testing"
)

type test struct {
	subjects []int32
	expect   int32
}

var tests = []test{
	{[]int32{4,5,6,7}, 4},
	{[]int32{2,3,4,5,6}, 4},
	{[]int32{1,2}, 0},
	{[]int32{1,1,1}, 0},
}

func TestFairRations(t *testing.T) {
	for _, test := range tests {
		input := make([]int32, len(test.subjects))
		copy(input, test.subjects)

		result := FairRations(input)

		if result != test.expect {
			fmt.Println("for", test.subjects, "expected", test.expect, "got", result)
			t.Fail()
		}
	}
}