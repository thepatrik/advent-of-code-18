package alchemicalreduction

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestCalculateUnits(t *testing.T) {
	b, err := ioutil.ReadFile("data")
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	str = strings.TrimSuffix(str, "\n")
	units := CalculateUnits(str)
	t.Logf("No of units: %d", units)
	if units != 9686 {
		t.Fail()
	}
}
