package main

import "github.com/leafo/weeklyloops/loops"

func main() {
	loop := loops.NewLoopWindow()
	loop.Title = "loop5"
	loop.Speed = 0.2

	loop.Draw = func(t float64, g *loops.Graphics) {
	}

	loop.Run()
}
