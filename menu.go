package main

import "github.com/nsf/termbox-go"
import "os"

type Cursor struct {
	Shape rune
	X     int
	Y     int
	Pos   int
	Size  int
}

func NewCursor(shape rune, x, y, size int) *Cursor {
	return &Cursor{
		Shape: shape,
		X:     x,
		Y:     y,
		Size:  size,
	}
}

func (this *Cursor) Draw(canvas *Canvas) {
	cell := termbox.Cell{
		Fg: configForeground,
		Bg: configBackground,
		Ch: this.Shape,
	}
	canvas.SetMap(this.X, this.Y+this.Pos, cell)
}

func (this *Cursor) ClearAll(canvas *Canvas) {
	cell := termbox.Cell{
		Fg: configForeground,
		Bg: configBackground,
	}
	for i := 0; i < this.Size; i++ {
		canvas.SetMap(this.X, this.Y+i, cell)
	}
}

func (this *Cursor) Move(canvas *Canvas, nextFunc func(signal int)) {
loop:
	for {
		if event := termbox.PollEvent(); event.Type == termbox.EventKey {
			switch event.Ch {
			case CONFIG_KEY_UP:
				this.Pos--
				if this.Pos < 0 {
					this.Pos = this.Size - 1
				}
				this.ClearAll(canvas)
				this.Draw(canvas)
				canvas.ReFlush()

			case CONFIG_KEY_DOWN:
				this.Pos++
				if this.Pos >= this.Size {
					this.Pos = 0
				}
				this.ClearAll(canvas)
				this.Draw(canvas)
				canvas.ReFlush()

			case CONFIG_KEY_RIGHT:
				break loop

			case CONFIG_KEY_QUIT:
				os.Exit(0)
			}
		}
	}

	nextFunc(this.Pos)
}

func drawMenu(canvas *Canvas, nextFunc func(signal int)) {
	title := []rune(gLangPkg["snake"])
	tips := [][]rune{
		[]rune(gLangPkg["move_tip"]),
		[]rune(gLangPkg["quit_tip"]),
	}
	items := [][]rune{
		[]rune(gLangPkg["start_game"]),
		[]rune(gLangPkg["end_game"]),
	}

	drawBlock(canvas, CONFIG_BORDER)
	drawCenter(canvas, title, 3)
	drawCenter(canvas, tips[0], 6)
	drawCenter(canvas, tips[1], 7)
	drawItem(canvas, items, 12)

	x := (canvas.Width - len(items[0]) - 4) / 2
	cursor := NewCursor(CONFIG_CURSOR, x, 12, 2)
	cursor.Draw(canvas)
	canvas.ReFlush()
	cursor.Move(canvas, nextFunc)
}

func drawCenter(canvas *Canvas, title []rune, line int) {
	l := len(title)
	pos := (canvas.Width - l - 2) / 2
	for _, t := range title {
		cell := termbox.Cell{
			Fg: configForeground,
			Bg: configBackground,
			Ch: t,
		}
		canvas.SetMap(pos, line, cell)
		pos++
	}
}

func drawItem(canvas *Canvas, items [][]rune, line int) {
	l := len(items[0])
	for i := range items {
		pos := (canvas.Width - l) / 2
		for _, t := range items[i] {
			cell := termbox.Cell{
				Fg: configForeground,
				Bg: configBackground,
				Ch: t,
			}
			canvas.SetMap(pos, line, cell)
			pos++
		}
		line++
	}
}
