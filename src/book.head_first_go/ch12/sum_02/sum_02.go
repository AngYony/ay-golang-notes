package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func OpenFile(fileName string) (*os.File, error) {
	fmt.Println("正在打开文件：", fileName)
	//打开文件，并返回指向该文件的指针，以及遇到的任何错误
	return os.Open(fileName)
}

func closeFile(file *os.File) {
	fmt.Println("正在关闭文件")
	file.Close()
}

func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := OpenFile(fileName)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	//添加defer，始终执行关闭文件操作
	defer closeFile(file)
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil
}

func main() {
	numbers, err := GetFloats(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0

	for _, number := range numbers {
		sum += number
	}

	fmt.Printf("总值：%0.2f\n", sum)
}
