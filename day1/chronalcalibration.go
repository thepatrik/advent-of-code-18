package cronalcalibration

import (
	"strconv"
)

func freq(s []string, i int) int {
	v, _ := strconv.Atoi(s[i])
	if i == len(s)-1 {
		return v
	}
	return v + freq(s, i+1)
}
