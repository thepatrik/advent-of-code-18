package dartminemadness

import (
	"strconv"
)

// Cart type
type Cart struct {
	direction    Direction
	id           string
	pos          *Position
	intersection turn
}

// NewCart constructs a new cart
func NewCart(direction Direction, id string, x, y int) *Cart {
	pos := &Position{x: x, y: y}
	return &Cart{direction: direction, id: id, pos: pos, intersection: initial}
}

// Intersection handles an intersection turn
func (cart *Cart) Intersection() Direction {
	switch cart.intersection {
	case initial:
		cart.intersection = left
	case skip:
		cart.intersection = right
	case right:
		cart.intersection = left
	case left:
		cart.intersection = skip
	}
	return cart.direction.Turn(cart.intersection)
}

// Move moves the cart
func (cart *Cart) Move(direction Direction) {
	switch direction {
	case north:
		cart.pos.y--
	case south:
		cart.pos.y++
	case east:
		cart.pos.x++
	case west:
		cart.pos.x--
	}
	cart.direction = direction
}

func (cart Cart) String() string {
	return cart.id + " " + strconv.Itoa(cart.pos.x) + "," + strconv.Itoa(cart.pos.y) + " " + cart.direction.String()
}
