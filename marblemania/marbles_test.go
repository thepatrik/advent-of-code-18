package marblemania

import (
	"testing"
)

func TestWinningScorePart1(t *testing.T) {
	players := 411
	totalPlays := 72059
	score := WinningScore(players, totalPlays)
	t.Logf("Score was: %d", score)
	if score != 429943 {
		t.Fail()
	}
}

func TestWinningScorePart2(t *testing.T) {
	players := 411
	totalPlays := 72059 * 100
	score := WinningScore(players, totalPlays)
	t.Logf("Score was: %d", score)
	if score != 3615691746 {
		t.Fail()
	}
}
