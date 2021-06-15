package main

import (
	prose "book_head_first_go/ch14/prose"
	"fmt"
)

func main() {
	phrases := []string{"AAA", "BBB"}
	fmt.Println("CCC", prose.JoinWithCommas(phrases))
}
