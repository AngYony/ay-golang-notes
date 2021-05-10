package main

func main() {

	//闭包
	//因为是闭包，在for迭代语句中，
	//每个defer语句延迟执行的函数引用的都是同一个i迭代变量，在循环结束后这个变量的值为3，因此最终输出的都是3。
	for i := 0; i < 3; i++ {
		defer func() {
			println(i)
		}()
	}

	for i := 0; i < 3; i++ {
		i := i //定义一个循环体内局部变量i
		defer func() { println(i) }()
	}

	for i := 0; i < 3; i++ {
		//通过函数传入i
		//defer 语句会马上对调用参数求值
		defer func(i int) { println(i) }(i)
	}
}
