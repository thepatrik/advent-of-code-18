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
	sum := CalcPotSum(pots, patterns, 20)
	t.Logf("Sum of pots was %d", sum)
	if sum != 325 {
		t.Fail()
	}
}

func TestPotSum(t *testing.T) {
	s := testutils.ReadFile("./data_initial")[0]
	pots := pots(s)
	strslice := testutils.ReadFile("./data")
	patterns := patterns(strslice)
	sum := CalcPotSum(pots, patterns, 20)
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
	iterations := 50000000000
	sum := CalcPotSum(pots, patterns, iterations)
	t.Logf("After 50 billion generations the sum is %d", sum)
	if sum != 1650000000055 {
		t.Fail()
	}
}

func BenchmarkPotSum(b *testing.B) {
	s := testutils.ReadFile("./data_initial")[0]
	pots := pots(s)
	strslice := testutils.ReadFile("./data")
	patterns := patterns(strslice)
	sum := calcSum(pots, patterns, 100000)
	b.Logf("Sum of pots was %d", sum)
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
