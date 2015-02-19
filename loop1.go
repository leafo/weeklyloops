package main

import "github.com/leafo/weeklyloops/loops"

func main() {
	loop := loops.NewLoopWindow()
	loop.Draw = func() {
	}

	loop.Run()
}
