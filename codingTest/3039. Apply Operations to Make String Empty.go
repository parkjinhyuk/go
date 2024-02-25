package codingTest

import (
	"math"
	"strings"
)

func lastNonEmptyString(s string) string {
	m := make(map[rune]int)
	for _, r := range s {
		v, ok := m[r]
		if ok {
			m[r] = v + 1
		} else {
			m[r] = 1
		}
	}

	min := math.MaxInt
	for _, v := range m {
		if v < min {
			min = v
		}
	}

	var lastStr string
	for len(s) != 0 {
		lastStr = s
		s, min = emptyString(&m, s, min)
	}

	return lastStr
}

func emptyString(m *map[rune]int, s string, min int) (string, int) {
	isAllMin := true
	for _, v := range *m {
		if v != min {
			isAllMin = false
			break
		}
	}
	if isAllMin && min != 1 {
		min = min - 1
	}

	for k, _ := range *m {
		s = strings.Replace(s, string(k), "", min)
	}

	nextMin := math.MaxInt
	for k, v := range *m {
		(*m)[k] = v - min
		if v-min == 0 {
			delete(*m, k)
			continue
		}

		if v-min < nextMin {
			nextMin = v - min
		}
	}

	return s, nextMin
}
