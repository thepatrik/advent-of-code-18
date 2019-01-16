package dartminemadness

import (
	"errors"
	"sort"
)

// Carts type
type Carts struct {
	slice []*Cart
}

// NewCarts constructs new carts type
func NewCarts() *Carts {
	slice := make([]*Cart, 0)
	return &Carts{slice: slice}
}

// Get cart at index
func (carts *Carts) Get(i int) *Cart {
	return carts.slice[i]
}

// Add adds a cart
func (carts *Carts) Add(cart *Cart) {
	carts.slice = append(carts.slice, cart)
}

// RemoveAll removes carts
func (carts *Carts) RemoveAll(c []*Cart) {
	for _, cart := range c {
		carts.Remove(cart)
	}
}

// Remove removes a cart
func (carts *Carts) Remove(cart *Cart) {
	index := func() int {
		for i, c := range carts.slice {
			if cart.id == c.id {
				return i
			}
		}
		return -1
	}()
	carts.slice = append(carts.slice[:index], carts.slice[index+1:]...)
}

// Len returns the length of carts
func (carts *Carts) Len() int {
	return len(carts.slice)
}

// Sort sorts the slice of carts by positions asc
func (carts *Carts) Sort() []*Cart {
	sort.SliceStable(carts.slice, func(i, j int) bool {
		if carts.slice[i].pos.x == carts.slice[j].pos.x {
			return carts.slice[i].pos.y < carts.slice[j].pos.y
		}
		return carts.slice[i].pos.x < carts.slice[j].pos.x
	})
	return carts.slice
}

// FindCart looks for a cart at position
func (carts *Carts) FindCart(x, y int) (*Cart, error) {
	var res *Cart
	for _, cart := range carts.slice {
		if cart.pos.x == x && cart.pos.y == y {
			res = cart
			return res, nil
		}
	}
	return res, errors.New("Could not find cart")
}

// HitTest performs a hit test
func (carts *Carts) HitTest(cart *Cart) (bool, []*Cart) {
	wrecks := make([]*Cart, 0)
	for _, c := range carts.slice {
		if c.id == cart.id {
			continue
		}
		if c.pos.x == cart.pos.x && c.pos.y == cart.pos.y {
			wrecks = append(wrecks, c)
		}
	}
	res := false
	if len(wrecks) > 0 {
		res = true
		wrecks = append(wrecks, cart)
	}
	return res, wrecks
}
