package chronalclassification

import (
	"advent-of-code-18/testutils"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func TestCalculateOperations(t *testing.T) {
	strslice := testutils.ReadFile("./data")
	samples := getSamples(strslice)
	instr := getInstructions(testutils.ReadFile("./test_program"))
	val := CalculateRegistry(samples, instr)
	t.Logf("Value is: %d", val)
	if val != 475 {
		t.Fail()
	}
}

func getInstructions(strSlice []string) []Instruction {
	instr := make([]Instruction, len(strSlice))
	for i, s := range strSlice {
		split := regexp.MustCompile("[^0-9]+").Split(strings.TrimRight(s, "]"), 5)
		code, _ := strconv.Atoi(split[0])
		a, _ := strconv.Atoi(split[1])
		b, _ := strconv.Atoi(split[2])
		c, _ := strconv.Atoi(split[3])
		instr[i] = Instruction{Code: code, InputA: a, InputB: b, OutputC: c}
	}
	return instr
}
