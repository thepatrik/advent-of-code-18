package subterraneansustainability

import (
	"sync"
)

// CalcPotSum calculates the sum of all pots that contain a plant.
func CalcPotSum(pots *Pots, patterns []Pattern, generations int) int {
	sumFunc := calcSum

	if generations > 1000000 {
		// Use forecast calculator func when the number of generations to
		// calculate exceeds 1 million
		sumFunc = calcForecast
	}

	return sumFunc(pots, patterns, generations)
}

func calcSum(pots *Pots, patterns []Pattern, generations int) int {
	return growGeneration(pots, patterns, 0, generations).Sum()
}

func calcForecast(pots *Pots, patterns []Pattern, generations int) int {
	// Find when the sum delta evens out
	sum, iterations, delta := func() (int, int, int) {
		nxtgen, gen, eq, delta := pots, 0, 0, 0

		// When the delta has been equal 10 times in a row, we
		// have probably found a pattern in the calculations.
		for eq < 10 {
			sum := nxtgen.Sum()
			// Step one generation a time to compare the sum delta
			// between each one.
			nxtgen = growGeneration(nxtgen, patterns, gen, gen+1)
			nxtdelta := nxtgen.Sum() - sum
			if nxtdelta == delta {
				eq++
			} else {
				eq = 0
			}
			delta = nxtdelta
			gen++
		}
		// return current sum, generation, and current delta value
		return nxtgen.Sum(), gen, delta
	}()
	// Now we can calculate (forecast) the complete sum
	return sum + (generations-iterations)*delta
}

func growGeneration(pots *Pots, patterns []Pattern, generation int, maxgen int) *Pots {
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
		return nxtpots
	}
	// Iterate until we hit desired generation
	return growGeneration(nxtpots, patterns, generation, maxgen)
}
