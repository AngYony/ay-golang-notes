package main

import (
	"fmt"
	"image"
	"log"
	"time"
)

type command int

//模拟两个command常量
const (
	right = command(0)
	left  = command(1)
)

type RoverDriver struct {
	//定义一个发送命令的通道
	commandc chan command
}

//定义向左的方法
func (r *RoverDriver) Left() {
	//向通道发送left命令值
	r.commandc <- left
}

//定义向右的方法
func (r *RoverDriver) Right() {
	//向通道发送right命令值
	r.commandc <- right
}

//定义结构的drive方法，能够访问RoverDriver的任何成员
func (r *RoverDriver) drive() {
	//当前位置初始值
	pos := image.Point{X: 0, Y: 0}
	//当前方向
	direction := image.Point{X: 1, Y: 0}

	updateInterval := 250 * time.Millisecond
	//创建初始计时器通道
	nextMove := time.After(updateInterval)

	for {
		select {
		//等待接收来自命令通道的命令
		case c := <-r.commandc:
			//判断命令的值，执行不同的分支操作
			switch c {
			//向右转
			case right:
				direction = image.Point{
					X: -direction.Y,
					Y: direction.X,
				}
				//向左转
			case left:
				direction = image.Point{
					X: direction.Y,
					Y: direction.X,
				}
			}
			log.Printf("new direction %v", direction)

		case <-nextMove: //从通道中取到值后将会击发计时器
			pos = pos.Add(direction)
			fmt.Println("当前位置：", pos)

			//为下一次事件循环创建新的计时器通道
			nextMove = time.After(updateInterval)
		}
	}
}

//创建通道并启动工作进程
func NewRoverDriver() *RoverDriver {
	r := &RoverDriver{
		commandc: make(chan command),
	}
	go r.drive()
	return r
}

func main() {
	r := NewRoverDriver()
	//此处休眠3秒，将始终触发time.after通道，将会连续输出“当前位置”信息，直到3秒结束
	time.Sleep(3 * time.Second)
	r.Left()
	time.Sleep(3 * time.Second)
	r.Right()
	time.Sleep(3 * time.Second)
}
