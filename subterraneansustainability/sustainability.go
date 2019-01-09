package subterraneansustainability

import (
	"sync"
)

// PotSum calculates the sum of all pots that contain a plant along with the
// delta for the last generation (used in part two).
func PotSum(pots *Pots, patterns []Pattern, generations int) (int, int) {
	return growGeneration(pots, patterns, 0, generations)
}

func growGeneration(pots *Pots, patterns []Pattern, generation int, maxgen int) (int, int) {
	noOfPots, offset := pots.Len()
	noOfPatterns := len(patterns)
	wg, mutex := sync.WaitGroup{}, sync.Mutex{}
	wg.Add(noOfPatterns)
	nxtpots := NewPots()

	for p := 0; p < noOfPatterns; p++ {
		go func(pattern *Pattern) {
			defer wg.Done()
			// Patterns can match two pots to the left of the first growing plant
			// and two to the right of the last plant
			for i := -2; i < noOfPots+2; i++ {
				pot := i + offset
				// No need to check patterns when N is false
				if pattern.N && pots.Compare(pot, pattern) {
					func() {
						mutex.Lock()
						defer mutex.Unlock()
						// Grow plant for next generation
						nxtpots.Grow(pot)
					}()
				}
			}
		}(&patterns[p])
	}
	wg.Wait() // Wait for patterns to finish

	if generation++; generation == maxgen {
		// Returns the calculated sum along with the delta for the last gen
		return nxtpots.Sum(), nxtpots.Sum() - pots.Sum()
	}
	// Iterate until we hit desired generation
	return growGeneration(nxtpots, patterns, generation, maxgen)
}
