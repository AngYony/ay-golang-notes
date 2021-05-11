package main

import (
	"fmt"
	"math/big"
)

func main() {
	var distance int64 = 41.3e12 //表示41.3乘以10的12次方
	fmt.Println(distance)

	lightSpeed := big.NewInt(299792)
	secondsPerDay := big.NewInt(86400)
	distance2 := new(big.Int)
	//由于传入的数值是10进制的，所以第二个参数指定为10
	distance2.SetString("240000000000000000000000000000000000000", 10)
	fmt.Println(distance2)
	seconds := new(big.Int)
	seconds.Div(distance2, lightSpeed)
	days := new(big.Int)
	days.Div(seconds, secondsPerDay)
	fmt.Println(days)

}
