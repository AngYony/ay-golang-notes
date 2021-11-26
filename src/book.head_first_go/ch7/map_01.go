package main

import "fmt"

func main() {
	// var myMap map[string]float64
	// myMap = make(map[string]float64)

	// ranks := make(map[string]int)

	// myMap["aa"] = 1.1
	// myMap["bb"] = 2.2

	// ranks["one"] = 1
	// ranks["two"] = 2
	// fmt.Println(myMap, ranks)

	// myMap := map[string]float64{"aa": 1.1, "bb": 2.2}
	// fmt.Println(myMap)

	// myMap2 := map[string]float64{}
	// fmt.Println(myMap2)

	// cnts := make(map[string]int)
	// cnts["aa"]++
	// cnts["aa"]++
	// cnts["bb"]++
	// fmt.Println(cnts)

	// myTestMap := map[string]int{}
	// myTestMap["A"] = 1
	// value, ok := myTestMap["A"]
	// fmt.Println(value, ok) //输出：1 true
	// myTestMap["B"] = 0     //直接复制为int类型的零值0
	// value, ok = myTestMap["B"]
	// fmt.Println(value, ok)     //输出：0 true
	// value, ok = myTestMap["C"] //直接访问未被赋过值的键
	// fmt.Println(value, ok)     //输出：0 false

	// //如果仅仅是为了判断值是否存在，可以直接使用“_”空白标识符来忽略value
	// _, ok = myTestMap["C"]
	// fmt.Println(ok) //输出：false
	// if !ok {
	// 	fmt.Println("不存在键为C的值")
	// }

	// fmt.Println(myTestMap)

	// delete(myTestMap, "A")
	// fmt.Println(myTestMap)

	// myTestMap["BB"] = 2
	// myTestMap["CC"] = 3

	myMap := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5}
	for key, val := range myMap {
		fmt.Println(key, val)
	}

}
