package codingTest

func maximumJumps(nums []int, target int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums)-1; i++ {
		if i != 0 && dp[i] == 0 {
			continue
		}

		for j := i + 1; j < len(nums); j++ {
			if abs(nums[j]-nums[i]) <= target {
				dp[j] = max(dp[j], dp[i]+1)
			}
		}
	}

	if dp[len(nums)-1] == 0 {
		return -1
	}
	return dp[len(nums)-1]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
