package arrays

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)


func TestBuildArrayOfThreeStrings(t *testing.T)  {
	arr := BuildArrayOfThreeStrings()
	assert.Equal(t, "first", arr[0])
	assert.Equal(t, "second", arr[1])
	assert.Equal(t, "third", arr[2])
}

func TestBuildSliceOfIntegers(t *testing.T) {
	size  := 5
	slice := BuildSliceOfIntegers(size)
	fmt.Println(slice)
	assert.Equal(t, 1, 1)
}