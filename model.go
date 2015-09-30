package main

import (
	"container/list"
	"math/rand"
	"time"
)

// Speeds
const SPEED_SLOWEST = 700 * time.Millisecond
const SPEED_FASTEST = 5 * time.Millisecond

type GameState int

const (
	GAME_INTRO GameState = iota
	GAME_STARTED
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
	IsEat bool

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
		this.Snake.PushBack(Point{i + 1, snakeY})
	}
	this.SnakeToBoard(true)

	this.RandomFood()
	this.FoodToBoard(true)

	this.UpdateSpeed()

	this.Timer = time.NewTimer(this.Speed)
}

func (this *Game) Play() {
	this.SnakeToBoard(false)

	header := this.Snake.Front().Value.(Point)

	switch this.Direction {
	case DIRECTION_LEFT:
	case DIRECTION_DOWN:
	case DIRECTION_UP:
	case DIRECTION_RIGHT:
	}

	if !this.IsEat {
		this.Snake.Remove(this.Snake.Back())
	}
	this.IsEat = false
	this.SnakeToBoard(true)
}

func (this *Game) SnakeToBoard(setOrUnset bool) {
	for e := this.Snake.Front(); e != nil; e = e.Next() {
		p := e.Value.(Point)
		if setOrUnset {
			this.Board[p.Y][p.X] = 1
		} else {
			this.Board[p.Y][p.X] = 0
		}
	}
}

func (this *Game) FoodToBoard(setOrUnset bool) {
	if setOrUnset {
		this.Board[this.Food.Y][this.Food.X] = 2
	} else {
		this.Board[this.Food.Y][this.Food.X] = 0
	}
}

func (this *Game) RandomFood() {
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

func (this *Game) UpdateSpeed() {
	this.Speed = SPEED_SLOWEST - SPEED_FASTEST*time.Duration(this.Score)
}
