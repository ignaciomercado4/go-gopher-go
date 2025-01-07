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

type Game struct {
	player *player
}

type player struct {
	positionX int
	positionY int
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.player.MoveForward()
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.player.MoveBackwards()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch {
	case g.player.positionY == 0:
		screen.Fill(color.Black)
	case g.player.positionY > 0 && g.player.positionY <= 10:
		screen.Fill(color.RGBA{0x66, 0xcc, 0xff, 0xff})
	case g.player.positionY > 10:
		screen.Fill(color.White)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

// player movement
func (p *player) MoveForward() {
	p.positionY++

}

func (p *player) MoveBackwards() {
	p.positionY--
}

func main() {
	game := &Game{
		player: &player{
			positionX: 0,
			positionY: 0,
		},
	}
	ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Go, Gopher, Go!")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
