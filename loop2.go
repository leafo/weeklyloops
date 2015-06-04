package main

import (
	"image/color"
	"math"
)
import "github.com/leafo/weeklyloops/loops"

var background = color.RGBA{16, 50, 8, 255}

var colors = []color.RGBA{
	color.RGBA{94, 92, 170, 255},
	color.RGBA{174, 111, 184, 255},
	color.RGBA{198, 137, 155, 255},
	color.RGBA{210, 187, 160, 255},
	color.RGBA{209, 224, 183, 255},
	color.RGBA{215, 241, 220, 255},
	color.RGBA{242, 250, 250, 255},
}

func main() {
	loop := loops.NewLoopWindow()

	loop.Title = "loop3"
	loop.Speed = 0.2

	rect := loops.NewRectangle(1.2, 0.5)

	loop.Draw = func(t float64, g *loops.Graphics) {
		if t < 0.02 {
		}

		rawT := t
		t = (math.Sin(t*math.Pi*2) + 1.0) / 2.0
		t = loops.Smoothstep(t)

		g.SetColor(background)
		g.SetMat(loops.NewIdentityMat4())
		g.DrawRect(-1, -1, 2, 2)

		innerRot := rawT
		rotMul := 1.0
		if innerRot > 0.5 {
			innerRot -= 0.5
			rotMul = -1.0
		}

		realRot := 0.0
		// 0.1 - 0.3
		if innerRot >= 0.1 && innerRot <= 0.3 {
			realRot = loops.Smoothstep(((innerRot - 0.1) / 0.2)) * math.Pi * rotMul
		}

		numLayers := len(colors)
		layer := 0
		for _, color := range colors {
			g.SetColor(color)
			layerScale := (float64(layer)/(float64(numLayers)-1.0))/2.0 + 0.5

			row := 0
			for y := -1.2; y < 1.2; y += 0.2 {
				col := 0
				for x := -3.5; x < 3.5; x += 0.5 {
					dir := 1.0
					wide := (row+col)%2 == 0
					ss := 0.8
					if wide {
						ss = 1.3
						dir = -1.0
					}

					vibe := 1.0 + math.Cos(rawT*math.Pi*2.0+float64(layer)*2)*0.1

					m := loops.NewIdentityMat4().
						Rotate(float32(realRot)).
						Scale(float32(vibe), float32(vibe), 1).
						Scale(0.2, 0.2, 1).
						Scale(float32(ss), float32(ss), 1).
						Translate(float32(x+t*dir), float32(y), 0).
						Scale(float32(layerScale), float32(layerScale), 0).
						Rotate(float32(math.Cos(t*math.Pi*2.0) * 0.1 * dir))

					g.SetMat(m)
					g.DrawShape(rect)

					col += 1
				}

				row += 1
			}

			layer += 1
		}

	}

	loop.Run()
}
