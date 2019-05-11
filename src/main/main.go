package main

import (
	"cca"
	"fmt"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var colorMatrix, _ = cca.GenerateMatrix(screenWidth, screenHeight, cca.RgbPallet)
var drawImage *image.RGBA

func update(screen *ebiten.Image) error {
	for r := range colorMatrix {
		for c := range colorMatrix[0] {
			RgbColor, err := cca.LookupColor(colorMatrix[r][c], cca.RgbPallet)
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
	drawImage = image.NewRGBA(image.Rect(0, 0, screenWidth, screenHeight))
	if err := ebiten.Run(update, screenWidth, screenHeight, 2, "CCA"); err != nil {
		log.Fatal(err)
	}
}
