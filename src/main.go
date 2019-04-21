package main

import (
	"base"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	//隐藏光标
	base.HideCursor()

	base.InitMap()

	var snake base.Snake

	snake.InitSnake()

	snake.Start()

	for {
	}
}
