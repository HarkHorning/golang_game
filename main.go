package main

import (
	"golang_game/animations"
	"golang_game/spritesheet"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Sprite struct {
	Img  *ebiten.Image
	X, Y float64
}

type Game struct {
	player            *Sprite
	playerSpriteSheet *spritesheet.SpriteSheet
	playerwalk        *animations.Animation
	sprites           []*Sprite
}

func (g *Game) Update() error {

	g.playerwalk.Update()

	if ebiten.IsKeyPressed((ebiten.KeyRight)) {
		g.player.X += 2
	}
	if ebiten.IsKeyPressed((ebiten.KeyLeft)) {
		g.player.X -= 2
	}
	if ebiten.IsKeyPressed((ebiten.KeyUp)) {
		g.player.Y -= 2
	}
	if ebiten.IsKeyPressed((ebiten.KeyDown)) {
		g.player.Y += 2
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	screen.Fill(color.RGBA{110, 120, 255, 1})

	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(g.player.X, g.player.Y)

	screen.DrawImage(
		g.player.Img.SubImage(
			g.playerSpriteSheet.Rect(g.playerwalk.Frame()),
		).(*ebiten.Image),
		&opts,
	)

	opts.GeoM.Reset()

	for _, sprite := range g.sprites {
		opts.GeoM.Translate(sprite.X, sprite.Y)

		screen.DrawImage(
			sprite.Img.SubImage(
				image.Rect(0, 0, 64, 64),
			).(*ebiten.Image),
			&opts,
		)

		opts.GeoM.Reset()
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Wizard Game!")
	chara, _, err := ebitenutil.NewImageFromFile("models/img/wizard_sheet.png")
	if err != nil {
		log.Fatal(err)
	}

	playerSpriteSheet := spritesheet.NewSpriteSheet(6, 1, 64)

	playerWalk := animations.NewAnimation(0, 6, 6, 15.0)

	game := Game{
		player: &Sprite{
			Img: chara,
			X:   64.0,
			Y:   0.0,
		},

		playerSpriteSheet: playerSpriteSheet,
		playerwalk:        playerWalk,

		sprites: []*Sprite{
			{
				Img: chara,
				X:   128.0,
				Y:   0.0,
			},
		},
	}

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
