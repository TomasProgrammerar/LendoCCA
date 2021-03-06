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
	if _, err := LookupColor(-1, RgbPallet); err == nil {
		t.Errorf("lookup of index -1 failed.\nExpected: Invalid color range\n Got: %s", err)
	}

	if _, err := LookupColor(16, RgbPallet); err == nil {
		t.Errorf("lookup of index 16 failed.\nExpected: Invalid color range\n Got: %s", err)
	}

	//Sanity check
	for colorIndex := 0; colorIndex < len(RgbPallet)-1; colorIndex++ {
		if lookedUpColor, err := LookupColor(colorIndex, RgbPallet); err != nil || !matchColor(lookedUpColor, RgbPallet[colorIndex]) {
			t.Errorf("lookup of index %d failed.\nExpected: %v\n Got: %s", colorIndex, RgbPallet[colorIndex], err)
		}
	}
}

func TestUpdateColor(t *testing.T) {
	//Range
	if updatedColor, err := UpdateColor(-1, len(RgbPallet)); err == nil {
		t.Errorf("update of index -1 failed.\nExpected: Error, Invalid color range\n Got: %v", updatedColor)
	}

	if updatedColor, err := UpdateColor(16, len(RgbPallet)); err == nil {
		t.Errorf("update of index 16 failed.\nExpected: Error, Invalid color range\n Got: %v", updatedColor)
	}

	//Loop around
	if updatedColor, err := UpdateColor(15, len(RgbPallet)); err != nil || updatedColor != 0 {
		t.Errorf("update of index 15 failed.\nExpected: %v\n Got: Color: %v Error: %s", 0, updatedColor, err)
	}

	//Increment
	if updatedColor, err := UpdateColor(0, len(RgbPallet)); err != nil || updatedColor != 1 {
		t.Errorf("update of index 0 failed.\nExpected: %v\n Got: Color: %v Error: %s", 1, updatedColor, err)
	}
}

func TestGenerateMatrix(t *testing.T) {
	testMatrix, err := GenerateMatrix(1000, 500, RgbPallet)

	if len(testMatrix[0]) != 1000 || len(testMatrix) != 500 || err != nil {
		t.Errorf("Wrong matrix dimensions .\nExpected: 1000x500\n Got: %dx%d Error: %s", len(testMatrix), len(testMatrix[0]), err)
	}

	for r := range testMatrix {
		for c := range testMatrix[0] {
			if testMatrix[r][c].Value > len(RgbPallet)-1 || testMatrix[r][c].Value < 0 {
				t.Errorf("Erroneous matrix value found.\nExpected: %d > value >= 0\nGot: [%d][%d]=%d", len(RgbPallet)-1, r, c, testMatrix[r][c].Value)
			}
		}
	}
}

func TestMatrixUpdate(t *testing.T) {
	testMatrix := [][]Pixel{
		{Pixel{0, false}, Pixel{0, false}},
		{Pixel{1, false}, Pixel{15, false}},
	}

	testMatrix = UpdateMatrix(testMatrix, RgbPallet)

	if testMatrix[0][0].Value != 1 ||
		testMatrix[0][1].Value != 1 ||
		testMatrix[1][0].Value != 1 ||
		testMatrix[1][1].Value != 0 {
		t.Errorf("Failed to update matrix.\nExpected: [[1,1],[1,0]]\nGot:%v", testMatrix)
	}

	testMatrix = UpdateMatrix(testMatrix, RgbPallet)

	if testMatrix[0][0].Value != 1 ||
		testMatrix[0][1].Value != 1 ||
		testMatrix[1][0].Value != 1 ||
		testMatrix[1][1].Value != 1 {
		t.Errorf("Failed to update matrix.\nExpected: [[1,1],[1,1]]\nGot:%v", testMatrix)
	}
}
