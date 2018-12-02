package cronalcalibration

import "strconv"

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
