package alchemicalreduction

import "unicode"

func wash(r []rune) ([]rune, int) {
	buff := make([]rune, 0)
	changes := 0

	for i, j := 0, 1; j <= len(r); j++ {
		if j < len(r) && (unicode.SimpleFold(r[i]) == r[j] || unicode.SimpleFold(r[j]) == r[i]) {
			changes++
			j++
		} else {
			buff = append(buff, r[i])
		}
		i = j
	}
	return buff, changes
}
