package main

import (
	"github.com/nsf/termbox-go"
)

type Point struct {
	termbox.Cell
	X int
	Y int
}

func NewPoint(ch rune, fg, bg termbox.Attribute, x, y int) *Point {
	return &Point{
		Cell: termbox.Cell{
			Ch: ch,
			Fg: fg,
			Bg: bg,
		},
		X: x,
		Y: y,
	}
}

func (this *Point) Draw(canvas *Canvas) {
	canvas.SetMap(this.X, this.Y, this.Cell)
}
