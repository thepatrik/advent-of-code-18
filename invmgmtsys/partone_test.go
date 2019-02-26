package invmgmtsys

import (
	"github.com/thepatrik/advent-of-code-18/testutils"
	"testing"
)

func TestChkSum(t *testing.T) {
	strslice := testutils.ReadFile("./data")
	chkSum := ChkSum(strslice)
	t.Logf("Check sum was: %d", chkSum)
	if chkSum != 6000 {
		t.Fail()
	}
}
