package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"json-to-struct/user"
)

func main() {
	var mdusers = make(map[string][]user.User)

	// 读取json文件内容
	jd, err := ioutil.ReadFile("json/users.json")
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(jd, &mdusers); err != nil {
		panic(err)
	}

	for _, us := range mdusers {
		for _, u := range us {
			fmt.Printf(
				"User Type: %v\nUser Age: %v\nUser Name: %v\nFacebook Url: %v \n",
				u.Type, u.Age, u.Name, u.Social.FaceBook,
			)
		}
	}

}
