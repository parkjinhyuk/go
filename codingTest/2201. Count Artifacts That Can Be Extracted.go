package codingTest

func digArtifacts(n int, artifacts [][]int, dig [][]int) int {
	excavated := make([][]bool, n)
	for i := 0; i < n; i++ {
		excavated[i] = make([]bool, n)
	}

	for _, d := range dig {
		excavated[d[0]][d[1]] = true
	}

	answer := 0
	for _, artifact := range artifacts {
		if covered(excavated, artifact) {
			answer++
		}
	}
	return answer
}

func covered(excavated [][]bool, artifact []int) bool {
	for i := artifact[0]; i < artifact[2]; i++ {
		for j := artifact[1]; j < artifact[3]; j++ {
			if !excavated[i][j] {
				return false
			}
		}
	}
	return true
}
