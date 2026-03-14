package main

import (
	"fmt"
	"learning-golang/src/github.com/headfirstgo/gadget"
)

type Player interface{
	Play(string)
	Stop()
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func TryOut(player Player) {
	player.Play("Test track")
	player.Stop()
	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
	}
	
}

func main() {
	mixtape := []string{"Jessie's Girl", "Whip it", "9 to 5"}

	var player Player = gadget.TapeRecorder{}
	playList(player, mixtape)

	player = gadget.TapeRecorder{}
	playList(player, mixtape)
	fmt.Println("---------------------------")

	TryOut(gadget.TapeRecorder{})
	TryOut(gadget.TapePlayer{})

}


