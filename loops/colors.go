package loops

import "image/color"

type SmoothPalette struct {
	Colors []color.RGBA
}

func (self *SmoothPalette) Color(t float64) color.RGBA {
	chunk := 1.0 / float64(len(self.Colors))

	idx := int(t / chunk)
	if idx >= len(self.Colors)-1 {
		return self.Colors[len(self.Colors)-1]
	}
	return mixColors(self.Colors[idx], self.Colors[idx+1], t-float64(idx)*chunk)
}

func mixColors(a, b color.RGBA, amount float64) color.RGBA {
	bAmount := 1.0 - amount
	return color.RGBA{
		uint8(float64(a.R)*amount + float64(b.R)*bAmount),
		uint8(float64(a.G)*amount + float64(b.G)*bAmount),
		uint8(float64(a.B)*amount + float64(b.B)*bAmount),
		uint8(float64(a.A)*amount + float64(b.A)*bAmount),
	}
}
