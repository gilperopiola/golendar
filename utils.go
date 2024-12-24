package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// isToday checks if a given date is today's date.
func isToday(date time.Time) bool {
	today := time.Now()
	return date.Year() == today.Year() && date.Month() == today.Month() && date.Day() == today.Day()
}

// weekdayToSpanish returns the Spanish name of a weekday.
func weekdayToSpanish(weekday time.Weekday) string {
	weekdays := []string{"Domingo", "Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado"}
	return weekdays[int(weekday)]
}

// daysIn returns the number of days in a given month and year.
func daysIn(month time.Month, year int) int {
	return time.Date(year, month+1, 0, 0, 0, 0, 0, time.Local).Day()
}

func leftClick() bool {
	return inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)
}

func newEmptyTile() Tile {
	return Tile{
		Name:   "?",
		Color:  *emptyTileColor,
		Active: false,
	}
}
