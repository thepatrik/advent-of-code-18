package starsalign

import (
	"advent-of-code-18/testutils"
	"fmt"
	"log"
	"testing"
)

func TestFindMessage(t *testing.T) {
	strslice := testutils.ReadFile("./data")
	points := points(strslice)
	sky := newSky(points)
	message := sky.FindMessage()
	log.Print(message)
	t.Logf("Waiting time would have been %d secs.", sky.iteration)
	if sky.iteration != 10027 {
		t.Fail()
	}
}

func points(strslice []string) []Point {
	points := make([]Point, 0)
	for _, s := range strslice {
		var point Point
		_, _ = fmt.Sscanf(s, "position=<%d, %d> velocity=<%d, %d>",
			&point.posx, &point.posy, &point.velx, &point.vely)
		points = append(points, point)
	}
	return points
}
