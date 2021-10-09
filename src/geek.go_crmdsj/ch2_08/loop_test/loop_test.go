package loop_test

import (
	"fmt"
	"testing"
)

func TestWhileLoop(t *testing.T) {
	n := 1
	for n < 5 {
		fmt.Println(n)
		n++

	}
}
