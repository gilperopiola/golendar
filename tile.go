package main

import (
	"image/color"
	"time"
)

// Tile represents a single day in the calendar.
type Tile struct {
	Date   time.Time
	Name   string
	Value  float64
	Color  color.Color
	Active bool
}

var (
	emptyTileColor = &color.RGBA{11, 11, 11, 255}
	todayColor     = &color.RGBA{245, 175, 55, 255}
	bgColor        = &color.RGBA{33, 33, 33, 255}
)
