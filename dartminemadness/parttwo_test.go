package dartminemadness

import (
	"github.com/thepatrik/advent-of-code-18/testutils"
	"testing"
)

func TestLastCartPosSmall(t *testing.T) {
	strslice := testutils.ReadFile("./data_s_2")
	raceTrack := raceTrack(strslice)
	pos := LastCartPos(raceTrack)
	t.Logf("Last position was at: %d,%d", pos.x, pos.y)
	if pos.x != 6 || pos.y != 4 {
		t.Fail()
	}
}

func TestLastCartPos(t *testing.T) {
	strslice := testutils.ReadFile("./data")
	raceTrack := raceTrack(strslice)
	pos := LastCartPos(raceTrack)
	t.Logf("Last position was at: %d,%d", pos.x, pos.y)
	if pos.x != 86 || pos.y != 18 {
		t.Fail()
	}
}
