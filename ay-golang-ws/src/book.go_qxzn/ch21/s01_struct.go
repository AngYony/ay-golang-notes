package main

import "fmt"

type location struct {
	lat  float64
	long float64
}

func main() {
	var spirit location
	spirit.lat = -14.5
	spirit.long = 12.44

	fmt.Printf("%+v", spirit)
}
