package subterraneansustainability

import "sync"

// PotSum calculates the sum of all pots that contain a plant
func PotSum(pots *Pots, patterns []Pattern, generations int) int {
	return growGeneration(pots, patterns, 0, generations)
}

func growGeneration(pots *Pots, patterns []Pattern, generation int, maxgen int) int {
	if generation == maxgen {
		return pots.Sum()
	}

	noOfPots, offset := pots.Len()
	noOfPatterns := len(patterns)
	wg, mutex := sync.WaitGroup{}, sync.Mutex{}
	wg.Add(noOfPatterns)
	nxtpots := NewPots()

	for p := 0; p < noOfPatterns; p++ {
		go func(pattern *Pattern) {
			defer wg.Done()
			// Patterns can match pots two to the left of the first growing plant
			// and two to the right of the last plant
			for i := -2; i < noOfPots+2; i++ {
				pot := i + offset
				// No need to check patterns when N is false
				if pattern.N && pots.Compare(pot, pattern) {
					func() {
						mutex.Lock()
						defer mutex.Unlock()
						nxtpots.Grow(pot)
					}()
				}
			}
		}(&patterns[p])
	}
	wg.Wait() // Wait for patterns to finish before starting next gen
	return growGeneration(nxtpots, patterns, generation+1, maxgen)
}
