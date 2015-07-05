package main

import (
	"log"
	"os"

	"github.com/nsf/termbox-go"
)

func main() {
	f, err := os.OpenFile("debug.log", os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	log.SetOutput(f)

	err = termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	canvas := NewCanvas(WIDTH, HEIGHT)

	drawMenuHandler(canvas)()
}

func drawMenuHandler(canvas *Canvas) func() {
	return func() {
		drawMenu(canvas, signalHandler(canvas))
	}
}

func signalHandler(canvas *Canvas) func(signal int) {
	return func(signal int) {
		switch signal {
		case 0:
			drawGame(canvas, drawMenuHandler(canvas))

		case 1:
			os.Exit(0)
		}
	}
}

func drawBlock(canvas *Canvas, c rune) {
	cell := termbox.Cell{
		Fg: configForeground,
		Bg: configBackground,
		Ch: c,
	}
	for i := 0; i < canvas.Width; i++ {
		canvas.Map[0][i] = cell
		canvas.Map[canvas.Height-1][i] = cell
	}
	for i := 0; i < canvas.Height; i++ {
		canvas.Map[i][0] = cell
		canvas.Map[i][canvas.Width-1] = cell
	}
}

const (
	VERSION = "0.1"
	NAME    = "snake"
)

const (
	WIDTH  = 32
	HEIGHT = 18
)
