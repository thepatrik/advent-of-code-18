package reposerecord

// CalcGuardScore Finds the guard ID who sleeps the
// most and multiplies the ID together with the
// minute he sleeps the most.
func CalcGuardScore(strslice []string) int {
	//map[guardID][minute(0-1400)]count
	m := getMap(strslice)

	// Find the sleepiest guard
	_, sleepiestGuard := func() (int, int) {
		hiscore, guard := -1, -1

		for id, min := range m {
			score := 0
			for _, count := range min {
				score += count
			}
			if score > hiscore {
				hiscore = score
				guard = id
			}
		}
		return hiscore, guard
	}()

	// Find the minute where the guard is most likely to be asleep
	_, minute := func(guard int) (int, int) {
		score, minute := -1, -1

		for time, count := range m[guard] {
			if count > score {
				score = count
				minute = time
			}
		}
		return score, minute
	}(sleepiestGuard)

	return sleepiestGuard * minute
}
