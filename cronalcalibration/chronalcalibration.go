package cronalcalibration

import (
	"strconv"
)

// SumFreq : Get the sum of all frequencies
func SumFreq(s []string, i int) int {
	v, _ := strconv.Atoi(s[i])
	if i == len(s)-1 {
		return v
	}
	return v + SumFreq(s, i+1)
}

// FindDupFreq : Find the first duplicate frequence
func FindDupFreq(snums []string) int {
	current := 0
	freqs := make([]int, 0)
	freqs = append(freqs, current)
	ix := 0
	for {
		s := snums[ix]
		v, _ := strconv.Atoi(s)
		current = current + v
		if findIx(current, freqs) {
			return current
		}
		ix++
		if ix == len(snums) {
			ix = 0 // index has wrapped around
		}
		freqs = append(freqs, current)
	}
}

func findIx(v int, ixs []int) bool {
	for _, ix := range ixs {
		if v == ix {
			return true
		}
	}
	return false
}
