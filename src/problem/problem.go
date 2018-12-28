package problem

import (
	"sort"
)

type stat struct {
	currVal int64
	initVal int64
}


// Complete the minTime function below.
func MinTime(machines []int64, goal int64) int64 {
	done := int64(0)
	day  := int64(0)
	//days := getDays(machines)

	stats := []stat{}
	for _, machine := range machines {
		stats = append(stats, stat{machine, machine})
	}

	sort.Slice(stats, func(i, j int) bool {
		return stats[i].initVal < stats[j].initVal
	})

	for done != goal {
		s := stats[0]
		stats = stats[1:]
		day = s.currVal
		s.currVal += s.initVal;
		stats = insertDay(stats, s)
		done++
	}

	return day
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