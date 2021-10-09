package fib

import (
	"fmt"
	"testing"
)

func TestFibList(t *testing.T) {
	a := 1
	b := 1
	fmt.Print(a)
	for i := 0; i < 5; i++ {

		fmt.Print(" ", b)
		temp := a
		a = b
		b = temp + a
	}

}

func TestFibWy(t *testing.T) {
	a, b := 1, 1
	fmt.Println(a, b)

	for i := 0; i < 5; i++ {
		temp := a + b
		fmt.Println(temp)
		a = b
		b = temp

	}
}
