package starsalign

import (
	"github.com/thepatrik/advent-of-code-18/testutils"
	"fmt"
	"testing"
)

func TestFindMessage(t *testing.T) {
	strslice := testutils.ReadFile("./data")
	stars := stars(strslice)
	sky := NewSky(stars)
	sky.Search()
	t.Log(sky)
	t.Logf("Waiting time would have been %d secs.", sky.iteration)
	if sky.iteration != 10027 {
		t.Fail()
	}
}

func stars(strslice []string) []Star {
	stars := make([]Star, 0)
	for _, s := range strslice {
		var star Star
		_, _ = fmt.Sscanf(s, "position=<%d, %d> velocity=<%d, %d>",
			&star.position.x, &star.position.y, &star.velocity.x, &star.velocity.y)
		stars = append(stars, star)
	}
	return stars
}
