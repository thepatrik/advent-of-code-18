package subterraneansustainability

import (
	"sort"
)

// Pots type
type Pots struct {
	m map[int]bool
}

// NewPots constructs pots
func NewPots() *Pots {
	return &Pots{m: make(map[int]bool)}
}

// Grow plant in a pot
func (p *Pots) Grow(pot int) {
	p.m[pot] = true
}

// Keys returns all keys (sorted asc)
func (p *Pots) Keys() *[]int {
	keys := make([]int, 0, len(p.m))
	for k := range p.m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return &keys
}

// Offset returns the indice of the first growing plant
func (p *Pots) Offset() int {
	return offset(*p.Keys())
}

// Len returns the length and offset
func (p *Pots) Len() (int, int) {
	keys := *p.Keys()
	return length(keys), offset(keys)
}

// Sum calculates the sum of all pots with a plant
func (p *Pots) Sum() int {
	sum := 0
	for k := range p.m {
		sum += k
	}
	return sum
}

// Slice turns pots to slice of bool
func (p *Pots) Slice() []bool {
	nxt := make([]bool, 0)
	if len(p.m) == 0 {
		return nxt
	}
	keys := *p.Keys()
	old := -100
	for _, ix := range keys {
		if len(nxt) > 0 {
			for i := old + 1; i < ix; i++ {
				nxt = append(nxt, false)
			}
		}
		nxt = append(nxt, true)
		old = ix
	}

	return nxt
}

// Compare if a pattern checks out
func (p *Pots) Compare(pot int, pattern *Pattern) bool {
	if p.m[pot-2] != pattern.L2 {
		return false
	}
	if p.m[pot-1] != pattern.L1 {
		return false
	}
	if p.m[pot] != pattern.C {
		return false
	}
	if p.m[pot+1] != pattern.R1 {
		return false
	}
	if p.m[pot+2] != pattern.R2 {
		return false
	}
	return true
}

// String returns a string representation
func (p *Pots) String() string {
	s := ""
	for _, v := range p.Slice() {
		s += boolToString(v)
	}
	return s
}

func length(keys []int) int {
	switch len(keys) {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return keys[len(keys)-1] - keys[0]
	}
}

func offset(keys []int) int {
	if len(keys) == 0 {
		return 0
	}
	return keys[0]
}
