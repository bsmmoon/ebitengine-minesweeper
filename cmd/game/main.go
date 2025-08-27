package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenW  = 320
	screenH  = 320
	tileSize = 32
	gridW    = 10
	gridH    = 10
)

type Game struct{}

func (g *Game) Update() error {
	// No updates yet
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw grid of rectangles
	for y := 0; y < gridH; y++ {
		for x := 0; x < gridW; x++ {
			ebitenutil.DrawRect(
				screen,
				float64(x*tileSize),
				float64(y*tileSize),
				tileSize-1, tileSize-1,
				color.RGBA{0x80, 0x80, 0x80, 0xff}, // gray
			)
		}
	}
}

func (g *Game) Layout(outsideW, outsideH int) (int, int) {
	return screenW, screenH
}

func main() {
	ebiten.SetWindowSize(screenW, screenH)
	ebiten.SetWindowTitle("Minesweeper")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
