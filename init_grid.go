package main

import (
	"math/rand"
	"time"
)

func initGrid(year int, month time.Month) [gridTall][gridWide]Tile {
	var grid [gridTall][gridWide]Tile

	firstOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)

	// Adjust for Monday start. More info at the bottom.
	weekdayOffset := (int(firstOfMonth.Weekday()) + 6) % 7

	day := 1
	daysInMonth := daysIn(month, year)

	for week := 0; week < gridWide; week++ {
		for dayOfWeek := 0; dayOfWeek < gridTall; dayOfWeek++ {
			tile := newEmptyTile()

			// Empty tiles before and after current month
			if week == 0 && dayOfWeek < weekdayOffset {
				grid[dayOfWeek][week] = tile
				continue
			}
			if day > daysInMonth {
				grid[dayOfWeek][week] = tile
				continue
			}

			// Tile exists
			date := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
			tile.Date = date
			tile.Name = weekdayToSpanish(date.Weekday())
			tile.Active = true

			// Set initial rnd value for now
			tile.Value = float64(rand.Intn(9)) / 2

			grid[dayOfWeek][week] = tile
			day++
		}
	}

	return grid
}

// We need this complex calc instead of just subtracting 1 from the Weekday() result
// because if not Sunday would be -1 and may cause an index out of range error.
//
// Original Weekday() | 		Calculation		  | Adjusted Index
// Sunday (0)		  | (0 + 6) % 7 = 6 % 7 = 6   | 6
// Monday (1)		  | (1 + 6) % 7 = 7 % 7 = 0   | 0
// Tuesday (2)		  | (2 + 6) % 7 = 8 % 7 = 1   | 1
// Wednesday (3)	  | (3 + 6) % 7 = 9 % 7 = 2   | 2
// Thursday (4)		  | (4 + 6) % 7 = 10 % 7 = 3  | 3
// Friday (5)		  | (5 + 6) % 7 = 11 % 7 = 4  | 4
// Saturday (6)		  | (6 + 6) % 7 = 12 % 7 = 5  | 5
