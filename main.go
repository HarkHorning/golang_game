package main

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	ChaSprite *ebiten.Image
	X, Y      float64
}

func (g *Game) Update() error {

	if ebiten.IsKeyPressed((ebiten.KeyRight)) {
		g.X += 2
	}
	if ebiten.IsKeyPressed((ebiten.KeyLeft)) {
		g.X -= 2
	}
	if ebiten.IsKeyPressed((ebiten.KeyUp)) {
		g.Y -= 2
	}
	if ebiten.IsKeyPressed((ebiten.KeyDown)) {
		g.Y += 2
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	screen.Fill(color.RGBA{110, 120, 255, 1})

	// ebitenutil.DebugPrint(screen, "Good, evnin!")

	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(g.X, g.Y)

	screen.DrawImage(
		g.ChaSprite.SubImage(
			image.Rect(0, 0, 64, 64),
		).(*ebiten.Image),
		&opts,
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Good, evnin!")
	Chara, _, err := ebitenutil.NewImageFromFile("models/img/wizard_sheet.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := ebiten.RunGame(&Game{ChaSprite: Chara, X: 60, Y: 0}); err != nil {
		log.Fatal(err)
	}
}
