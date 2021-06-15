package main

import (
	"encoding/json"
	"fmt"
)

type OrderItem struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
type Order struct {
	ID         string    `json:"id"`
	Name       string    `json:"name,omitempty"` //省略空的字段
	Quantity   int       `json:"quantity"`
	TotalPrice float64   `json:"total_price"`
	Item       OrderItem `json:"item"`
}

func unmarshal() {
	s := `{"id":"1234","quantity":3,"total_price":30,"item":{"id":"item1","name":"小三","price":100}}`
	var o Order
	err := json.Unmarshal([]byte(s), &o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", o)
}

func marshal() {
	o := Order{
		ID:         "1234",
		Quantity:   3,
		TotalPrice: 30,
		Item: OrderItem{
			ID:    "item1",
			Name:  "小三",
			Price: 100,
		},
	}
	b, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", b)
}

func unmarshal2() {
	res := `{
"data": [
    {
        "synonym":"",
        "weight":"0.6",
        "word": "真丝",
        "tag":"材质"
    },
    {
        "synonym":"",
        "weight":"0.8",
        "word": "韩都衣舍",
        "tag":"品牌"
    },
    {
        "synonym":"连身裙;联衣裙",
        "weight":"1.0",
        "word": "连衣裙",
        "tag":"品类"
    }
]
}`

	type my struct {
		Data []struct {
			Synonym string `json:"synonym"`
		} `json:"data"`
	}
	a := my{}

	err := json.Unmarshal([]byte(res), &a)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", a.Data[2].Synonym)
}

func main() {
	unmarshal2()
}
