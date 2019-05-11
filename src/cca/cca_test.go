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
	if updatedColor, err := updateColor(-1, len(rgbPallet)); err == nil {
		t.Errorf("update of index -1 failed.\nExpected: Error, Invalid color range\n Got: %v", updatedColor)
	}

	if updatedColor, err := updateColor(16, len(rgbPallet)); err == nil {
		t.Errorf("update of index 16 failed.\nExpected: Error, Invalid color range\n Got: %v", updatedColor)
	}

	//Loop around
	if updatedColor, err := updateColor(15, len(rgbPallet)); err != nil || updatedColor != 0 {
		t.Errorf("update of index 15 failed.\nExpected: %v\n Got: Color: %v Error: %s", 0, updatedColor, err)
	}

	//Increment
	if updatedColor, err := updateColor(0, len(rgbPallet)); err != nil || updatedColor != 1 {
		t.Errorf("update of index 0 failed.\nExpected: %v\n Got: Color: %v Error: %s", 1, updatedColor, err)
	}
}

func TestGenerateMatrix(t *testing.T) {
	testMatrix, err := generateMatrix(1000, 1000, rgbPallet)

	if len(testMatrix) != 1000 || len(testMatrix[0]) != 1000 || err != nil {
		t.Errorf("Wrong matrix dimensions .\nExpected: 1000x1000\n Got: %dx%d Error: %s", len(testMatrix), len(testMatrix[0]), err)
	}

	for c := range testMatrix[0] {
		for r := range testMatrix {
			if testMatrix[r][c] > len(rgbPallet)-1 || testMatrix[r][c] < 0 {
				t.Errorf("Erroneous matrix value found.\nExpected: %d > value >= 0\nGot: [%d][%d]=%d", len(rgbPallet)-1, r, c, testMatrix[r][c])
			}
		}
	}
}
