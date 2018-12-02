package cronalcalibration

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestSumFreq(t *testing.T) {
	strslice := read()
	res := SumFreq(strslice, 0)
	t.Logf("Sum of frequencies are: %d", res)
}

func TestFindDupFreq(t *testing.T) {
	strslice := read()
	res := FindDupFreq(strslice)
	t.Logf("First duplicate frequency is: %d", res)
}

func read() []string {
	f, _ := os.Open("./data")
	defer f.Close()

	strslice := make([]string, 0)
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')

		if err == io.EOF {
			fmt.Print(line)
			break
		}
		line = strings.TrimSuffix(line, "\n")
		strslice = append(strslice, line)
	}
	return strslice
}
