package main

import (
	"bytes"
	"fmt"
	"image"
	"log"

	rplatformer "github.com/hajimehoshi/ebiten/v2/examples/resources/images/platformer"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	SCREEN_HEIGHT = 900
	SCREEN_WIDTH  = 900
	unit          = 0.5
)

var (
	leftSprite  *ebiten.Image
	rightSprite *ebiten.Image
	idleSprite  *ebiten.Image
)

type Game struct {
	player *player
}

type player struct {
	positionX int
	positionY int
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(rplatformer.Right_png))
	if err != nil {
		panic(err)
	}
	rightSprite = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(rplatformer.Left_png))
	if err != nil {
		panic(err)
	}
	leftSprite = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(rplatformer.MainChar_png))
	if err != nil {
		panic(err)
	}
	idleSprite = ebiten.NewImageFromImage(img)
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.player.MoveForward()
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.player.MoveBackwards()
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.player.MoveLeft()
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.player.MoveRight()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.drawPlayer(screen)

	tutorial := "Move: arrow keys\nLet Gopher eat all those apples!"
	msg := fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f\n%s", ebiten.ActualTPS(), ebiten.ActualFPS(), tutorial)
	ebitenutil.DebugPrint(screen, msg)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 900, 900
}

// player
func (p *player) drawPlayer(screen *ebiten.Image) {
	s := idleSprite
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Scale(0.5, 0.5)
	op.GeoM.Translate(float64(p.positionX)/unit, float64(p.positionY)/unit)
	screen.DrawImage(s, op)
}

func (p *player) MoveForward() {
	p.positionY--
}

func (p *player) MoveBackwards() {
	p.positionY++
}

func (p *player) MoveLeft() {
	p.positionX--
}

func (p *player) MoveRight() {
	p.positionX++
}

func main() {
	game := &Game{
		player: &player{
			positionX: 450,
			positionY: 450,
		},
	}
	ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Go, Gopher, Go!")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
