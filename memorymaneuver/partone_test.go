package memorymaneuver

import (
	"advent-of-code-18/testutils"
	"testing"
)

func TestMetadataSum(t *testing.T) {
	nums := testutils.ReadIntFile("./data")
	sum := CalcMetadataSum(nums)
	t.Logf("Metadata sum was: %d", sum)
	if sum != 40984 {
		t.Fail()
	}
}
