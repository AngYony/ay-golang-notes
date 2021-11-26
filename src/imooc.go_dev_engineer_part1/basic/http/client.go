package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	var url string = "https://www.imooc.com"
	request, err := http.NewRequest(http.MethodGet, url, nil)

	request.Header.Add("User-Agent", "")

	client := http.Client{

		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("redirect:", request)
			return nil

		},
	}
	//resp, err := http.DefaultClient.Do(request)
	resp, err := client.Do(request)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", s)
}
