package marblemania

import (
	"testing"
)

func TestWinningScoreSmall(t *testing.T) {
	players := 10
	totalPlays := 1618
	score := WinningScore(players, totalPlays)
	t.Logf("Score was: %d", score)
	if score != 8317 {
		t.Fail()
	}
}

func TestWinningScore(t *testing.T) {
	players := 411
	totalPlays := 72059
	score := WinningScore(players, totalPlays)
	t.Logf("Score was: %d", score)
	if score != 429943 {
		t.Fail()
	}
}
