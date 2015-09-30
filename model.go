package main

import (
	"container/list"
	"fmt"
	"math/rand"
	"time"
)

// Speeds
const SPEED_SLOWEST = 700 * time.Millisecond
const SPEED_FASTEST = 5 * time.Millisecond

type GameState int

const (
	GAME_STARTED GameState = iota
	GAME_PAUESD
	GAME_OVER
)

type Direction int

const (
	DIRECTION_LEFT = iota
	DIRECTION_DOWN
	DIRECTION_UP
	DIRECTION_RIGHT
)

type Point struct {
	X int
	Y int
}

type Game struct {
	Board [][]int
	State GameState
	Score int
	Speed time.Duration

	Snake *list.List
	Food  Point

	Direction Direction

	Timer *time.Timer
}

func NewGame() *Game {
	this := new(Game)
	this.ResetGame()
	return this
}

func (this *Game) ResetGame() {
	this.Board = make([][]int, BOARD_HEIGHT)
	for i := 0; i < BOARD_HEIGHT; i++ {
		this.Board[i] = make([]int, BOARD_WIDTH)
	}

	this.Direction = DIRECTION_RIGHT

	this.Snake = list.New()
	snakeY := BOARD_HEIGHT / 2
	for i := 0; i < 3; i++ {
		this.Snake.PushFront(Point{i + 1, snakeY})
	}
	this.snakeToBoard(true)

	this.randomFood()
	this.foodToBoard(true)

	this.updateSpeed()

	this.Timer = time.NewTimer(this.Speed)
}

func (this *Game) Play() {
	globalDebugText = fmt.Sprintf("Food: %v State: %d", this.Food, this.State)

	if this.State != GAME_STARTED {
		return
	}

	if this.checkFail() {
		this.Timer.Stop()
		return
	}

	this.snakeToBoard(false)

	header := this.Snake.Front().Value.(Point)
	var newHeader Point

	switch this.Direction {
	case DIRECTION_LEFT:
		newHeader = Point{header.X - 1, header.Y}
	case DIRECTION_RIGHT:
		newHeader = Point{header.X + 1, header.Y}
	case DIRECTION_DOWN:
		newHeader = Point{header.X, header.Y + 1}
	case DIRECTION_UP:
		newHeader = Point{header.X, header.Y - 1}
	}

	this.Snake.PushFront(newHeader)
	if this.checkEat() {
		this.randomFood()
		this.foodToBoard(true)
		this.Score++
		this.updateSpeed()
	} else {
		this.Snake.Remove(this.Snake.Back())
	}
	this.snakeToBoard(true)

	this.Timer.Reset(this.Speed)
}

func (this *Game) Pause() {
	switch this.State {
	case GAME_STARTED:
		this.Timer.Stop()
		this.State = GAME_PAUESD

	case GAME_PAUESD:
		this.Timer.Reset(this.Speed)
		this.State = GAME_STARTED
	}
}

func (this *Game) SetDirection(d Direction) {
	switch this.Direction {
	case DIRECTION_LEFT:
		if d == DIRECTION_RIGHT {
			return
		}
	case DIRECTION_RIGHT:
		if d == DIRECTION_LEFT {
			return
		}
	case DIRECTION_DOWN:
		if d == DIRECTION_UP {
			return
		}
	case DIRECTION_UP:
		if d == DIRECTION_DOWN {
			return
		}
	}

	this.Direction = d
}

func (this *Game) snakeToBoard(setOrUnset bool) {
	e := this.Snake.Front()
	p := e.Value.(Point)
	if setOrUnset {
		this.Board[p.Y][p.X] = 2
	} else {
		this.Board[p.Y][p.X] = 0
	}

	for {
		e = e.Next()
		if e == nil {
			break
		}
		p := e.Value.(Point)
		if setOrUnset {
			this.Board[p.Y][p.X] = 1
		} else {
			this.Board[p.Y][p.X] = 0
		}
	}
}

func (this *Game) foodToBoard(setOrUnset bool) {
	if setOrUnset {
		this.Board[this.Food.Y][this.Food.X] = 3
	} else {
		this.Board[this.Food.Y][this.Food.X] = 0
	}
}

func (this *Game) randomFood() {
	unUsedPoints := make([]Point, 0, 8)

	for y := 0; y < BOARD_HEIGHT; y++ {
		for x := 0; x < BOARD_WIDTH; x++ {
			if this.Board[y][x] == 0 {
				unUsedPoints = append(unUsedPoints, Point{x, y})
			}
		}
	}

	l := len(unUsedPoints)
	if l == 0 {
		this.Food = Point{}
		return
	}

	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(l)

	this.Food = unUsedPoints[index]
}

func (this *Game) updateSpeed() {
	this.Speed = SPEED_SLOWEST - SPEED_FASTEST*time.Duration(this.Score)
}

func (this *Game) checkEat() bool {
	header := this.Snake.Front().Value.(Point)
	return header == this.Food
}

func (this *Game) checkFail() bool {
	e := this.Snake.Front()
	header := e.Value.(Point)

	for {
		e = e.Next()
		if e == nil {
			break
		}
		p := e.Value.(Point)
		if header == p {
			return true
		}
	}

	if header.X < 0 || header.X > BOARD_WIDTH || header.Y < 0 || header.Y > BOARD_HEIGHT {
		return true
	}

	return false
}
