package main

import (
	"fmt"
	"image/color"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/basicfont"
)

const (
	tileSizePx    = 75
	tilePaddingPx = 9
	gridWide      = 5 // Number of weeks in all months (except February?)
	gridTall      = 7 // Days of the week
	screenWidePx  = gridWide*tileSizePx + gridWide*tilePaddingPx
	screenTallPx  = gridTall*tileSizePx + gridTall*tilePaddingPx
)

type App struct {
	MonthGrid [gridTall][gridWide]Tile
}

func NewGolendar() *App {
	return &App{
		MonthGrid: initGrid(2024, time.October),
	}
}

func (app *App) Update() error {
	if leftClick() {
		mouseX, mouseY := ebiten.CursorPosition()
		tileX := mouseX / (tileSizePx + tilePaddingPx)
		tileY := mouseY / (tileSizePx + tilePaddingPx)

		if tileX >= 0 && tileX < gridWide && tileY >= 0 && tileY < gridTall {
			tile := app.MonthGrid[tileY][tileX]

			if tile.Active {
				fmt.Printf("%s\n", tile.Date.Format(time.DateOnly))
			}
		}
	}
	return nil
}

// Draw renders the game screen.
func (app *App) Draw(screen *ebiten.Image) {
	for tileY := 0; tileY < gridTall; tileY++ {
		for tileX := 0; tileX < gridWide; tileX++ {
			tile := app.MonthGrid[tileY][tileX]

			// Screen position in px
			x := float64(tileX*tileSizePx + tileX*tilePaddingPx)
			y := float64(tileY*tileSizePx + tileX*tilePaddingPx) // tileX is not a typo. Adds the cascade effect

			// If date not in current month, render empty tile and continue
			if !tile.Active {
				vector.DrawFilledRect(screen, float32(x), float32(y), tileSizePx, tileSizePx, emptyTileColor, false)
				continue
			}

			// If we're here the date exists
			drawColor := *todayColor
			isToday := isToday(tile.Date)

			if !isToday {
				r := uint8(math.Min(20+tile.Value*60, 255))
				g := uint8(math.Min(210-tile.Value*35, 255))
				b := uint8(math.Min(140-tile.Value*25, 255))

				drawColor = color.RGBA{r, g, b, 255}
			}

			// Draw the square
			vector.DrawFilledRect(screen, float32(x), float32(y), tileSizePx, tileSizePx, drawColor, false)

			// Draw the labels
			shortDateLabel := fmt.Sprintf("%s%d", tile.Name[:1], tile.Date.Day())
			valueDateLabel := "0"
			valueDateLabelXOffset := 0 // for .5 values

			// If value is round
			if math.Mod(tile.Value, 1) < epsilon {
				valueDateLabel = fmt.Sprintf("%.0f", tile.Value)
			} else {
				// If .5
				valueDateLabel = fmt.Sprintf("%.1f", tile.Value)
				valueDateLabelXOffset = -15
			}

			shortDateLabelColor := color.White
			if isToday {
				shortDateLabelColor = color.Black
			}

			text.Draw(screen, shortDateLabel, basicfont.Face7x13, int(x)+5, int(y)+15, shortDateLabelColor)                                 // L1, M2, ...
			text.Draw(screen, valueDateLabel, robotoFont, int(x)+tileSizePx/2-8+valueDateLabelXOffset, int(y)+tileSizePx/2+14, color.White) // 0, 0.5, ...
		}
	}
}

// Layout defines the game's screen dimensions.
func (app *App) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidePx, screenTallPx
}

// epsilon is a tiny value used to safely compare floats.
const epsilon = 1e-8
