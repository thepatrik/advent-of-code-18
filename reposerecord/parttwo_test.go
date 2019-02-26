package reposerecord

import (
	"github.com/thepatrik/advent-of-code-18/testutils"
	"testing"
)

func TestCalcGuardFrequency(t *testing.T) {
	strslice := testutils.ReadFile("./data")
	freq := CalcGuardFrequency(strslice)
	t.Logf("Guard frequence was: %d", freq)
	if freq != 132484 {
		t.Fail()
	}
}
