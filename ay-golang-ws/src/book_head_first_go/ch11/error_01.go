package main

import (
	"fmt"
	"log"
)

type ComedyError string

func (c ComedyError) Error() string {
	return string(c)
}

type OverheatError float64

func (o OverheatError) Error() string {
	return fmt.Sprintf("值：%0.2f", o)
}

func checkTemperature(actual float64, safe float64) error {
	excess := actual - safe
	if excess > 0 {
		var wyerr error
		wyerr = OverheatError(excess)
		return wyerr
	}
	return nil
}

func main() {
	var err error //声明一个error类型的变量
	err = ComedyError("这是一个错误信息")
	fmt.Println(err)

	err = checkTemperature(121.322, 100.0)
	if err != nil {
		log.Fatal(err)
	}
}
