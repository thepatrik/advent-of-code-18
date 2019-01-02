package memorymaneuver

import (
	"advent-of-code-18/testutils"
	"testing"
)

func TestCalcRootNodeVal(t *testing.T) {
	nums := testutils.ReadIntFile("./data")
	val := CalcRootNodeVal(nums)
	t.Logf("Value of the root node is: %d", val)
	if val != 37067 {
		t.Fail()
	}
}
