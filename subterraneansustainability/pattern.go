package subterraneansustainability

// Pattern type
type Pattern struct {
	L2 bool
	L1 bool
	C  bool
	R1 bool
	R2 bool
	N  bool
}

// String generates a string from the current pattern
func (p *Pattern) String() string {
	s := ""
	s += boolToString(p.L2)
	s += boolToString(p.L1)
	s += boolToString(p.C)
	s += boolToString(p.R1)
	s += boolToString(p.R2)
	s += " " + boolToString(p.N)
	return s
}
