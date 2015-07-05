package main

import "container/list"

type SnakeBody struct {
	*list.List
	Direction
}

func NewSnakeBody() *SnakeBody {
	body := new(SnakeBody)
	body.List = list.New()
	for i := 5; i >= 2; i-- {
		point := NewPoint(CONFIG_BODY, configForeground, configBackground, i, 8)
		body.PushBack(point)
	}
	body.Direction = DIRECTION_RIGHT

	return body
}

func (this *SnakeBody) ChangeDirection(pressedKey rune) {
	switch pressedKey {
	case CONFIG_KEY_UP:
		if this.Direction != DIRECTION_DOWN {
			this.Direction = DIRECTION_UP
		}

	case CONFIG_KEY_DOWN:
		if this.Direction != DIRECTION_UP {
			this.Direction = DIRECTION_DOWN
		}

	case CONFIG_KEY_LEFT:
		if this.Direction != DIRECTION_RIGHT {
			this.Direction = DIRECTION_LEFT
		}

	case CONFIG_KEY_RIGHT:
		if this.Direction != DIRECTION_LEFT {
			this.Direction = DIRECTION_RIGHT
		}

	}
}

func (this *SnakeBody) Move(increase bool) {
	header := this.Front().Value.(*Point)
	newHeader := *header

	switch this.Direction {
	case DIRECTION_UP:
		newHeader.Y--
	case DIRECTION_DOWN:
		newHeader.Y++
	case DIRECTION_LEFT:
		newHeader.X--
	case DIRECTION_RIGHT:
		newHeader.X++
	}

	this.PushFront(&newHeader)

	if !increase {
		this.Remove(this.Back())
	}
}

func (this *SnakeBody) Draw(canvas *Canvas) {
	for e := this.Front(); e != nil; e = e.Next() {
		point := e.Value.(*Point)
		canvas.SetMap(point.X, point.Y, point.Cell)
	}
}
