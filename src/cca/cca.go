package cca

import (
	"errors"
	"log"
	"math/rand"
	"time"
)

//Represents a standard RGB color value
type color struct {
	R uint8
	G uint8
	B uint8
}

//RgbPallet is the color lookup table to deduce what colors the
//matrix indexes correspond to
//table taken from https://www.december.com/html/spec/color16codes.html
var RgbPallet = []color{
	color{0, 0, 0},       //Black
	color{128, 128, 128}, //Grey
	color{192, 192, 192}, //Silver
	color{255, 255, 255}, //White
	color{128, 0, 0},     //Maroon
	color{255, 0, 0},     //red
	color{128, 128, 0},   //olive
	color{255, 255, 0},   //yellow
	color{0, 128, 0},     //green
	color{0, 255, 0},     //lime
	color{0, 128, 128},   //teal
	color{0, 255, 255},   //aqua
	color{0, 0, 128},     //navy
	color{0, 0, 255},     //blue
	color{128, 0, 128},   //purple
	color{255, 0, 255},   //magenta
}

func newMatrix(width, height int) [][]int {
	colorMatrix := make([][]int, height)
	for r := range colorMatrix {
		colorMatrix[r] = make([]int, width)
	}
	return colorMatrix
}

//make sure we are inbounds
func checkOob(r, c, rOffset, cOffset, rDim, cDim int) bool {
	return r+rOffset < 0 ||
		r+rOffset >= rDim ||
		c+cOffset < 0 ||
		c+cOffset >= cDim
}

//Checks all pixels around the current to see if it should be incremented
func checkAdjecency(r, c int, colorMatrix [][]int) bool {
	if !checkOob(r, c, -1, 0, len(colorMatrix), len(colorMatrix[0])) {
		if shouldIncrement(colorMatrix[r][c], colorMatrix[r-1][c]) {
			return true
		}
	}
	if !checkOob(r, c, 1, 0, len(colorMatrix), len(colorMatrix[0])) {
		if shouldIncrement(colorMatrix[r][c], colorMatrix[r+1][c]) {
			return true
		}
	}
	if !checkOob(r, c, 0, 1, len(colorMatrix), len(colorMatrix[0])) {
		if shouldIncrement(colorMatrix[r][c], colorMatrix[r][c+1]) {
			return true
		}
	}
	if !checkOob(r, c, 0, -1, len(colorMatrix), len(colorMatrix[0])) {
		if shouldIncrement(colorMatrix[r][c], colorMatrix[r][c-1]) {
			return true
		}
	}
	return false
}

func shouldIncrement(c1, c2 int) bool {
	return c1+1 == c2 || (c1 == 15 && c2 == 0)
}

//LookupColor takes an matrix index and a pallet and returns the color corresponding to the index
func LookupColor(colorIndex int, pallet []color) (color, error) {
	if colorIndex > len(pallet)-1 || colorIndex < 0 {
		return color{}, errors.New("Invalid color index range")
	}

	return pallet[colorIndex], nil
}

//UpdateColor takes a matrix index and updates it to the next logical step
func UpdateColor(colorIndex, maxColors int) (int, error) {
	if colorIndex > maxColors-1 || colorIndex < 0 {
		return 0, errors.New("Invalid color index range")
	} else if colorIndex == maxColors-1 {
		return 0, nil
	}

	return colorIndex + 1, nil
}

//GenerateMatrix creates a widthxheight	matrix index values within the provided color pallet
func GenerateMatrix(width, height int, pallet []color) ([][]int, error) {
	colorMatrix := newMatrix(width, height)

	source := rand.NewSource(time.Now().UnixNano())
	colorGen := rand.New(source)

	for r := range colorMatrix {
		for c := range colorMatrix[0] {
			colorMatrix[r][c] = colorGen.Intn(len(pallet))
		}
	}

	return colorMatrix, nil
}

func UpdateMatrix(colorMatrix [][]int, pallet []color) [][]int {
	updatedMatrix := newMatrix(len(colorMatrix[0]), len(colorMatrix))
	var err error

	for r := range colorMatrix {
		for c := range colorMatrix[0] {
			if checkAdjecency(r, c, colorMatrix) {
				if updatedMatrix[r][c], err = UpdateColor(colorMatrix[r][c], len(pallet)); err != nil {
					log.Fatal(err)
				}
			}
		}
	}

	return updatedMatrix
}
