package main

import "github.com/nsf/termbox-go"

const (
	// Colors
	COLOR_BACKGROUND   = termbox.ColorBlue
	COLOR_BOARD        = termbox.ColorBlack
	COLOR_INSTRUCTIONS = termbox.ColorYellow
)

const (
	DEFAULT_MARGIN_WIDTH  = 2
	DEFAULT_MARGIN_HEIGHT = 1
	TITLE_START_X         = DEFAULT_MARGIN_WIDTH
	TITLE_START_Y         = DEFAULT_MARGIN_HEIGHT
)

// Text in the UI
const TITLE = "Snake written in go"

func render(g *Game) {
	termbox.Clear(COLOR_BACKGROUND, COLOR_BACKGROUND)
	tbprint(TITLE_START_X, TITLE_START_Y, COLOR_INSTRUCTIONS, COLOR_BACKGROUND, TITLE)

	termbox.Flush()
}

// tbprint draws a string.
func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}
