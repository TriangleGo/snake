package main

import (
	"os"
	"time"

	"github.com/nsf/termbox-go"
)

type Direction int

func drawGame(canvas *Canvas, comeBackFunc func()) {

	body := NewSnakeBody()
	food := NewFood(CONFIG_FOOD, configForeground, configBackground)
	food.RandowPos(canvas)

	go listenKey()

	gIsLost = false

	for {
		canvas.ClearMap()
		drawBlock(canvas, CONFIG_BORDER)

		if gIsLost {
			drawLost(canvas)
			canvas.ReFlush()
			<-gChComeBack
			break
		}

		body.ChangeDirection(gPressedKey)
		body.Move(gIsEat)
		body.Draw(canvas)

		if gIsEat {
			food.RandowPos(canvas)
		}
		food.Draw(canvas)
		gIsEat = food.IsEat(body)

		gIsLost = isLost(body, WIDTH, HEIGHT)
		canvas.ReFlush()
		time.Sleep(300 * time.Millisecond)
	}

	comeBackFunc()
}

func drawLost(canvas *Canvas) {
	tips := [][]rune{
		[]rune("你 输 了 ！ "),
		[]rune("使 用 空 格 键 返 回 菜 单 "),
	}

	y := 6
	for i := range tips {
		l := len(tips[i])
		x := (canvas.Width - l - 2) / 2
		for _, t := range tips[i] {
			cell := termbox.Cell{
				Ch: t,
				Fg: configForeground,
				Bg: configBackground,
			}
			canvas.SetMap(x, y, cell)
			x++
		}
		y += 2
	}
}

func isLost(body *SnakeBody, width, height int) bool {
	e := body.Front()
	header := e.Value.(*Point)
	x, y := header.X, header.Y

	if x == 0 || x == width-1 || y == 0 || y == height-1 {
		return true
	}

	for e = e.Next(); e != nil; e = e.Next() {
		node := e.Value.(*Point)
		if *node == *header {
			return true
		}
	}

	return false
}

func listenKey() {
loop:
	for {
		event := termbox.PollEvent()
		if event.Type == termbox.EventKey {

			switch event.Ch {
			case CONFIG_KEY_UP, CONFIG_KEY_DOWN, CONFIG_KEY_LEFT, CONFIG_KEY_RIGHT:
				gPressedKey = event.Ch

			case CONFIG_KEY_QUIT:
				os.Exit(0)
			}

			if gIsLost && event.Key == termbox.KeySpace {
				gChComeBack <- true
				break loop
			}
		}
	}
}

const (
	DIRECTION_UP = iota
	DIRECTION_DOWN
	DIRECTION_LEFT
	DIRECTION_RIGHT
)

var (
	gPressedKey rune
	gIsLost     bool
	gIsEat      bool

	gChComeBack = make(chan bool)
)
