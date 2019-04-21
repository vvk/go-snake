package base

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Postion struct {
	X int
	Y int
}

//显示字符到命令行
func Rander(x int, y int, chr byte){
	GotoPostion(x, y)
	fmt.Fprintf(os.Stdout, "%c", chr)
}

//生成指定范围随机数
func RandomInt(min int, max int) int{
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max - min) + min
}

