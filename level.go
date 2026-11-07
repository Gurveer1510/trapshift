package main

const (
	spawnX = 100
	spawnY = 200
)

type Platform struct {
	PlatformX float32
	PlatformY float32
	Width     float32
	Height    float32
}

type Spike struct {
	Hidden    bool
	Triggered bool
	SpikeX    float32
	SpikeY    float32
	Width     float32
	Height    float32
}

func newGame() *Game {
	return &Game{
		PlayerX: spawnX,
		PlayerY: spawnY,
		platforms: []Platform{
			{PlatformX: 350, PlatformY: 380, Width: 160, Height: 20},
			{PlatformX: 0, PlatformY: 450, Width: 420, Height: 126},
			{PlatformX: 520, PlatformY: 450, Width: 440, Height: 126},
		},
		Spikes: []Spike{
			{SpikeX: 250, SpikeY: 430, Width: 24, Height: 20},
			{SpikeX: 600, SpikeY: 430, Width: 24, Height: 20, Hidden: true},
		},
	}
}
