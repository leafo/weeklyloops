package main

import (
	"image/color"
	"math"

	"github.com/leafo/weeklyloops/loops"
)

var background = color.RGBA{28, 30, 22, 255}

func main() {
	loop := loops.NewLoopWindow()

	loop.Title = "loop2"
	loop.Speed = 0.2

	circle := loops.NewCircle(1.0, 5)

	pal := loops.SmoothPalette{
		[]color.RGBA{
			color.RGBA{0, 0, 0, 255},
			color.RGBA{94, 92, 170, 255},
			color.RGBA{174, 111, 184, 255},
			color.RGBA{198, 137, 155, 255},
			color.RGBA{210, 187, 160, 255},
			color.RGBA{209, 224, 183, 255},
			color.RGBA{215, 241, 220, 255},
			color.RGBA{242, 250, 250, 255},
			color.RGBA{102, 120, 120, 255},
		},
	}

	loop.Draw = func(t float64, g *loops.Graphics) {
		g.SetColor(background)
		g.SetMat(loops.NewIdentityMat4())
		g.DrawRect(-1, -1, 2, 2)

		//waver := math.Sin(t * math.Pi)
		t = t + math.Cos(t*math.Pi*30)/100.0

		g.SetColor(color.RGBA{200, 200, 200, 255})

		layers := 12
		for i := 0; i < layers-1; i += 1 {
			phase := (math.Sin(t*math.Pi*2+float64(i))+1.0)/6.0 + 2.0/3.0
			scale := float64(layers-i)/float64(layers)/2.0 + 0.5
			rot := -t*math.Pi*2/5.0*float64(i+1) + float64(i)

			m := loops.NewIdentityMat4().
				Scale(float32(phase), float32(phase), 1).
				Rotate(float32(rot)).
				Scale(float32(scale), float32(scale), 1)

			g.SetMat(m)
			g.SetColor(pal.Color(float64(i) / float64(layers)))
			g.DrawShape(circle)
		}
	}

	loop.Run()
}
