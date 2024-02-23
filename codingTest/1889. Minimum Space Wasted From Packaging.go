package codingTest

import (
	"math"
	"sort"
)

func minWastedSpace(packages []int, boxes [][]int) int {
	sort.Ints(packages)
	for _, box := range boxes {
		sort.Ints(box)
	}

	sum := make([]int, len(packages))
	for i, p := range packages {
		if i == 0 {
			sum[i] = p
			continue
		}

		sum[i] = p + sum[i-1]
	}

	result := math.MaxInt
	for _, box := range boxes {
		wastedSpace := 0
		idx, prevIdx := 0, 0
		for _, b := range box {
			idx = upperBound(&packages, b)
			if idx == 0 {
				continue
			}

			if prevIdx == 0 {
				wastedSpace += b*idx - sum[idx-1]
				prevIdx = idx
				continue
			}

			wastedSpace += b*(idx-prevIdx) - (sum[idx-1] - sum[prevIdx-1])
			prevIdx = idx

			if idx == len(packages) {
				break
			}
		}

		if idx != len(packages) {
			continue
		}

		if wastedSpace < result {
			result = wastedSpace
		}
	}

	if result == math.MaxInt {
		return -1
	}

	if result > 1000000007 {
		result = result % 1000000007
	}

	return result
}

func upperBound(packages *[]int, b int) int {
	low, high, mid := 0, len(*packages)-1, 0
	for low <= high {
		mid = (low + high) / 2
		if (*packages)[mid] > b {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return low
}
