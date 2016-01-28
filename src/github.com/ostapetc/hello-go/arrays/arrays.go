package arrays

func BuildArrayOfThreeStrings() [3]string {
	var arr [3]string
	arr[0] = "first"
	arr[1] = "second"
	arr[2] = "third"

	return arr
}

func BuildSliceOfIntegers(size int) []int {
	s := []int{}

	for i := 0; i < size; i++ {
		s[i] = i
	}

	return s
}


