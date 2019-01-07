package chronalcoordinates

type coord struct {
	x int
	y int
}

// LargestArea calculates the largest area (that isn't infinate)
func LargestArea(coords []coord, w, h int) (int, int) {
	maxSum := 10000
	inf := make(map[coord]bool)
	m := make(map[coord]int)
	regions := 0

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			mc := coord{0, 0}
			min := -1
			total := 0
			for _, c := range coords {
				dist := abs(x-c.x) + abs(y-c.y)
				if dist < min || min == -1 {
					min = dist
					mc = c
				} else if dist == min {
					mc = coord{-1, -1}
				}

				total += abs(x-c.x) + abs(y-c.y) // Part two
			}

			inf[mc] = inf[mc] || (x == 0 || y == 0 || x == w || y == h)
			m[mc]++

			// Part two
			if total < maxSum {
				regions++
			}
		}
	}

	// Find max area
	max := 0
	for k, v := range m {
		if v > max && !inf[k] {
			max = v
		}
	}

	return max, regions
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
