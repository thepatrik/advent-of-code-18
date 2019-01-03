package starsalign

import (
	"math"
	"strings"
)

// Sky type
type Sky struct {
	points    []Point
	iteration int
}

// Point type
type Point struct {
	posx int
	posy int
	velx int
	vely int
}

// BoundingBox type
type BoundingBox struct {
	top    int
	left   int
	right  int
	bottom int
}

// newSky constructs a Sky
func newSky(points []Point) Sky {
	return Sky{points: points}
}

func (sky *Sky) fwd() {
	for i, point := range sky.points {
		sky.points[i].posx += point.velx
		sky.points[i].posy += point.vely
	}
	sky.iteration++
}

func (sky *Sky) rwd() {
	for i, point := range sky.points {
		sky.points[i].posx -= point.velx
		sky.points[i].posy -= point.vely
	}
	sky.iteration--
}

func (sky *Sky) dimensions() (int, int) {
	bbox := sky.bounds()
	w := bbox.right - bbox.left + 1
	h := bbox.bottom - bbox.top + 1
	return w, h
}

func (sky *Sky) size() int64 {
	w, h := sky.dimensions()
	return int64(w * h)
}

func (sky *Sky) bounds() BoundingBox {
	box := BoundingBox{top: math.MaxInt32, left: math.MaxInt32, right: math.MinInt32, bottom: math.MinInt32}
	for _, point := range sky.points {
		if point.posx > box.right {
			box.right = point.posx
		} else if point.posx < box.left {
			box.left = point.posx
		}

		if point.posy < box.top {
			box.top = point.posy
		} else if point.posy > box.bottom {
			box.bottom = point.posy
		}
	}
	return box
}

func (sky *Sky) plot() string {
	w, h := sky.dimensions()
	matrix := make([][]bool, h)
	for i := 0; i < h; i++ {
		matrix[i] = make([]bool, w)
	}

	bbox := sky.bounds()
	for _, point := range sky.points {
		x := point.posx - bbox.left
		y := point.posy - bbox.top
		matrix[y][x] = true
	}

	var builder strings.Builder
	builder.WriteRune('\n')
	for h := 0; h < len(matrix); h++ {
		for w := 0; w < len(matrix[h]); w++ {
			if matrix[h][w] {
				builder.WriteRune('#')
			} else {
				builder.WriteRune('.')
			}
		}
		builder.WriteRune('\n')
	}
	return builder.String()
}

func (sky *Sky) traverse() {
	size := sky.size()
	sky.fwd()

	// When the size starts to increase, we have hit bottom.
	if sky.size() > size {
		// Rewind and unfold
		sky.rwd()
		return
	}
	sky.traverse()
}

// FindMessage looks for a message in the sky
func (sky *Sky) FindMessage() string {
	sky.traverse()
	return sky.plot()
}
