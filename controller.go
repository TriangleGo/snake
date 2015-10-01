package main

import (
	"time"

	"github.com/nsf/termbox-go"
)

const ANIMATION_SPEED = 10 * time.Millisecond

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	g := NewGame()
	Render(g)

	for {
		select {
		case ev := <-eventQueue:
			if ev.Type != termbox.EventKey {
				break
			}
			switch {
			case ev.Ch == 'h' || ev.Key == termbox.KeyArrowLeft:
				g.SetDirection(DIRECTION_LEFT)
			case ev.Ch == 'j' || ev.Key == termbox.KeyArrowDown:
				g.SetDirection(DIRECTION_DOWN)
			case ev.Ch == 'k' || ev.Key == termbox.KeyArrowUp:
				g.SetDirection(DIRECTION_UP)
			case ev.Ch == 'l' || ev.Key == termbox.KeyArrowRight:
				g.SetDirection(DIRECTION_RIGHT)
			case ev.Key == termbox.KeySpace:
				g.PauseOrRestart()
			case ev.Ch == 'q' || ev.Key == termbox.KeyEsc:
				return
			}

		case <-g.Timer.C:
			g.Play()

		default:
			Render(g)
			time.Sleep(ANIMATION_SPEED)
		}
	}
}
