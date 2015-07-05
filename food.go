package main

import (
	"math/rand"

	"github.com/nsf/termbox-go"
)

type Food struct {
	*Point
	Ch rune
	Fg termbox.Attribute
	Bg termbox.Attribute
}

func NewFood(ch rune, fg, bg termbox.Attribute) *Food {
	return &Food{
		Ch: ch,
		Fg: fg,
		Bg: bg,
	}
}

func (this *Food) RandowPos(canvas *Canvas) {
	points := canvas.FetchNotSettedPoint()
	this.Point = &points[rand.Intn(len(points))]
	this.Point.X = this.X
	this.Point.Y = this.Y
	this.Point.Ch = this.Ch
}

func (this *Food) IsEat(body *SnakeBody) bool {
	header := body.Front().Value.(*Point)
	return this.X == header.X && this.Y == header.Y
}
