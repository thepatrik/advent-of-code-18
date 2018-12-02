package cronalcalibration

import "strconv"

// FindDupFreq : Find the first duplicate frequence
func FindDupFreq(snums []string) int {
	current := 0
	freqs := make(map[int]bool)
	freqs[current] = true
	i := 0
	for {
		v, _ := strconv.Atoi(snums[i])
		current = current + v
		if _, ok := freqs[current]; ok == true {
			return current
		}

		if i++; i == len(snums) {
			i = 0 // index has wrapped around
		}
		freqs[current] = true
	}
}
