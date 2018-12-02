package cronalcalibration

import (
	"advent-of-code-18/testutils"
	"testing"
)

func TestFindDupFreq(t *testing.T) {
	strslice := testutils.ReadFile("./data")
	res := FindDupFreq(strslice)
	t.Logf("First duplicate frequency is: %d", res)
}
