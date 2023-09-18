package main

import (
	_ "embed"
	"flag"
	"fmt"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	ModeTitle Mode = iota
	ModeGame
	ModeGameOver
)

const (
	screenWidth   = 1020
	screenHeight  = 860
	titleFontSize = fontSize * 1.5
	fontSize      = 24
	smallFontSize = fontSize / 2
)

var (
	positionsLane1  = float64(screenWidth/2) - 450
	positionsLane2  = float64(screenWidth/2) - 200
	positionsLane3  = float64(screenWidth/2) + 300
	positionsLane4  = float64(screenWidth/2) + 100
	titleArcadeFont font.Face
	arcadeFont      font.Face
	smallArcadeFont font.Face
	mplusNormalFont font.Face
)

type Mode int

type Game struct {
	mode           Mode
	gameoverCount  int
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
	car1           *ebiten.Image
	car1PosX       float64
	car1PosY       float64
	car2           *ebiten.Image
	car2PosX       float64
	car2PosY       float64
	car3           *ebiten.Image
	car3PosX       float64
	car3PosY       float64
}

var flagCRT = flag.Bool("crt", false, "enable the CRT effect")

var crtGo []byte

type GameWithCRTEffect struct {
	ebiten.Game

	crtShader *ebiten.Shader
}

func NewGame(crt bool) ebiten.Game {
	g := &Game{}
	g.SetupElements()
	if crt {
		return &GameWithCRTEffect{Game: g}
	}
	return g
}

func (g *GameWithCRTEffect) DrawFinalScreen(screen ebiten.FinalScreen, offscreen *ebiten.Image, geoM ebiten.GeoM) {
	if g.crtShader == nil {
		s, err := ebiten.NewShader(crtGo)
		if err != nil {
			panic(fmt.Sprintf("flappy: failed to compiled the CRT shader: %v", err))
		}
		g.crtShader = s
	}

	os := offscreen.Bounds().Size()

	op := &ebiten.DrawRectShaderOptions{}
	op.Images[0] = offscreen
	op.GeoM = geoM
	screen.DrawRectShader(os.X, os.Y, g.crtShader, op)
}

func (g *Game) moveLanes(f float64) {

	g.laneLPosY += f
	g.laneL1PosY += f
	g.laneL2PosY += f
	g.laneL3PosY += f
	g.laneL4PosY += f
	g.laneRPosY += f
	g.laneR1PosY += f
	g.laneR2PosY += f
	g.laneR3PosY += f
	g.laneR4PosY += f

	if g.laneLPosY > 940 {

		g.laneLPosX = float64((screenWidth)/2) / 2
		g.laneLPosY = -100
	} else if g.laneL1PosY > 940 {

		g.laneL1PosX = float64((screenWidth)/2) / 2
		g.laneL1PosY = -100
	} else if g.laneL2PosY > 940 {

		g.laneL2PosX = float64((screenWidth)/2) / 2
		g.laneL2PosY = -100
	} else if g.laneL3PosY > 940 {

		g.laneL3PosX = float64((screenWidth)/2) / 2
		g.laneL3PosY = -100
	} else if g.laneL4PosY > 940 {

		g.laneL4PosX = float64((screenWidth)/2) / 2
		g.laneL4PosY = -100
	}

	if g.laneRPosY > 940 {

		g.laneRPosX = float64(screenWidth)/2 + float64((screenWidth)/2)/2
		g.laneRPosY = -100
	} else if g.laneR1PosY > 940 {

		g.laneR1PosX = float64(screenWidth)/2 + float64((screenWidth)/2)/2
		g.laneR1PosY = -100
	} else if g.laneR2PosY > 940 {

		g.laneR2PosX = float64(screenWidth)/2 + float64((screenWidth)/2)/2
		g.laneR2PosY = -100
	} else if g.laneR3PosY > 940 {

		g.laneR3PosX = float64(screenWidth)/2 + float64((screenWidth)/2)/2
		g.laneR3PosY = -100
	} else if g.laneR4PosY > 940 {

		g.laneR4PosX = float64(screenWidth)/2 + float64((screenWidth)/2)/2
		g.laneR4PosY = -100
	}
}

func (g *Game) moveCars(f float64) {

	g.car1PosY += f
	if g.car1PosY > 940 {

		g.car1PosX = float64(getPoisitonLane())
		g.car1PosY = -100
	}

	g.car2PosY += f
	if g.car2PosY > 940 {

		g.car2PosX = float64(getPoisitonLane())
		g.car2PosY = -100
	}

	g.car3PosY += f
	if g.car3PosY > 940 {

		g.car3PosX = float64(getPoisitonLane())
		g.car3PosY = -100
	}
}

func (g *Game) isCarHit() bool {
	if g.playerPosY > g.car1PosY && g.playerPosY-g.car1PosY <= 125 && (math.Abs(g.playerPosX-g.car1PosX) <= 67) {

		return true
	} else if g.playerPosY > g.car2PosY && g.playerPosY-g.car2PosY <= 125 && (math.Abs(g.playerPosX-g.car2PosX) <= 67) {

		return true
	} else if g.playerPosY > g.car3PosY && g.playerPosY-g.car3PosY <= 125 && (math.Abs(g.playerPosX-g.car3PosX) <= 67) {

		return true
	}
	return false
}

func (g *Game) Update() error {

	switch g.mode {
	case ModeTitle:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.mode = ModeGame
		}
	case ModeGame:
		flag := true
		for _, k := range inpututil.PressedKeys() {
			if k == ebiten.KeyRight && g.playerPosX < float64(screenWidth)-150 {
				g.playerPosX += 8

			}

			if k == ebiten.KeyLeft && g.playerPosX > 0 {
				g.playerPosX -= 8
			}

			if k == ebiten.KeyArrowUp {
				g.moveLanes(float64(20))
				g.moveCars(float64(20))
				flag = false
			} else {
				flag = true
			}
		}

		if flag {
			g.moveLanes(float64(10))
			g.moveCars(float64(10))
		}

		if g.isCarHit() {
			g.mode = ModeGameOver
			g.gameoverCount = 30
		}

	case ModeGameOver:
		if g.gameoverCount > 0 {
			g.gameoverCount--
		}
		if g.gameoverCount == 0 && inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.SetupElements()
			g.mode = ModeTitle
		}

	}

	return nil
}

func getPoisitonLane() int {
	posLane := rand.Intn(3)

	switch posLane {
	case 0:
		return int(positionsLane1)
	case 1:
		return int(positionsLane2)
	case 2:
		return int(positionsLane3)
	case 3:
		return int(positionsLane4)
	}
	return int(positionsLane1)
}

func (g *Game) drawRoad(screen *ebiten.Image) {

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
}

func (g *Game) drawCars(screen *ebiten.Image) {

	car := &ebiten.DrawImageOptions{}
	car.GeoM.Translate(g.playerPosX, g.playerPosY)
	screen.DrawImage(g.player, car)

	car1 := &ebiten.DrawImageOptions{}
	car1.GeoM.Translate(g.car1PosX, g.car1PosY)
	screen.DrawImage(g.car1, car1)

	car2 := &ebiten.DrawImageOptions{}
	car2.GeoM.Translate(g.car2PosX, g.car2PosY)
	screen.DrawImage(g.car2, car2)

	car3 := &ebiten.DrawImageOptions{}
	car3.GeoM.Translate(g.car3PosX, g.car3PosY)
	screen.DrawImage(g.car3, car3)
}

func (g *Game) Draw(screen *ebiten.Image) {

	screen.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})
	g.drawRoad(screen)
	if g.mode != ModeTitle {
		g.drawCars(screen)
	}
	var titleTexts []string
	var texts []string
	switch g.mode {
	case ModeTitle:
		titleTexts = []string{"RACE MANIAC"}
		texts = []string{"", "", "", "", "", "", "", "PRESS SPACE TO START PLAYING"}
	case ModeGameOver:
		texts = []string{"", "GAME OVER!  PRESS SPACE TO PLAY AGAIN"}
	}
	for i, l := range titleTexts {
		x := (screenWidth - len(l)*titleFontSize) / 2
		text.Draw(screen, l, mplusNormalFont, x, (i+4)*titleFontSize, color.White)
	}
	for i, l := range texts {
		x := (screenWidth - len(l)*fontSize) / 2
		text.Draw(screen, l, mplusNormalFont, x, (i+4)*fontSize, color.White)
	}

	if g.mode == ModeTitle {
		msg := []string{
			"Race maniac,a simple 2d game implemented in go lang",
		}
		for i, l := range msg {
			x := (screenWidth - len(l)*smallFontSize) / 2
			text.Draw(screen, l, mplusNormalFont, x, screenHeight-4+(i-1)*smallFontSize, color.White)
		}
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {

	return 1020, 860
}

func main() {
	flag.Parse()

	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("HighwayManiac")

	//if err := ebiten.RunGame(g); err != nil {
	//log.Fatalf(" Run game error : %v", err)
	//}
	if err := ebiten.RunGame(NewGame(*flagCRT)); err != nil {
		panic(err)
	}
}

func (g *Game) SetupElements() {

	imgCar, _, err := ebitenutil.NewImageFromFile("c.png")
	if err != nil {
		log.Fatalf("Load image error : %v", err)
	}

	g.player = imgCar
	g.playerPosX = float64(screenWidth/2) - 500
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

	car1, _, err := ebitenutil.NewImageFromFile("c1.png")
	if err != nil {
		log.Fatalf("Load image error : %v", err)
	}

	g.car1 = car1
	g.car1PosX = float64(screenWidth/2) - 200
	g.car1PosY = float64(screenHeight/2) - 300

	car2, _, err := ebitenutil.NewImageFromFile("c2.png")
	if err != nil {
		log.Fatalf("Load image error : %v", err)
	}

	g.car2 = car2
	g.car2PosX = float64(screenWidth/2) + 350
	g.car2PosY = -100

	car3, _, err := ebitenutil.NewImageFromFile("c3.png")
	if err != nil {
		log.Fatalf("Load image error : %v", err)
	}

	g.car3 = car3
	g.car3PosX = float64(screenWidth/2) + 100
	g.car3PosY = -300

}
