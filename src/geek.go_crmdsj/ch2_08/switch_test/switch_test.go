package switch_test

import (
	"fmt"
	"testing"
)

func TestSwitch(t *testing.T) {
	for i := 1; i <= 10; i++ {
		switch i {
		case 1, 2, 3:
			fmt.Println("值1,2,3中的一个")
			break
		case 4:
			fmt.Println("值为4")
		case 5:
			fmt.Println("值为5")
		default:
			fmt.Println("大于5的值")

		}
	}

}
