package main

func greeting2(myChannel chan string) {
	myChannel <- "hi" //发送操作会导致该goroutine阻塞

}

func main() {
	myChannel := make(chan string)
	//myChannel <- "你看"
	go greeting2(myChannel)
	//fmt.Println(<-myChannel)
}
