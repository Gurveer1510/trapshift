package main

func CheckHorizontalOverlap(x1, y1, w1, h1, x2, y2, w2, h2 float32) bool {
	return x1 < x2+w2 && x1+w1 > x2
}

func CheckVerticalOverlap(x1, y1, w1, h1, x2, y2, w2, h2 float32) bool {
	return y1 < y2+h2 && y1+h1 > y2
}

func (g *Game) resolvePlatformCollisions() {
	for _, platform := range g.platforms {
		if !CheckHorizontalOverlap(g.PlayerX, g.PlayerY, playerSize, playerSize,
			platform.PlatformX, platform.PlatformY, platform.Width, platform.Height) {
			continue
		}

		top := platform.PlatformY
		bottom := top + platform.Height
		if g.prevBottomY <= top && g.VelocityY >= 0 && g.PlayerY+playerSize >= top {
			g.VelocityY = 0
			g.PlayerY = top - playerSize
			g.PlayerOnGround = true
		}
		if g.prevTopY >= bottom && g.VelocityY <= 0 && g.PlayerY <= bottom {
			g.VelocityY = 0
			g.PlayerY = bottom
		}
	}
}

func (g *Game) checkSpikeCollision() bool {
	for _, spike := range g.Spikes {
		if CheckHorizontalOverlap(g.PlayerX, g.PlayerY, playerSize, playerSize,
			spike.SpikeX, spike.SpikeY, spike.Width, spike.Height) &&
			CheckVerticalOverlap(g.PlayerX, g.PlayerY, playerSize, playerSize,
				spike.SpikeX, spike.SpikeY, spike.Width, spike.Height) && (!spike.Hidden || spike.Triggered)  {
			return true
		}
	}
	return false
}

func (g *Game) ProximityCheck() {
	for i, spike := range g.Spikes {
		if spike.SpikeX - (g.PlayerX+playerSize) <= 50 {
			g.Spikes[i].Triggered = true
		}
	}
}
