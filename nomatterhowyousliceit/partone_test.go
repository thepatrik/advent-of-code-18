package nomatterhowyousliceit

import (
	"advent-of-code-18/testutils"
	"testing"
)

func TestGetOverlappingFabrics(t *testing.T) {
	strslice := testutils.ReadFile("./data")
	res := GetOverlappingFabrics(strslice)
	t.Logf("Number of overlapping fabrics was: %d", res)
	if res != 101196 {
		t.Fail()
	}
}

func TestGetOverlappingFabricsSmall(t *testing.T) {
	strslice := testutils.ReadFile("./data_small")
	res := GetOverlappingFabrics(strslice)
	t.Logf("Number of overlapping fabrics was: %d", res)
	if res != 4 {
		t.Fail()
	}
}
