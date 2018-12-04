package nomatterhowyousliceit

import (
	"advent-of-code-18/testutils"
	"log"
	"testing"
)

func TestGetNonOverlappingFabricID(t *testing.T) {
	strslice := testutils.ReadFile("./data")
	res := GetNonOverlappingFabricID(strslice)
	log.Println(res)
	t.Logf("The ID of the non-overlapping fabric was: %d", res)
	if res != 243 {
		t.Fail()
	}
}

func TestGetNonOverlappingFabricIDSmall(t *testing.T) {
	strslice := testutils.ReadFile("./data_small")
	res := GetNonOverlappingFabricID(strslice)
	t.Logf("The ID of the non-overlapping fabric was: %d", res)
	if res != 3 {
		t.Fail()
	}
}
