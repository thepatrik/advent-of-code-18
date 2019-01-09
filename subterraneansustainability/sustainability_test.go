package subterraneansustainability

import (
	"advent-of-code-18/testutils"
	"fmt"
	"testing"
)

func TestPotSumSmall(t *testing.T) {
	s := testutils.ReadFile("./data_s_initial")[0]
	pots := pots(s)
	strslice := testutils.ReadFile("./data_s")
	patterns := patterns(strslice)
	sum, _ := PotSum(pots, patterns, 20)
	t.Logf("Sum of pots was %d", sum)
	if sum != 325 {
		t.Fail()
	}
}

func BenchmarkPotSum(b *testing.B) {
	s := testutils.ReadFile("./data_initial")[0]
	pots := pots(s)
	strslice := testutils.ReadFile("./data")
	patterns := patterns(strslice)
	sum, _ := PotSum(pots, patterns, 100000)
	b.Logf("Sum of pots was %d", sum)
}

func TestPotSum(t *testing.T) {
	s := testutils.ReadFile("./data_initial")[0]
	pots := pots(s)
	strslice := testutils.ReadFile("./data")
	patterns := patterns(strslice)
	sum, _ := PotSum(pots, patterns, 20)
	t.Logf("Sum of pots was %d", sum)
	if sum != 1672 {
		t.Fail()
	}
}

func TestPotSum50Billion(t *testing.T) {
	s := testutils.ReadFile("./data_initial")[0]
	pots := pots(s)
	strslice := testutils.ReadFile("./data")
	patterns := patterns(strslice)
	sum, delta := PotSum(pots, patterns, 1000)
	// After x generations the delta between each generation stabilizes
	// (on my data it is 33). Take the sum of 1000 generations and add
	// the sum of deltas for the rest (49999999000).
	fiftybillion := sum + (50000000000-1000)*delta
	t.Logf("After 50 billion generations the sum is %d", fiftybillion)
	if fiftybillion != 1650000000055 {
		t.Fail()
	}
}

func pots(s string) *Pots {
	pots := NewPots()
	for i, r := range s {
		if r == '#' {
			pots.Grow(i)
		}
	}
	return pots
}

func patterns(strslice []string) []Pattern {
	patterns := make([]Pattern, 0)
	for _, s := range strslice {
		patternstr, expected := "", ""
		_, _ = fmt.Sscanf(s, "%s => %s",
			&patternstr, &expected)
		pattern := Pattern{
			L2: patternstr[0] == '#',
			L1: patternstr[1] == '#',
			C:  patternstr[2] == '#',
			R1: patternstr[3] == '#',
			R2: patternstr[4] == '#',
			N:  expected[0] == '#'}
		patterns = append(patterns, pattern)
	}
	return patterns
}
