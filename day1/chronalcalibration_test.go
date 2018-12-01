package cronalcalibration

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

func TestFreq(t *testing.T) {
	strslice := read()
	res := freq(strslice, 0)
	log.Printf("Got frequency: %d", res)
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
