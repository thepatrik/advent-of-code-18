package invmgmtsys

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestChkSum(t *testing.T) {
	strslice := read()
	res := ChkSum(strslice)
	t.Logf("Check sum was: %d", res)
}

func TestFindCmnLetters(t *testing.T) {
	strslice := read()
	res := FindCmnLetters(strslice)
	t.Logf("Common letters was: %s", res)
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
