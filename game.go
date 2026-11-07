package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
	prevTopY       float32
	prevBottomY    float32
	platforms      []Platform
	Spikes         []Spike
	DeathCount     int
}

func (g *Game) Update() error {
	g.handleInput()

	if g.PlayerY > screenHeight {
		g.Respawn()
		return nil
	}

	g.VelocityY += gravity
	g.prevTopY = g.PlayerY
	g.prevBottomY = g.PlayerY + playerSize
	g.PlayerY += g.VelocityY
	g.PlayerOnGround = false

	if g.checkSpikeCollision() {
		g.Respawn()
		return nil
	}

	g.resolvePlatformCollisions()
	return nil
}

func (g *Game) handleInput() {
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
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Deaths: %d", g.DeathCount))
	vector.FillRect(screen, g.PlayerX, g.PlayerY, playerSize, playerSize, color.RGBA{215, 31, 190, 255}, false)
	for _, spike := range g.Spikes {
		vector.FillRect(screen, spike.SpikeX, spike.SpikeY, spike.Width, spike.Height, color.RGBA{255, 0, 0, 255}, false)
	}
	for _, platform := range g.platforms {
		vector.FillRect(screen, platform.PlatformX, platform.PlatformY, platform.Width, platform.Height, color.RGBA{69, 163, 64, 255}, false)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Respawn() {
	g.PlayerX = spawnX
	g.PlayerY = spawnY
	g.VelocityY = 0
	g.DeathCount++
	g.PlayerOnGround = false
}
