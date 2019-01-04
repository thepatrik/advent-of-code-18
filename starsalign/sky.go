package starsalign

import (
	"strings"
)

// Sky type
type Sky struct {
	stars     []Star
	iteration int
}

// Star type
type Star struct {
	position Position
	velocity Velocity
}

// Position (x, y). X represents left (negative) or
// right (positive), while Y represents up (negative)
// or down (positive).
type Position struct {
	x int
	y int
}

// Velocity (x, y) per second.
type Velocity struct {
	x int
	y int
}

// BoundingBox type
type BoundingBox struct {
	top    int
	left   int
	right  int
	bottom int
}

// NewSky constructs a brand new Sky.
func NewSky(stars []Star) Sky {
	return Sky{stars: stars}
}

// Tick iterates the clock one second, and moves each
// star's position according to the velocity.
func (sky *Sky) Tick() {
	for i := range sky.stars {
		sky.stars[i].position.x += sky.stars[i].velocity.x
		sky.stars[i].position.y += sky.stars[i].velocity.y
	}
	sky.iteration++
}

// Rewind iterates the clock one second backwards, and
// moves each star's position according to the velocity.
func (sky *Sky) Rewind() {
	for i := range sky.stars {
		sky.stars[i].position.x -= sky.stars[i].velocity.x
		sky.stars[i].position.y -= sky.stars[i].velocity.y
	}
	sky.iteration--
}

// Bounds returns a bounding box around the current
// alignment of the stars.
func (sky Sky) Bounds() BoundingBox {
	bbox := BoundingBox{top: 1<<31 - 1, left: 1<<31 - 1, right: -1 << 31, bottom: -1 << 31}

	for _, star := range sky.stars {
		if star.position.x > bbox.right {
			bbox.right = star.position.x
		} else if star.position.x < bbox.left {
			bbox.left = star.position.x
		}
		if star.position.y < bbox.top {
			bbox.top = star.position.y
		} else if star.position.y > bbox.bottom {
			bbox.bottom = star.position.y
		}
	}
	return bbox
}

// Dimensions returns the dimensions (width*height)
// based on the bounding box.
func (sky Sky) Dimensions() (int, int) {
	bbox := sky.Bounds()
	w := bbox.right - bbox.left + 1
	h := bbox.bottom - bbox.top + 1
	return w, h
}

// Area returns the size from the dimensions.
func (sky Sky) Area() int64 {
	w, h := sky.Dimensions()
	return int64(w * h)
}

// Matrix turns the sky into a plottable two-dimensional
// slice.
func (sky Sky) Matrix() [][]bool {
	w, h := sky.Dimensions()
	matrix := make([][]bool, h)
	for i := 0; i < h; i++ {
		matrix[i] = make([]bool, w)
	}

	bbox := sky.Bounds()
	for _, star := range sky.stars {
		x := star.position.x - bbox.left
		y := star.position.y - bbox.top
		matrix[y][x] = true
	}
	return matrix
}

// String returns a string based on a matrix of the
// star positions.
func (sky Sky) String() string {
	var builder strings.Builder
	builder.WriteRune('\n')
	matrix := sky.Matrix()
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

// Search looks for a message in the sky. Iterates until the
// message is found, or we run out of memory :-)
func (sky *Sky) Search() {
	area := sky.Area()
	sky.Tick()

	if sky.Area() > area {
		// When the area starts to increase, we have hit bottom.
		sky.Rewind()
		return
	}
	sky.Search()
}
