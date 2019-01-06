package sumofparts

import (
	"advent-of-code-18/testutils"
	"testing"
)

func TestCompletionTime(t *testing.T) {
	strslice := testutils.ReadFile("./data")
	steps := steps(strslice)
	concurrency := 5
	time := CompletionTime(steps, concurrency)
	t.Logf("Completion time was: %d", time)
	if time != 1226 {
		t.Fail()
	}
}
