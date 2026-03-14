package main

import (
	"fmt"
)

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("Honk!")
}

type NoiseMaker interface {
	MakeSound()
}

func Play(n NoiseMaker) {
	n.MakeSound()
}

type Robot string

func (r Robot) MakeSound() {
	fmt.Println("Beep Boop!")
}

func (r Robot) Walk() {
	fmt.Println("Powering legs")
}

func main() {
	var toy NoiseMaker
	toy = Whistle("Toyco Canary")
	toy.MakeSound()
	toy = Horn("Toyco Blaster")
	toy.MakeSound()
	fmt.Println("########################")
	var noiseMaker NoiseMaker = Robot("Botco Ambler")
	noiseMaker.MakeSound()
	var robot Robot = noiseMaker.(Robot)
	robot.Walk() 
}
