package chronalclassification

//Instruction type
type Instruction struct {
	Code    int
	InputA  int
	InputB  int
	OutputC int
}

// Sample type
type Sample struct {
	RegBefore []int
	Instr     Instruction
	RegAfter  []int
}

type operation func() []int

func (s *Sample) addr() []int {
	res := make([]int, len(s.RegBefore))
	copy(res, s.RegBefore)
	res[s.Instr.OutputC] = res[s.Instr.InputA] + res[s.Instr.InputB]
	return res
}

func (s *Sample) addi() []int {
	res := make([]int, len(s.RegBefore))
	copy(res, s.RegBefore)
	res[s.Instr.OutputC] = res[s.Instr.InputA] + s.Instr.InputB
	return res
}

func (s *Sample) mulr() []int {
	res := make([]int, len(s.RegBefore))
	copy(res, s.RegBefore)
	res[s.Instr.OutputC] = res[s.Instr.InputA] * res[s.Instr.InputB]
	return res
}

func (s *Sample) muli() []int {
	res := make([]int, len(s.RegBefore))
	copy(res, s.RegBefore)
	res[s.Instr.OutputC] = res[s.Instr.InputA] * s.Instr.InputB
	return res
}

func (s *Sample) banr() []int {
	res := make([]int, len(s.RegBefore))
	copy(res, s.RegBefore)
	res[s.Instr.OutputC] = res[s.Instr.InputA] & res[s.Instr.InputB]
	return res
}

func (s *Sample) bani() []int {
	res := make([]int, len(s.RegBefore))
	copy(res, s.RegBefore)
	res[s.Instr.OutputC] = res[s.Instr.InputA] & s.Instr.InputB
	return res
}

func (s *Sample) borr() []int {
	res := make([]int, len(s.RegBefore))
	copy(res, s.RegBefore)
	res[s.Instr.OutputC] = res[s.Instr.InputA] | res[s.Instr.InputB]
	return res
}

func (s *Sample) bori() []int {
	res := make([]int, len(s.RegBefore))
	copy(res, s.RegBefore)
	res[s.Instr.OutputC] = res[s.Instr.InputA] | s.Instr.InputB
	return res
}

func (s *Sample) setr() []int {
	res := make([]int, len(s.RegBefore))
	copy(res, s.RegBefore)
	res[s.Instr.OutputC] = res[s.Instr.InputA]
	return res
}

func (s *Sample) seti() []int {
	res := make([]int, len(s.RegBefore))
	copy(res, s.RegBefore)
	res[s.Instr.OutputC] = s.Instr.InputA
	return res
}

func (s *Sample) gtir() []int {
	res := make([]int, len(s.RegBefore))
	copy(res, s.RegBefore)
	if s.Instr.InputA > res[s.Instr.InputB] {
		res[s.Instr.OutputC] = 1
	} else {
		res[s.Instr.OutputC] = 0
	}
	return res
}

func (s *Sample) gtri() []int {
	res := make([]int, len(s.RegBefore))
	copy(res, s.RegBefore)
	if res[s.Instr.InputA] > s.Instr.InputB {
		res[s.Instr.OutputC] = 1
	} else {
		res[s.Instr.OutputC] = 0
	}
	return res
}

func (s *Sample) gtrr() []int {
	res := make([]int, len(s.RegBefore))
	copy(res, s.RegBefore)
	if res[s.Instr.InputA] > res[s.Instr.InputB] {
		res[s.Instr.OutputC] = 1
	} else {
		res[s.Instr.OutputC] = 0
	}
	return res
}

func (s *Sample) eqir() []int {
	res := make([]int, len(s.RegBefore))
	copy(res, s.RegBefore)
	if s.Instr.InputA == res[s.Instr.InputB] {
		res[s.Instr.OutputC] = 1
	} else {
		res[s.Instr.OutputC] = 0
	}
	return res
}

func (s *Sample) eqri() []int {
	res := make([]int, len(s.RegBefore))
	copy(res, s.RegBefore)
	if res[s.Instr.InputA] == s.Instr.InputB {
		res[s.Instr.OutputC] = 1
	} else {
		res[s.Instr.OutputC] = 0
	}
	return res
}

func (s *Sample) eqrr() []int {
	res := make([]int, len(s.RegBefore))
	copy(res, s.RegBefore)
	if res[s.Instr.InputA] == res[s.Instr.InputB] {
		res[s.Instr.OutputC] = 1
	} else {
		res[s.Instr.OutputC] = 0
	}
	return res
}

func (s *Sample) eq(slice []int) bool {
	if len(s.RegAfter) != len(slice) {
		return false
	}
	for i, v := range s.RegAfter {
		if v != slice[i] {
			return false
		}
	}
	return true
}
