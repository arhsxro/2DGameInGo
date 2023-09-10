package main

import (
	_ "image/jpeg"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	player *ebiten.Image
}

func (g *Game) Update() error {

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	screen.DrawImage(g.player, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {

	return 640, 480
}

func main() {

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("HighwayManiac")

	img, _, err := ebitenutil.NewImageFromFile("car.jpeg")
	if err != nil {
		log.Fatalf("Load image error : %v", err)
	}

	g := &Game{}
	g.player = img

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatalf(" Run game error : %v", err)
	}
}
