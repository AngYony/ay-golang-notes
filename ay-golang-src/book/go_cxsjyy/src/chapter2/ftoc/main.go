package main

import "fmt"

func main() {
	//const freezingF, boilingF = 32.0, 212.0
	//fmt.Printf("%gF = %gC\n", freezingF, fToC(freezingF)) //32F = 0C
	//fmt.Printf("%gF = %gC\n", boilingF, fToC(boilingF))   //212F = 100C
	//
	//x := 1
	//p := &x         //p是整型指针，指向x
	//fmt.Println(*p) // 输出：1
	//
	//*p = 2         //*p指代的就是x，因此相当于是：x=2
	//fmt.Println(x) //输出：2
	//
	//var w, y int
	//fmt.Println(&y)
	//z := &y
	//fmt.Println(z)
	//fmt.Println(&w == &w, &w == &y, &y == nil) //输出：true false false
	//
	//
	var p1 = wy()   //p1是一个指针
	fmt.Println(p1) //输出：0xc00005c080

	var p2 = wy()   //p2是一个指针
	fmt.Println(p2) //输出：0xc00005c088

	fmt.Println(p1 == p2) //输出：false

	m := 1
	incr(&m)              //传入m指针，m加1
	fmt.Println(m)        //输出2
	fmt.Println(incr(&m)) //再次传入m指针，m再次加1，并返回m的值，所以输出：3
	fmt.Println(m)        //输出：3
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

func wy() *int {
	v := 1
	return &v //返回一个指针
}

//参数是一个整型指针
func incr(p *int) int {
	//获取指针指向的变量，并加1
	*p++
	return *p //返回指向的变量的值
}
