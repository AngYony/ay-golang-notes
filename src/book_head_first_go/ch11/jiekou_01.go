package main

import (
	"book_head_first_go/ch11/gadget"
	"fmt"
)

type Player interface {
	Play(string)
	Stop()
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	var player Player
	player = gadget.TapePlayer{}
	mixtape := []string{"AAA", "BBB", "CCC"}
	playList(player, mixtape)

	player = gadget.TapeRecorder{}
	playList(player, mixtape)

	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
	} else {
		fmt.Println("断言失败")
	}

}
