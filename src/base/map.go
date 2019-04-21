package base

//初始化地图
func InitMap() {
	for i := STARTX; i <= ENDX; i++ {
		var str byte
		if i == STARTX || i == ENDX {
			str = '+'
		} else {
			str = '-'
		}

		Rander(i, STARTX, str)
		Rander(i, ENDY, str)
	}

	for i := STARTX; i <= ENDY; i++ {
		var str byte
		if i == STARTX || i == ENDY {
			str = '+'
		} else {
			str = '|'
		}

		Rander(STARTX, i, str)
		Rander(ENDX, i, str)

	}
}
