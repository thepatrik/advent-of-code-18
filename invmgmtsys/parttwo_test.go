package invmgmtsys

import (
	"github.com/thepatrik/advent-of-code-18/testutils"
	"testing"
)

func TestFindCmnLetters(t *testing.T) {
	strslice := testutils.ReadFile("./data")
	letters := FindCmnLetters(strslice)
	t.Logf("Common letters was: %s", letters)
	if letters != "pbykrmjmizwhxlqnasfgtycdv" {
		t.Fail()
	}
}
