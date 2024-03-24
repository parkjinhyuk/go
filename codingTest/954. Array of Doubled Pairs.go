package codingTest

import "sort"

func canReorderDoubled(arr []int) bool {
	sort.Ints(arr)

	for {
		if len(arr) == 0 {
			return true
		}

		if arr[0]%2 == -1 {
			return false
		}

		target := 0
		if arr[0] < 0 {
			target = arr[0] / 2
		} else {
			target = arr[0] * 2
		}
		findIdx := binarySearch(arr, target)

		if findIdx == len(arr) {
			return false
		}

		arr = append(arr[1:findIdx], arr[findIdx+1:]...)
	}
}

func binarySearch(arr []int, target int) int {
	low := 1
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return len(arr)
}
