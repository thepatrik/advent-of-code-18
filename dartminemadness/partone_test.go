package dartminemadness

import (
	"github.com/thepatrik/advent-of-code-18/testutils"
	"testing"
)

func TestCrashDetectSmall(t *testing.T) {
	strslice := testutils.ReadFile("./data_s")
	raceTrack := raceTrack(strslice)
	pos := DetectFirstCrash(raceTrack)
	t.Logf("Crashed at: %d,%d", pos.x, pos.y)
	if pos.x != 7 || pos.y != 3 {
		t.Fail()
	}
}

func TestCrashDetect(t *testing.T) {
	strslice := testutils.ReadFile("./data")
	raceTrack := raceTrack(strslice)
	pos := DetectFirstCrash(raceTrack)
	t.Logf("Crashed at: %d,%d", pos.x, pos.y)
	if pos.x != 26 || pos.y != 92 {
		t.Fail()
	}
}
