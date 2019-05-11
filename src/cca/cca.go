package cca

import (
	"errors"
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
		return -1, errors.New("Invalid color index range")
	} else if colorIndex == maxColors-1 {
		return 0, nil
	}

	return colorIndex + 1, nil
}

//GenerateMatrix creates a widthxheight	matrix index values within the provided color pallet
func GenerateMatrix(width, height int, pallet []color) ([][]int, error) {
	colorMatrix := make([][]int, height)
	for r := range colorMatrix {
		colorMatrix[r] = make([]int, width)
	}

	source := rand.NewSource(time.Now().UnixNano())
	colorGen := rand.New(source)

	for c := range colorMatrix[0] {
		for r := range colorMatrix {
			colorMatrix[r][c] = colorGen.Intn(len(pallet))
		}
	}

	return colorMatrix, nil
}
