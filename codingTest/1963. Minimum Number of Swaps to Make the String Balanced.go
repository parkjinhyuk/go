package codingTest

type stack []int32

func (s stack) Push(v int32) stack {
	return append(s, v)
}

func (s stack) Pop() stack {
	l := len(s)
	return s[:l-1]
}

func minSwaps(s string) int {
	swaps := 0

	var st stack
	for _, v := range s {
		if v == '[' {
			st = st.Push(v)
			continue
		}

		if len(st) > 0 {
			st = st.Pop()
			continue
		}

		swaps++
		st = st.Push('[')
	}

	return swaps
}
