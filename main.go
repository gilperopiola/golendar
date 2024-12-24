package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := NewGolendar()

	ebiten.SetWindowSize(screenWidePx, screenTallPx)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Golendar")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatalf("Error running Golendar: %v", err)
	}
}
