package main

import (
	"fmt"
	"sort"
)

type stat struct {
	currVal int64
	initVal int64
}

func main() {
	machines := []int64{5,1,4,3,2,7}
	stats := []stat{}

	for _, machine := range machines {
		stats = append(stats, stat{machine, machine})
	}

	sort.Slice(stats, func(i, j int) bool {
		return stats[i].initVal < stats[j].initVal
	})

	fmt.Println(stats)

	for t := 0; t < 5; t++ {
		s := stats[0]
		stats = stats[1:]
		s.currVal += s.initVal;
		insertDay(stats, s)

		fmt.Println(stats)
	}
}

func insertDay(stats []stat, s stat) []stat {
	stats = append(stats, s)
	j := len(stats) - 1

	for j > 0 && stats[j].currVal < stats[j - 1].currVal {
		swap(stats, j, j - 1)
		j--
	}

	return stats
}

func swap(stats []stat, i, j int) {
	tmp := stats[i]
	stats[i] = stats[j]
	stats[j] = tmp
}