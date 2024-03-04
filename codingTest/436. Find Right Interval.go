package codingTest

import "sort"

type IntervalWithIdx struct {
	idx      int
	interval []int
}

func findRightInterval(intervals [][]int) []int {
	intervalsWithIdx := make([]IntervalWithIdx, len(intervals))
	for i, v := range intervals {
		intervalsWithIdx[i].idx = i
		intervalsWithIdx[i].interval = v
	}

	sort.Slice(intervalsWithIdx, func(i, j int) bool {
		return intervalsWithIdx[i].interval[0] < intervalsWithIdx[j].interval[0]
	})

	answer := make([]int, len(intervals))
	for _, v := range intervalsWithIdx {
		answer[v.idx] = lowerBound(intervalsWithIdx, v.interval[1])
	}

	return answer
}

func lowerBound(intervalsWithIdx []IntervalWithIdx, i int) int {
	low, high, mid := 0, len(intervalsWithIdx)-1, 0
	for low <= high {
		mid = (low + high) / 2
		if intervalsWithIdx[mid].interval[0] >= i {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	if low == len(intervalsWithIdx) {
		return -1
	}

	return intervalsWithIdx[low].idx
}
