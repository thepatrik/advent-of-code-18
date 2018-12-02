package cronalcalibration

import (
	"advent-of-code-18/testutils"
	"testing"
)

func TestSumFreq(t *testing.T) {
	strslice := testutils.ReadFile("./data")
	res := SumFreq(strslice, 0)
	t.Logf("Sum of frequencies are: %d", res)
}
