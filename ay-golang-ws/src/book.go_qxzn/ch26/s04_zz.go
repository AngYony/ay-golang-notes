package main

import "fmt"

type stats struct {
	level int
	count int
}
type character struct {
	name  string
	stats stats
}

func levelUp(s *stats) {
	s.level++
	s.count = 5 * s.count
}
func main() {
	player := character{name: "张三"}
	levelUp(&player.stats)

	fmt.Println(player)
}
