package parser

import (
	"fmt"
	"googlecrawler/crawler/fetcher"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%s\n", contents)
	result := ParseCityList(contents)
	fmt.Printf("%d", len(result.Requests))

	if len(result.Requests) != 470 {
		t.Errorf("数量不同")
	}

}
