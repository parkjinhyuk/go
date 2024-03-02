package codingTest

func maxTurbulenceSize(arr []int) int {
	if len(arr) == 1 {
		return 1
	}

	max := 1
	for i := 0; i < len(arr)-1; i++ {
		m := 1
		b := true
		if arr[i] < arr[i+1] {
			b = false
		}
		for ; i < len(arr)-1; i++ {
			if arr[i] == arr[i+1] {
				break
			}

			if b2i(arr[i] < arr[i+1])^b2i(b) == b2i(true) {
				m++
				b = !b
			} else {
				i--
				break
			}
		}

		if max < m {
			max = m
		}
	}

	return max
}

func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}
