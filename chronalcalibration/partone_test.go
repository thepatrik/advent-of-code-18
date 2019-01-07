package chronalcalibration

import (
	"advent-of-code-18/testutils"
	"testing"
)

func TestSumFreq(t *testing.T) {
	strslice := testutils.ReadFile("./data")
	sum := SumFreq(strslice, 0)
	t.Logf("Sum of frequencies are: %d", sum)
	if sum != 531 {
		t.Fail()
	}
}
