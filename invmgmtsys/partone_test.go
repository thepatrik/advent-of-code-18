package invmgmtsys

import (
	"advent-of-code-18/testutils"
	"testing"
)

func TestChkSum(t *testing.T) {
	strslice := testutils.ReadFile("./data")
	res := ChkSum(strslice)
	t.Logf("Check sum was: %d", res)
}
