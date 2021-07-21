package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		_ = r.ParseForm()
		fmt.Println("path:", r.URL.Path)

		//获取a的值
		a, _ := strconv.Atoi(r.Form["a"][0])
		b, _ := strconv.Atoi(r.Form["b"][0])

		w.Header().Set("Content-Type", "application/json")
		//转换为二进制数据（字节数组)
		jData, _ := json.Marshal(map[string]int{
			"data": a + b,
		})

		_, _ = w.Write(jData)
	})

	_ = http.ListenAndServe(":8000", nil)
}
