package main

import "github.com/nsf/termbox-go"

type Canvas struct {
	Map    [][]termbox.Cell
	Width  int
	Height int
}

func NewCanvas(w, h int) *Canvas {
	c := new(Canvas)

	c.Width = w
	c.Height = h

	c.Map = make([][]termbox.Cell, h)
	for i := range c.Map {
		c.Map[i] = make([]termbox.Cell, w)
	}

	return c
}

func (this *Canvas) SetMap(x, y int, cell termbox.Cell) {
	this.Map[y][x] = cell
}

func (this *Canvas) ClearMap() {
	for y := 0; y < this.Height; y++ {
		for x := 0; x < this.Width; x++ {
			this.Map[y][x] = termbox.Cell{}
		}
	}
}

func (this *Canvas) FetchNotSettedPoint() []Point {
	points := make([]Point, 0, this.Width*this.Height/2)
	for y := 0; y < this.Height; y++ {
		for x := 0; x < this.Width; x++ {
			if this.Map[y][x].Ch == 0 {
				point := Point{
					X: x,
					Y: y,
				}
				points = append(points, point)
			}
		}
	}
	return points
}

func (this *Canvas) ReFlush() {
	termbox.Clear(configForeground, configBackground)

	for y := 0; y < this.Height; y++ {
		for x := 0; x < this.Width; x++ {
			termbox.SetCell(
				x, y,
				this.Map[y][x].Ch,
				this.Map[y][x].Fg,
				this.Map[y][x].Bg,
			)
		}
	}

	termbox.Flush()
}
