package main

import (
	"fmt"
	"log"

	"github.com/ghtalpo/ebpvm/pvm"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	logicalScreenWidth   = 640
	logicalScreenHeight  = 400
	physicalScreenWidth  = 640 * 2
	physicalScreenHeight = 400 * 2
)

// Game implements ebiten.Game interface.
type Game struct{}

// NewGame is
func NewGame() *Game {
	return &Game{}
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Write your game's logical update.
	return pvm.OnUpdate()
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
	pvm.OnDraw(screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return logicalScreenWidth, logicalScreenHeight
}

func main() {
	game := NewGame()
	pvm.Start()
	// Sepcify the window size as you like. Here, a doulbed size is specified.
	ebiten.SetWindowSize(physicalScreenWidth, physicalScreenHeight)
	ebiten.SetWindowTitle("EbPvm")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
	pvm.Stop()

	fmt.Println("done")
}
