package chronalclassification

import (
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/thepatrik/advent-of-code-18/testutils"
)

func TestCalculateSamples(t *testing.T) {
	strslice := testutils.ReadFile("./data")
	samples := getSamples(strslice)
	res := CalculateSamples(samples)
	t.Logf("No of samples: %d", res)
	if res != 590 {
		t.Fail()
	}
}

// This is stupid but gets the work done...
func getSamples(strslice []string) []*Sample {
	samples := make([]*Sample, 0)
	currentSample := &Sample{}
	for _, s := range strslice {
		split := regexp.MustCompile("[^0-9]+").Split(strings.TrimRight(s, "]"), 5)
		if len(split) > 1 {
			i := 0
			if split[0] == "" {
				i = 1
			}
			code, _ := strconv.Atoi(split[i])
			i++
			a, _ := strconv.Atoi(split[i])
			i++
			b, _ := strconv.Atoi(split[i])
			i++
			c, _ := strconv.Atoi(split[i])
			curr := []int{code, a, b, c}
			if strings.HasPrefix(s, "Before:") {
				currentSample.RegBefore = curr
			} else if strings.HasPrefix(s, "After") {
				currentSample.RegAfter = curr
				samples = append(samples, currentSample)
				currentSample = &Sample{}
			} else {
				currentSample.Instr = Instruction{Code: code, InputA: a, InputB: b, OutputC: c}
			}
		}
	}
	return samples
}
