package main

import (
	"github.com/nsf/termbox-go"
)

const (
	CONFIG_BORDER = '#'
	CONFIG_CURSOR = '*'
	CONFIG_BODY   = 'o'
	CONFIG_FOOD   = '$'
)

const (
	CONFIG_KEY_LEFT    = 'h'
	CONFIG_KEY_RIGHT   = 'l'
	CONFIG_KEY_UP      = 'k'
	CONFIG_KEY_DOWN    = 'j'
	CONFIG_KEY_QUIT    = 'q'
	CONFIG_KEY_CONFIRM = 'w'
)

var (
	configForeground = termbox.ColorWhite
	configBackground = termbox.ColorBlack
)
