package base

import (
	"fmt"
	"math/rand"
	"time"
)

type Snake struct {
	length        int
	pos           []Postion
	direction     int
	lastDirection int
	score         int
	status        int
}

//var gameOverChan chan bool= make(chan bool)

func (snake *Snake) InitSnake() {
	snake.length = SNAKE_LENGTH
	snake.direction = getSnakeDirection()
	snake.score = 0
	snake.status = GAME_INIT

	snake.pos = make([]Postion, SNAKE_LENGTH)
	snake.pos[0] = Postion{X: SNAKE_START_X, Y: SNAKE_START_Y}

	x, y := snake.pos[0].X, snake.pos[0].Y
	for i := 1; i < SNAKE_LENGTH; i++ {
		switch snake.direction {
		case UP:
			y--
		case RIGHT:
			x++
		case DOWN:
			y++
		case LEFT:
			x--
		}
		snake.pos[i] = Postion{X: x, Y: y}
	}

	snake.rander()
}

func (snake *Snake) Start() {
	snake.status = GAME_RUNING
	var food Food
	food.Generate(*snake)

	go snake.setSnakeDirection()

	for {
		time.Sleep(time.Second / 3)

		if snake.direction == GAME_PAUSE {
			continue
		}

		snakeHead := snake.pos[0]

		//检测墙
		if snakeHead.X <= STARTX || snakeHead.X >= ENDX || snakeHead.Y <= STARTY || snakeHead.Y >= ENDY {
			snake.over()
			return
		}

		//检测自身
		for i := 1; i < snake.length; i++ {
			if snakeHead.X == snake.pos[i].X && snakeHead.Y == snake.pos[i].Y {
				snake.over()
				return
			}
		}

		//食物检测
		if snakeHead.X == food.X && snakeHead.Y == food.Y {
			snake.graw()

			Rander(food.X, food.Y, ' ')
			food.Generate(*snake)
		}

		snake.move()
	}
}

//移动
func (snake *Snake) move() {
	snakeTail := snake.pos[snake.length-1]
	Rander(snakeTail.X, snakeTail.Y, ' ')

	for i := snake.length - 1; i > 0; i-- {
		snake.pos[i] = snake.pos[i-1]
	}

	snakeHead := snake.pos[0]
	switch snake.direction {
	case UP:
		snakeHead.Y--
	case RIGHT:
		snakeHead.X++
	case DOWN:
		snakeHead.Y++
	case LEFT:
		snakeHead.X--
	}

	snake.pos[0] = snakeHead
	snake.rander()
}

func (snake *Snake) rander() {
	for i := 0; i < snake.length; i++ {
		var chr byte = '*'
		if i == 0 {
			chr = '#'
		}
		Rander(snake.pos[i].X, snake.pos[i].Y, chr)
	}
}

//增长
func (snake *Snake) graw() {

	snakeNew := snake.pos[snake.length-1]
	switch snake.direction {
	case UP:
		snakeNew.Y++
	case RIGHT:
		snakeNew.X--
	case DOWN:
		snakeNew.Y--
	case LEFT:
		snakeNew.X++
	}

	snake.length++
	snake.score++

	snake.pos = append(snake.pos, snakeNew)
}

func (snake *Snake) over() {
	//close(gameOverChan)
	snake.status = GAME_OVER
	GotoPostion(STARTX, ENDY+1)
	fmt.Println("得分：", snake.score)
}

//根据键盘输入改变方便
func (snake *Snake) setSnakeDirection() {
	for {
		/*select {
		case <-gameOverChan:
			return
		default:
		}*/

		if snake.status == GAME_OVER {
			return
		}

		snake.status = GAME_RUNING
		switch Direction() {
		//方向上  W|w|↑
		case 72, 87, 119:
			if snake.direction != DOWN {
				snake.direction = UP
			}
		//方向右
		case 100, 68, 77:
			if snake.direction != LEFT {
				snake.direction = RIGHT
			}
		//方向下
		case 83, 115, 80:
			if snake.direction != UP {
				snake.direction = DOWN
			}
		//方向左
		case 65, 97, 75:
			if snake.direction != RIGHT {
				snake.direction = LEFT
			}
		//暂停  空格键
		case 32:
			if snake.direction == GAME_PAUSE {
				snake.direction = snake.lastDirection
			} else {
				snake.status = GAME_PAUSE
				snake.lastDirection = snake.direction
				snake.direction = GAME_PAUSE
			}
		}
	}
}

func getSnakeDirection() int {
	directorList := [4]int{UP, RIGHT, DOWN, LEFT}
	return rand.Intn(len(directorList))
}
