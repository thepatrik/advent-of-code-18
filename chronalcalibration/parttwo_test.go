package chronalcalibration

import (
	"advent-of-code-18/testutils"
	"testing"
)

func TestFindDupFreq(t *testing.T) {
	strslice := testutils.ReadFile("./data")
	duplicate := FindDupFreq(strslice)
	t.Logf("First duplicate frequency is: %d", duplicate)
	if duplicate != 76787 {
		t.Fail()
	}
}
