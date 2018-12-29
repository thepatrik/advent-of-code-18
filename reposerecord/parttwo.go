package reposerecord

// CalcGuardFrequency Finds the guard ID who is most
// frequently asleep on the same minute multiplied
// with the same minute.
func CalcGuardFrequency(strslice []string) int {
	//map[guardID][minute(0-1400)]count
	m := getMap(strslice)

	// Find the guard with the highest time frequence asleep
	_, guard, minute := func() (int, int, int) {
		guard, hiscore, minute := -1, -1, -1

		for id, time := range m {
			for min, count := range time {
				if count > hiscore {
					guard = id
					hiscore = count
					minute = min
				}
			}
		}
		return hiscore, guard, minute
	}()

	return guard * minute
}
