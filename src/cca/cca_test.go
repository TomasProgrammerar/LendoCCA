package cca

import (
	"testing"
)

func TestLookupColor(t *testing.T) {
	if testColor, err := lookupColor(-1, rgbPallet); err == nil {
		t.Errorf("lookup of index -1 failed.\nExpected: panic\n Got: %v", testColor)
	}

	if testColor, err := lookupColor(16, rgbPallet); err == nil {
		t.Errorf("lookup of index 16 failed.\nExpected: panic\n Got: %v", testColor)
	}

	for colorIndex := 0; colorIndex < 16; colorIndex++ {
		if testColor, err := lookupColor(colorIndex, rgbPallet); err != nil ||
			testColor.R != rgbPallet[colorIndex].R ||
			testColor.G != rgbPallet[colorIndex].G ||
			testColor.B != rgbPallet[colorIndex].B {
			t.Errorf("lookup of index %d failed.\nExpected: %v\n Got: %s", colorIndex, rgbPallet[colorIndex], err)
		}
	}
}
