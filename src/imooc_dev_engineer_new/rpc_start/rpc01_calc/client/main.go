package main

import (
	"encoding/json"
	"fmt"

	"github.com/kirinlabs/HttpRequest"
)

type ResponseData struct {
	Data int `json:"data"`
}

func main() {
	fmt.Println(Add(1, 2))

}

func Add(a, b int) int {
	req := HttpRequest.NewRequest()
	resp, _ := req.Get(fmt.Sprintf("http://localhost:8000/%s?a=%d&b=%d", "add", a, b))
	body, _ := resp.Body()
	//fmt.Println(string(body))

	rspData := ResponseData{}
	_ = json.Unmarshal(body, &rspData)

	return rspData.Data
}
