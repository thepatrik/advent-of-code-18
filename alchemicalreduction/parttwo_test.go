package alchemicalreduction

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestFindShortestPolymerLen(t *testing.T) {
	b, err := ioutil.ReadFile("data")
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	str = strings.TrimSuffix(str, "\n")
	len := FindShortestPolymerLen(str)
	t.Logf("Length of polymer: %d", len)
	if len != 5524 {
		t.Fail()
	}
}
