package base

type Food struct {
	Postion
}

func (food *Food) Generate(snake Snake) {
	for {
		x := RandomInt(STARTX+1, ENDX-1)
		y := RandomInt(STARTY+1, ENDY-1)

		for i := 0; i < snake.length; i++ {
			if x != snake.pos[i].X && y != snake.pos[i].Y {
				food.X = x
				food.Y = y
				Rander(x, y, '*')
				return
			}
		}
	}
}
