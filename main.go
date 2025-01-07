package main

import (
	"bytes"
	"fmt"
	"image"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	rplatformer "github.com/hajimehoshi/ebiten/v2/examples/resources/images/platformer"
)

const (
	SCREEN_HEIGHT = 900
	SCREEN_WIDTH  = 900
	unit          = 1.0
	sampleRate    = 44100 // Audio sample rate
)

var (
	leftSprite   *ebiten.Image
	rightSprite  *ebiten.Image
	idleSprite   *ebiten.Image
	appleSprite  *ebiten.Image
	audioContext *audio.Context
	biteSound    *audio.Player
)

type Game struct {
	player *player
	apples []*apple
}

type player struct {
	positionX int
	positionY int
}

type apple struct {
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

	img, _, err = ebitenutil.NewImageFromFile("apple.png")
	if err != nil {
		panic(err)
	}
	appleSprite = ebiten.NewImageFromImage(img)

	audioContext = audio.NewContext(sampleRate)
	biteSoundFile, err := ebitenutil.OpenFile("bite.mp3")
	if err != nil {
		panic(err)
	}
	biteDecoded, err := mp3.DecodeWithSampleRate(sampleRate, biteSoundFile)
	if err != nil {
		panic(err)
	}
	biteSound, err = audio.NewPlayer(audioContext, biteDecoded)
	if err != nil {
		panic(err)
	}
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

	for i := 0; i < len(g.apples); {
		if g.player.checkCollision(g.apples[i]) {
			biteSound.Rewind()
			biteSound.Play()

			g.apples = append(g.apples[:i], g.apples[i+1:]...)
		} else {
			i++
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.drawPlayer(screen)

	for _, apple := range g.apples {
		apple.drawApple(screen)
	}

	tutorial := "Move: arrow keys\nLet Gopher eat all those apples!"
	msg := fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f\n%s", ebiten.ActualTPS(), ebiten.ActualFPS(), tutorial)
	ebitenutil.DebugPrint(screen, msg)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
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
	p.positionY -= 3
}

func (p *player) MoveBackwards() {
	p.positionY += 3
}

func (p *player) MoveLeft() {
	p.positionX -= 3
}

func (p *player) MoveRight() {
	p.positionX += 3
}

func (p *player) checkCollision(a *apple) bool {
	const collisionThreshold = 50
	distanceX := p.positionX - a.positionX
	distanceY := p.positionY - a.positionY
	return distanceX*distanceX+distanceY*distanceY <= collisionThreshold*collisionThreshold
}

// apple
func (a *apple) drawApple(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.3, 0.3) //
	op.GeoM.Translate(float64(a.positionX)/unit, float64(a.positionY)/unit)
	screen.DrawImage(appleSprite, op)
}

func generateRandomApples(count int) []*apple {
	rand.Seed(time.Now().UnixNano())
	apples := make([]*apple, count)
	for i := 0; i < count; i++ {
		apples[i] = &apple{
			positionX: rand.Intn(SCREEN_WIDTH),
			positionY: rand.Intn(SCREEN_HEIGHT),
		}
	}
	return apples
}

func main() {
	game := &Game{
		player: &player{
			positionX: SCREEN_WIDTH / 2,
			positionY: SCREEN_HEIGHT / 2,
		},
		apples: generateRandomApples(10),
	}
	ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Go, Gopher, Go!")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
