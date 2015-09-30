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
	render(g)

	for {
		select {
		case ev := <-eventQueue:
			if ev.Type != termbox.EventKey {
				break
			}
			switch {
			case ev.Ch == 'h' || ev.Key == termbox.KeyArrowLeft:
			case ev.Ch == 'j' || ev.Key == termbox.KeyArrowDown:
			case ev.Ch == 'k' || ev.Key == termbox.KeyArrowUp:
			case ev.Ch == 'l' || ev.Key == termbox.KeyArrowRight:
			case ev.Key == termbox.KeySpace:
			case ev.Ch == 'q' || ev.Key == termbox.KeyEsc:
				return
			}

		default:
			render(g)
			time.Sleep(ANIMATION_SPEED)
		}
	}
}
