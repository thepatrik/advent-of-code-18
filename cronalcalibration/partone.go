package cronalcalibration

import "strconv"

// SumFreq : Get the sum of all frequencies
func SumFreq(s []string, i int) int {
	v, _ := strconv.Atoi(s[i])
	if i == len(s)-1 {
		return v
	}
	return v + SumFreq(s, i+1)
}
