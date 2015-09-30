package main

import (
	"fmt"
	"strings"

	"github.com/nsf/termbox-go"
)

const IS_DEBUG = true

var globalDebugText string

// Colors
const (
	COLOR_BACKGROUND   = termbox.ColorBlue
	COLOR_BOARD        = termbox.ColorBlack
	COLOR_INSTRUCTIONS = termbox.ColorYellow
	COLOR_DANGER       = termbox.ColorRed
)

var pieceColors = []termbox.Attribute{
	termbox.ColorBlack,
	termbox.ColorGreen,
	termbox.ColorCyan,
	termbox.ColorMagenta,
}

const (
	DEFAULT_MARGIN_WIDTH  = 2
	DEFAULT_MARGIN_HEIGHT = 1

	CELL_WITDH = 2

	TITLE_START_X = DEFAULT_MARGIN_WIDTH
	TITLE_START_Y = DEFAULT_MARGIN_HEIGHT
	TITLE_END_Y   = TITLE_START_Y + 1

	BOARD_WIDTH   = 10
	BOARD_HEIGHT  = 16
	BOARD_START_X = DEFAULT_MARGIN_WIDTH
	BOARD_START_Y = TITLE_END_Y + DEFAULT_MARGIN_HEIGHT
	BOARD_END_X   = BOARD_START_X + BOARD_WIDTH*CELL_WITDH
	BOARD_END_Y   = BOARD_START_Y + BOARD_HEIGHT

	INSTRUCTIONS_START_X = BOARD_END_X + DEFAULT_MARGIN_WIDTH
	INSTRUCTIONS_START_Y = BOARD_START_Y
)

// Text in the UI
const TITLE = "Snake written in go"

var instructions = []string{
	"instructions:",
	"",
	"left:  h | ←",
	"down:  j | ↓",
	"up:    k | ↑",
	"right: l | →",
	"Pause: <Space>",
	"Exit:  q | <Esc>",
	"",
	"Score: %v",
	"Speed: %v",
	"",
	"GAME OVER",
	"",
	"DEBUG: %s",
}

func Render(g *Game) {
	termbox.Clear(COLOR_BACKGROUND, COLOR_BACKGROUND)
	tbprint(TITLE_START_X, TITLE_START_Y, COLOR_INSTRUCTIONS, COLOR_BACKGROUND, TITLE)

	for y := 0; y < BOARD_HEIGHT; y++ {
		for x := 0; x < BOARD_WIDTH; x++ {
			colorIndex := g.Board[y][x]
			cellColor := pieceColors[colorIndex]

			for i := 0; i < CELL_WITDH; i++ {
				termbox.SetCell(BOARD_START_X+CELL_WITDH*x+i, BOARD_START_Y+y, ' ', cellColor, cellColor)
			}
		}
	}

	for y, instruction := range instructions {
		textColor := COLOR_INSTRUCTIONS

		switch {
		case strings.HasPrefix(instruction, "Score:"):
			instruction = fmt.Sprintf(instruction, g.Score)

		case strings.HasPrefix(instruction, "Speed:"):
			instruction = fmt.Sprintf(instruction, g.Speed)

		case strings.HasPrefix(instruction, "GAME OVER"):
			if g.State == GAME_OVER {
				textColor = COLOR_DANGER
			} else {
				instruction = ""
			}

		case strings.HasPrefix(instruction, "DEBUG"):
			if IS_DEBUG {
				instruction = fmt.Sprintf(instruction, globalDebugText)
			} else {
				instruction = ""
			}
		}

		tbprint(INSTRUCTIONS_START_X, INSTRUCTIONS_START_Y+y, textColor, COLOR_BACKGROUND, instruction)
	}

	termbox.Flush()
}

// tbprint draws a string.
func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}
