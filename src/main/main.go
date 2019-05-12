package main

import (
	"cca"
	"flag"
	"fmt"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var colorMatrix [][]cca.Pixel
var drawImage *image.RGBA

func update(screen *ebiten.Image) error {
	colorMatrix = cca.UpdateMatrix(colorMatrix, cca.RgbPallet)

	for r := range colorMatrix {
		for c := range colorMatrix[0] {
			RgbColor, err := cca.LookupColor(colorMatrix[r][c].Value, cca.RgbPallet)
			if err != nil {
				log.Fatal(err)
			}

			drawImage.Pix[4*c+r*4*(len(colorMatrix[0])-1)] = RgbColor.R
			drawImage.Pix[4*c+1+r*4*(len(colorMatrix[0])-1)] = RgbColor.G
			drawImage.Pix[4*c+2+r*4*(len(colorMatrix[0])-1)] = RgbColor.B
			drawImage.Pix[4*c+3+r*4*(len(colorMatrix[0])-1)] = 0xff
		}
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	screen.ReplacePixels(drawImage.Pix)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
	return nil
}

func main() {
	threshold := flag.Int("threshold", 1, "required number of preceeding neighbours required before transforming")
	speed := flag.Int("speed", 10, "number of updates per second to the simulation (cap at 60)")
	width := flag.Int("width", 640, "screen width")
	height := flag.Int("height", 480, "screen height")
	colors := flag.Int("colors", 16, "total number of colors used in the simulation")
	flag.Parse()

	cca.InitSimParams(*threshold, *speed, *colors)
	colorMatrix, _ = cca.GenerateMatrix(*width, *height, cca.RgbPallet)
	ebiten.SetMaxTPS(*speed)
	drawImage = image.NewRGBA(image.Rect(0, 0, *width, *height))
	if err := ebiten.Run(update, *width, *height, 2, "CCA"); err != nil {
		log.Fatal(err)
	}
}
