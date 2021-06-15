package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand" //这里指的是包的导入路径
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	//Unix()方法将把时间转换成Unix时间格式，这是一个整数，是自1970年1月1日以来的秒数
	seconds := time.Now().Unix() //获取当前日期和时间的整数形式
	rand.Seed(seconds)           //播种随机数生成器

	target := rand.Intn(100) + 1 //rand是包的名称
	fmt.Println(target)

	//创建一个bufio.Reader,用于读取键盘输入
	reader := bufio.NewReader(os.Stdin)

	for guesses := 0; guesses < 10; guesses++ {

		fmt.Println("你还剩", 10-guesses, "次机会")

		fmt.Print("请输入一个数字：")
		//读取用户输入的内容，直到他们按下<enter>
		input, err := reader.ReadString('\n')

		//如果出现错误，打印错误信息
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)  //删除换行符
		guess, err := strconv.Atoi(input) //将输入字符串转换为整数
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("你猜低了")
		} else if guess > target {
			fmt.Println("你猜高了")
		} else {
			fmt.Println("猜对了")
			break
		}
	}

}
