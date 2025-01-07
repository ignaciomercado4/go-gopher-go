package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	SCREEN_HEIGHT = 900
	SCREEN_WIDTH  = 900
)

type Game struct{}

type player struct {
	positionX int
	positionY int
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {

	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x66, 0xcc, 0xff, 0xff})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

// player movement

func main() {
	game := &Game{}
	ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Go, Gopher, Go!")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
