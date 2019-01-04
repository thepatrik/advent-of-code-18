package marblemania

import (
	"container/ring"
)

// WinningScore calculates the winning score
func WinningScore(players, highestVal int) int {
	circle := &ring.Ring{Value: 0}
	scores := make([]int, players)

	for nextMarble := 1; nextMarble <= highestVal; nextMarble++ {
		if nextMarble%23 == 0 {
			circle = circle.Move(-8)
			removed := circle.Unlink(1)
			scores[nextMarble%players] += nextMarble + removed.Value.(int)
			circle = circle.Move(1)
		} else {
			circle = circle.Move(1)

			s := &ring.Ring{Value: nextMarble}

			circle.Link(s)
			circle = circle.Move(1)
		}
	}
	highestScore := 0
	for _, score := range scores {
		if score > highestScore {
			highestScore = score
		}
	}
	return highestScore
}
