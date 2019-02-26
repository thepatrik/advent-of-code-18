package sumofparts

import (
	"github.com/thepatrik/advent-of-code-18/testutils"
	"testing"
)

func TestStepOrder(t *testing.T) {
	strslice := testutils.ReadFile("./data")
	steps := steps(strslice)
	order := Order(steps)
	t.Logf("Step order is: %s", order)
	if order != "HEGMPOAWBFCDITVXYZRKUQNSLJ" {
		t.Fail()
	}
}
