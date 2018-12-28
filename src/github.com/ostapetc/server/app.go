package main

import (
	"fmt"
)
func main() {
	fmt.Println(firstUniqChar("leetcode"))
	fmt.Println(firstUniqChar("artartm"))
}

func firstUniqChar(s string) int {
	cmap := getCountMap(s)

	for i, code := range s {
		if 	cmap[int(code)] < 2 {
			return i
		}
	}

	return - 1
}

func getCountMap(s string) map[int]int {
	cmap := make(map[int]int)

	for _, code := range s {
		cmap[int(code)]++
	}

	return cmap
}