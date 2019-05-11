package cca

import (
	"testing"
)

//Helper function to check if colors match
func matchColor(c1, c2 color) bool {
	return (c1.R == c2.R && c1.G == c2.G && c1.B == c2.B)
}

func TestLookupColor(t *testing.T) {
	//Range
	if _, err := lookupColor(-1, rgbPallet); err == nil {
		t.Errorf("lookup of index -1 failed.\nExpected: Invalid color range\n Got: %s", err)
	}

	if _, err := lookupColor(16, rgbPallet); err == nil {
		t.Errorf("lookup of index 16 failed.\nExpected: Invalid color range\n Got: %s", err)
	}

	//Sanity check
	for colorIndex := 0; colorIndex < 16; colorIndex++ {
		if lookedUpColor, err := lookupColor(colorIndex, rgbPallet); err != nil || !matchColor(lookedUpColor, rgbPallet[colorIndex]) {
			t.Errorf("lookup of index %d failed.\nExpected: %v\n Got: %s", colorIndex, rgbPallet[colorIndex], err)
		}
	}
}

func TestUpdateColor(t *testing.T) {
	//Range
	if updatedColor, err := updateColor(-1, rgbPallet); err == nil {
		t.Errorf("update of index -1 failed.\nExpected: Error, Invalid color range\n Got: %v", updatedColor)
	}

	if updatedColor, err := updateColor(16, rgbPallet); err == nil {
		t.Errorf("update of index 16 failed.\nExpected: Error, Invalid color range\n Got: %v", updatedColor)
	}

	//Loop around
	if updatedColor, err := updateColor(15, rgbPallet); err != nil || !matchColor(updatedColor, rgbPallet[0]) {
		t.Errorf("update of index 15 failed.\nExpected: %v\n Got: Color: %v Error: %s", rgbPallet[0], updatedColor, err)
	}

	//Increment
	if updatedColor, err := updateColor(0, rgbPallet); err != nil || !matchColor(updatedColor, rgbPallet[1]) {
		t.Errorf("update of index 0 failed.\nExpected: %v\n Got: Color: %v Error: %s", rgbPallet[1], updatedColor, err)
	}
}
