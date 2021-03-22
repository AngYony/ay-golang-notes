package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	boilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g℃", c)
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func main() {
	//fmt.Println(Celsius(2.0) == Fahrenheit(2.0)) //编译错误
	fmt.Println(Celsius(2.0) == Celsius(Fahrenheit(2.0)))          //输出true
	fmt.Println(float64(Celsius(2.0)) == float64(Fahrenheit(2.0))) //输出true

	c := FToC(212.0)
	fmt.Println(c)          //输出100℃
	fmt.Println(c.String()) //输出100℃
	fmt.Printf("%v\n", c)   //输出100℃，不需要显式调用字符串
	fmt.Printf("%s\n", c)   //输出100℃
	fmt.Printf("%g\n", c)   //输出100，不调用字符串
	fmt.Println(float64(c)) //输出100，不调用字符串

}
