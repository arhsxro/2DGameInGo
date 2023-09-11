package main

import (
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const screenWidth int = 1020
const screenHeight int = 620

type Game struct {
	player         *ebiten.Image
	playerPosX     float64
	playerPosY     float64
	background     *ebiten.Image
	backgroundPosX float64
	backgroundPosY float64
	lane1          *ebiten.Image
	lane1PosX      float64
	lane1PosY      float64
	lane2          *ebiten.Image
	lane2PosX      float64
	lane2PosY      float64
	lane3          *ebiten.Image
	lane3PosX      float64
	lane3PosY      float64
}

func (g *Game) Update() error {

	for _, k := range inpututil.AppendPressedKeys([]ebiten.Key{ebiten.KeyRight, ebiten.KeyLeft}) {
		if k == ebiten.KeyRight && g.playerPosX < float64(screenWidth)-200 {
			g.playerPosX += 8
		}

		if k == ebiten.KeyLeft && g.playerPosX > 0 {
			g.playerPosX -= 8
		}
	}

	//if inpututil.AppendPressedKeys([]ebiten.Key{ebiten.KeyRight}) != nil && g.playerPosX+300 < float64(screenWidth)-150 {
	//g.playerPosX += 5
	//}

	//if inpututil.IsKeyJustPressed(ebiten.KeyLeft) && g.playerPosX-300 > 0 {
	//g.playerPosX -= 300
	//}

	//g.playerPosX += 1
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	background := &ebiten.DrawImageOptions{}
	background.GeoM.Translate(g.backgroundPosX, g.backgroundPosY)
	screen.DrawImage(g.background, background)

	lane1 := &ebiten.DrawImageOptions{}
	lane1.GeoM.Translate(g.lane1PosX, g.lane1PosY)
	screen.DrawImage(g.lane1, lane1)

	lane2 := &ebiten.DrawImageOptions{}
	lane2.GeoM.Translate(g.lane2PosX, g.lane2PosY)
	screen.DrawImage(g.lane2, lane2)

	lane3 := &ebiten.DrawImageOptions{}
	lane3.GeoM.Translate(g.lane3PosX, g.lane3PosY)
	screen.DrawImage(g.lane3, lane3)

	car := &ebiten.DrawImageOptions{}
	car.GeoM.Translate(g.playerPosX, g.playerPosY)
	screen.DrawImage(g.player, car)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {

	return 1020, 620
}

func main() {

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("HighwayManiac")

	imgCar, _, err := ebitenutil.NewImageFromFile("Car.png")
	if err != nil {
		log.Fatalf("Load image error : %v", err)
	}

	g := &Game{}
	g.player = imgCar
	g.playerPosX = float64(screenWidth/2) - 120
	g.playerPosY = float64(screenHeight/2) + 100

	imgBackground, _, err := ebitenutil.NewImageFromFile("bck.jpg")
	if err != nil {
		log.Fatalf("Load image error : %v", err)
	}

	g.background = imgBackground
	g.backgroundPosX = float64(0)
	g.backgroundPosY = float64(0)

	lane1 := ebiten.NewImage(15, 80)
	lane1.Fill(color.White)

	lane2 := ebiten.NewImage(15, 60)
	lane2.Fill(color.White)

	lane3 := ebiten.NewImage(10, 30)
	lane3.Fill(color.White)

	g.lane1 = lane1
	g.lane1PosX = float64(screenWidth)/2 + 30
	g.lane1PosY = 525

	g.lane2 = lane2
	g.lane2PosX = float64(screenWidth)/2 + 30
	g.lane2PosY = 400

	g.lane3 = lane3
	g.lane3PosX = float64(screenWidth)/2 + 30
	g.lane3PosY = 320

	if err := ebiten.RunGame(g); err != nil {
		log.Fatalf(" Run game error : %v", err)
	}
}
