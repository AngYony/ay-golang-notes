package main

import "fmt"

func abc(mych chan string) {
	mych <- "a"
	mych <- "b"
	mych <- "c"
}

func def(mych chan string) {
	mych <- "d"
	mych <- "e"
	mych <- "f"
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go abc(ch1) //发送abc
	go def(ch2) //发送def

	fmt.Print(<-ch1)
	fmt.Print(<-ch2)
	fmt.Print(<-ch1)
	fmt.Print(<-ch2)
	fmt.Print(<-ch1)
	fmt.Print(<-ch2)
}
