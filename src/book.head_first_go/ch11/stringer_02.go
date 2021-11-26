package main

import "fmt"

type Gallons float64

func (g Gallons) String() string {
	return fmt.Sprintf("%0.2f gal", g)
}

type Liters float64

func (l Liters) String() string {
	return fmt.Sprintf("%0.2f L", l)
}

type Milliliters float64

func (m Milliliters) String() string {
	return fmt.Sprintf("%0.2f ml", m)
}

type wy float64

func main() {
	fmt.Println(Gallons(12.322))
	fmt.Println(Liters(12.3324324))
	fmt.Println(Milliliters(12.67454))
	fmt.Println(wy(1.2344))
}
