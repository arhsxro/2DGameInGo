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
const screenHeight int = 860

type Game struct {
	player         *ebiten.Image
	playerPosX     float64
	playerPosY     float64
	background     *ebiten.Image
	backgroundPosX float64
	backgroundPosY float64
	laneL          *ebiten.Image
	laneLPosX      float64
	laneLPosY      float64
	laneL1         *ebiten.Image
	laneL1PosX     float64
	laneL1PosY     float64
	laneL2         *ebiten.Image
	laneL2PosX     float64
	laneL2PosY     float64
	laneL3         *ebiten.Image
	laneL3PosX     float64
	laneL3PosY     float64
	laneL4         *ebiten.Image
	laneL4PosX     float64
	laneL4PosY     float64
	laneR          *ebiten.Image
	laneRPosX      float64
	laneRPosY      float64
	laneR1         *ebiten.Image
	laneR1PosX     float64
	laneR1PosY     float64
	laneR2         *ebiten.Image
	laneR2PosX     float64
	laneR2PosY     float64
	laneR3         *ebiten.Image
	laneR3PosX     float64
	laneR3PosY     float64
	laneR4         *ebiten.Image
	laneR4PosX     float64
	laneR4PosY     float64
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

	g.laneLPosY += 5
	g.laneL1PosY += 5
	g.laneL2PosY += 5
	g.laneL3PosY += 5
	g.laneL4PosY += 5
	g.laneRPosY += 5
	g.laneR1PosY += 5
	g.laneR2PosY += 5
	g.laneR3PosY += 5
	g.laneR4PosY += 5

	if g.laneLPosY == 940 {
		laneL := ebiten.NewImage(15, 80)
		laneL.Fill(color.White)
		g.laneL = laneL
		g.laneLPosX = float64((screenWidth)/2) / 2
		g.laneLPosY = -100
	} else if g.laneL1PosY == 940 {
		laneL1 := ebiten.NewImage(15, 80)
		laneL1.Fill(color.White)
		g.laneL1 = laneL1
		g.laneL1PosX = float64((screenWidth)/2) / 2
		g.laneL1PosY = -100
	} else if g.laneL2PosY == 940 {
		laneL2 := ebiten.NewImage(15, 80)
		laneL2.Fill(color.White)
		g.laneL2 = laneL2
		g.laneL2PosX = float64((screenWidth)/2) / 2
		g.laneL2PosY = -100
	} else if g.laneL3PosY == 940 {
		laneL3 := ebiten.NewImage(15, 80)
		laneL3.Fill(color.White)
		g.laneL3 = laneL3
		g.laneL3PosX = float64((screenWidth)/2) / 2
		g.laneL3PosY = -100
	} else if g.laneL4PosY == 940 {
		laneL4 := ebiten.NewImage(15, 80)
		laneL4.Fill(color.White)
		g.laneL4 = laneL4
		g.laneL4PosX = float64((screenWidth)/2) / 2
		g.laneL4PosY = -100
	}

	if g.laneRPosY == 940 {
		laneR := ebiten.NewImage(15, 80)
		laneR.Fill(color.White)
		g.laneR = laneR
		g.laneRPosX = float64(screenWidth)/2 + float64((screenWidth)/2)/2
		g.laneRPosY = -100
	} else if g.laneR1PosY == 940 {
		laneR1 := ebiten.NewImage(15, 80)
		laneR1.Fill(color.White)
		g.laneR1 = laneR1
		g.laneR1PosX = float64(screenWidth)/2 + float64((screenWidth)/2)/2
		g.laneR1PosY = -100
	} else if g.laneR2PosY == 940 {
		laneR2 := ebiten.NewImage(15, 80)
		laneR2.Fill(color.White)
		g.laneR2 = laneR2
		g.laneR2PosX = float64(screenWidth)/2 + float64((screenWidth)/2)/2
		g.laneR2PosY = -100
	} else if g.laneR3PosY == 940 {
		laneR3 := ebiten.NewImage(15, 80)
		laneR3.Fill(color.White)
		g.laneR3 = laneR3
		g.laneR3PosX = float64(screenWidth)/2 + float64((screenWidth)/2)/2
		g.laneR3PosY = -100
	} else if g.laneR4PosY == 940 {
		laneR4 := ebiten.NewImage(15, 80)
		laneR4.Fill(color.White)
		g.laneR4 = laneR4
		g.laneR4PosX = float64(screenWidth)/2 + float64((screenWidth)/2)/2
		g.laneR4PosY = -100
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	background := &ebiten.DrawImageOptions{}
	background.GeoM.Translate(g.backgroundPosX, g.backgroundPosY)
	screen.DrawImage(g.background, background)

	laneL := &ebiten.DrawImageOptions{}
	laneL.GeoM.Translate(g.laneLPosX, g.laneLPosY)
	screen.DrawImage(g.laneL, laneL)

	laneL1 := &ebiten.DrawImageOptions{}
	laneL1.GeoM.Translate(g.laneL1PosX, g.laneL1PosY)
	screen.DrawImage(g.laneL1, laneL1)

	laneL2 := &ebiten.DrawImageOptions{}
	laneL2.GeoM.Translate(g.laneL2PosX, g.laneL2PosY)
	screen.DrawImage(g.laneL2, laneL2)

	laneL3 := &ebiten.DrawImageOptions{}
	laneL3.GeoM.Translate(g.laneL3PosX, g.laneL3PosY)
	screen.DrawImage(g.laneL3, laneL3)

	laneL4 := &ebiten.DrawImageOptions{}
	laneL4.GeoM.Translate(g.laneL4PosX, g.laneL4PosY)
	screen.DrawImage(g.laneL4, laneL4)

	laneR := &ebiten.DrawImageOptions{}
	laneR.GeoM.Translate(g.laneRPosX, g.laneRPosY)
	screen.DrawImage(g.laneR, laneR)

	laneR1 := &ebiten.DrawImageOptions{}
	laneR1.GeoM.Translate(g.laneR1PosX, g.laneR1PosY)
	screen.DrawImage(g.laneR1, laneR1)

	laneR2 := &ebiten.DrawImageOptions{}
	laneR2.GeoM.Translate(g.laneR2PosX, g.laneR2PosY)
	screen.DrawImage(g.laneR2, laneR2)

	laneR3 := &ebiten.DrawImageOptions{}
	laneR3.GeoM.Translate(g.laneR3PosX, g.laneR3PosY)
	screen.DrawImage(g.laneR3, laneR3)

	laneR4 := &ebiten.DrawImageOptions{}
	laneR4.GeoM.Translate(g.laneR4PosX, g.laneR4PosY)
	screen.DrawImage(g.laneR4, laneR4)

	car := &ebiten.DrawImageOptions{}
	car.GeoM.Translate(g.playerPosX, g.playerPosY)
	screen.DrawImage(g.player, car)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {

	return 1020, 860
}

func main() {

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("HighwayManiac")

	g := &Game{}
	g.SetupElements()

	if err := ebiten.RunGame(g); err != nil {
		log.Fatalf(" Run game error : %v", err)
	}
}

func (g *Game) SetupElements() {

	imgCar, _, err := ebitenutil.NewImageFromFile("c.png")
	if err != nil {
		log.Fatalf("Load image error : %v", err)
	}

	g.player = imgCar
	g.playerPosX = float64(screenWidth/2) - 120
	g.playerPosY = float64(screenHeight/2) + 200

	imgBackground, _, err := ebitenutil.NewImageFromFile("back.png")
	if err != nil {
		log.Fatalf("Load image error : %v", err)
	}

	g.background = imgBackground
	g.backgroundPosX = float64(0)
	g.backgroundPosY = float64(0)

	lane := ebiten.NewImage(13, 80)
	lane.Fill(color.White)

	g.laneL = lane
	g.laneLPosX = float64((screenWidth)/2) / 2
	g.laneLPosY = 700

	g.laneL1 = lane
	g.laneL1PosX = float64((screenWidth)/2) / 2
	g.laneL1PosY = 500

	g.laneL2 = lane
	g.laneL2PosX = float64((screenWidth)/2) / 2
	g.laneL2PosY = 300

	g.laneL3 = lane
	g.laneL3PosX = float64((screenWidth)/2) / 2
	g.laneL3PosY = 100

	g.laneL4 = lane
	g.laneL4PosX = float64((screenWidth)/2) / 2
	g.laneL4PosY = -100

	g.laneR = lane
	g.laneRPosX = float64(screenWidth)/2 + float64((screenWidth)/2)/2
	g.laneRPosY = 700

	g.laneR1 = lane
	g.laneR1PosX = float64(screenWidth)/2 + float64((screenWidth)/2)/2
	g.laneR1PosY = 500

	g.laneR2 = lane
	g.laneR2PosX = float64(screenWidth)/2 + float64((screenWidth)/2)/2
	g.laneR2PosY = 300

	g.laneR3 = lane
	g.laneR3PosX = float64(screenWidth)/2 + float64((screenWidth)/2)/2
	g.laneR3PosY = 100

	g.laneR4 = lane
	g.laneR4PosX = float64(screenWidth)/2 + float64((screenWidth)/2)/2
	g.laneR4PosY = -100

}
