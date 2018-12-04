package nomatterhowyousliceit

import (
	"regexp"
	"strconv"
	"strings"
)

const size int = 1000

// GetFabric returns the fabric to work on
func GetFabric(slice []string) [size][size]int {
	var fabric [size][size]int
	for _, s := range slice {
		_, x, y, w, h := GetValues(s)

		for dx := 0; dx < w; dx++ {
			for dy := 0; dy < h; dy++ {
				fabric[x+dx][y+dy]++
			}
		}
	}
	return fabric
}

// GetValues parses and splits the string into ints
func GetValues(s string) (int, int, int, int, int) {
	split := regexp.MustCompile("[^0-9]+").Split(strings.TrimLeft(s, "#"), 5)
	n, _ := strconv.Atoi(split[0])
	x, _ := strconv.Atoi(split[1])
	y, _ := strconv.Atoi(split[2])
	w, _ := strconv.Atoi(split[3])
	h, _ := strconv.Atoi(split[4])
	return n, x, y, w, h
}
