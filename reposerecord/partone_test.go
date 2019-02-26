package reposerecord

import (
	"github.com/thepatrik/advent-of-code-18/testutils"
	"testing"
)

func TestCalcGuardScore(t *testing.T) {
	strslice := testutils.ReadFile("./data")
	score := CalcGuardScore(strslice)
	t.Logf("Guard score was: %d", score)
	if score != 74743 {
		t.Fail()
	}
}
