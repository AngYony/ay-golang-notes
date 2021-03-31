package main

import "fmt"

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Whistle_MakeSound:", w)
}

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("Horn_MakeSound:", h)
}

type NoiseMaker interface {
	MakeSound()
}

type Robat string

func (r Robat) MakeSound() {
	fmt.Println("Robat_MakeSound:", r)
}

//定义一个额外的方法
func (r Robat) Walk() {
	fmt.Println("Robat_Walk:", r)
}

func play(n NoiseMaker) {
	n.MakeSound()
}

func main() {
	//声明一个接口类型变量
	var toy NoiseMaker
	toy = Whistle("AA")
	toy.MakeSound()

	toy = Horn("BB")
	toy.MakeSound()

	play(Whistle("CC"))
	play(Horn("DD"))
	play(Robat("EE"))

	var noiseMaker NoiseMaker = Robat("FFF")
	noiseMaker.MakeSound()

	var robot Robat = noiseMaker.(Robat)
	robot.Walk()

	wy, ok := noiseMaker.(Robat)
	if ok {
		wy.Walk()
	} else {
		fmt.Println("失败")
	}

}
