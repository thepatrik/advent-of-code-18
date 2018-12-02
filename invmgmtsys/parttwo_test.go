package invmgmtsys

import (
	"advent-of-code-18/testutils"
	"testing"
)

func TestFindCmnLetters(t *testing.T) {
	strslice := testutils.ReadFile("./data")
	res := FindCmnLetters(strslice)
	t.Logf("Common letters was: %s", res)
}
