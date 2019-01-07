package chronalcoordinates

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"testing"
)

func TestLargestArea(t *testing.T) {
	coords, w, h := input()
	area, region := LargestArea(coords, w, h)
	t.Logf("Area was: %d", area)
	if area != 4166 {
		t.Fail()
	}
	t.Logf("Region was: %d", region)
	if region != 42250 {
		t.Fail()
	}
}

func input() ([]coord, int, int) {
	type Input struct {
		n int
		x int
		y int
	}

	file := flag.String("file", "./data", "file used for input")
	flag.Parse()
	fileHandle, _ := os.Open(*file)
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	input := []Input{}
	n := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			i := Input{}
			_, _ = fmt.Sscanf(line, "%d, %d", &i.x, &i.y)
			i.n = n
			n++
			input = append(input, i)
		}
	}

	var coords []coord
	w, h := 0, 0
	for _, i := range input {
		if i.y > h {
			h = i.y
		}
		if i.x > w {
			w = i.x
		}
		coords = append(coords, coord{x: i.x, y: i.y})
	}
	return coords, w, h
}
