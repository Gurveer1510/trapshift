package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 960
	screenHeight = 576
	gravity      = 0.5
	jumpSpeed    = -10.0
	moveSpeed    = 3.5
	playerSize   = 28
)

type Game struct {
	PlayerX        float32
	PlayerY        float32
	PlayerOnGround bool
	VelocityY      float32
	FloorY         float32
	prevTopY    float32
	prevBottomY    float32
	platforms      []Platform
}

type Platform struct {
	PlatformX float32
	PlatformY float32
	Width     float32
	Height    float32
}

func (g *Game) Update() error {
	// if ebiten.IsKeyPressed(ebiten.KeyW) {
	// 	nextY := g.PlayerY - moveSpeed

	// 	if nextY >= 0 {
	// 		g.PlayerY = nextY
	// 	}
	// }
	// if ebiten.IsKeyPressed(ebiten.KeyS) {
	// 	nextY := g.PlayerY + moveSpeed
	// 	if nextY <= screenHeight-playerSize {
	// 		g.PlayerY = nextY
	// 	}
	// }
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) && g.PlayerOnGround {
		g.VelocityY = jumpSpeed
		g.PlayerOnGround = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		nextX := g.PlayerX - moveSpeed

		if nextX >= 0 {
			g.PlayerX = nextX
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		nextX := g.PlayerX + moveSpeed

		if nextX <= screenWidth-playerSize {
			g.PlayerX = nextX
		}
	}

	g.VelocityY += gravity
	g.prevTopY = g.PlayerY
	g.prevBottomY = g.PlayerY + playerSize
	g.PlayerY += g.VelocityY
	g.PlayerOnGround = false

	for _, platform := range g.platforms {
		if CheckHorizontalOverlap(g.PlayerX, g.PlayerY, playerSize, playerSize, platform.PlatformX, platform.PlatformY, platform.Width, platform.Height) && g.prevBottomY <= platform.PlatformY && g.VelocityY >= 0 && g.PlayerY+playerSize >= platform.PlatformY {
			g.VelocityY = 0
			g.PlayerY = platform.PlatformY - playerSize
			g.PlayerOnGround = true
		}

		if CheckHorizontalOverlap(g.PlayerX, g.PlayerY, playerSize, playerSize, platform.PlatformX, platform.PlatformY, platform.Width, platform.Height) && g.prevTopY >= platform.PlatformY+platform.Height && g.VelocityY <= 0 && g.PlayerY <= platform.PlatformY+platform.Height {
			g.VelocityY = 0
			g.PlayerY = platform.PlatformY + platform.Height
		}

	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	vector.FillRect(screen, g.PlayerX, g.PlayerY, playerSize, playerSize, color.RGBA{215, 31, 190, 255}, false)
	for _, platform := range g.platforms {
		vector.FillRect(screen, platform.PlatformX, platform.PlatformY, platform.Width, platform.Height, color.RGBA{69, 163, 64, 255}, false)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	p1 := Platform{
		PlatformX: 350,
		PlatformY: 380,
		Width:     160,
		Height:    20,
	}
	p2 := Platform{
		PlatformX: 0,
		PlatformY: 450,
		Width:     420,
		Height:    126,
	}
	p3 := Platform{
		PlatformX: 520,
		PlatformY: 450,
		Width:     440,
		Height:    126,
	}
	game := &Game{
		PlayerX:   100,
		PlayerY:   200,
		FloorY:    450,
		platforms: []Platform{p1, p2, p3},
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("TrapShift")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func CheckHorizontalOverlap(x1, y1, w1, h1, x2, y2, w2, h2 float32) bool {
	if x1 < x2+w2 && x1+w1 > x2 {
		return true
	}
	return false
}

func CheckVerticalOverlap(x1, y1, w1, h1, x2, y2, w2, h2 float32) bool {
	if y1 < y2+h2 && y1+h1 > y2 {
		return true
	}
	return false
}
