package loops

import (
	"fmt"
	"math"
)

type Vec2 [2]float32

func degToRadians(deg float32) float32 {
	return deg / 180 * math.Pi
}

func NewVec2(x, y float32) Vec2 {
	return Vec2{x, y}
}

func NewVec2FromAngle(deg float32) Vec2 {
	rad := degToRadians(deg)

	return Vec2{
		float32(math.Cos(float64(rad))),
		float32(math.Sin(float64(rad))),
	}
}

func (self Vec2) Scale(s float32) Vec2 {
	return Vec2{
		self[0] * s,
		self[1] * s,
	}
}

func (self Vec2) Add(other Vec2) Vec2 {
	return Vec2{
		self[0] + other[0],
		self[1] + other[1],
	}
}

func (self Vec2) Sub(other Vec2) Vec2 {
	return Vec2{
		self[0] - other[0],
		self[1] - other[1],
	}
}

func (self Vec2) Length() float32 {
	pow2 := self[0]*self[0] + self[1]*self[1]
	if pow2 == 0 {
		return 0
	}

	return float32(math.Sqrt(float64(pow2)))
}

func (self Vec2) Normalize() Vec2 {
	length := self.Length()
	if length == 0 {
		return Vec2{}
	}

	return Vec2{
		self[0] / length,
		self[1] / length,
	}
}

func (self Vec2) Rotate(rads float32) Vec2 {
	x := float64(self[0])
	y := float64(self[1])

	c := math.Cos(float64(rads))
	s := math.Sin(float64(rads))

	return Vec2{
		float32(x*c - y*s),
		float32(y*c + x*s),
	}
}

func (self Vec2) RotateAngle(deg float32) Vec2 {
	return self.Rotate(degToRadians(deg))
}

func (self Vec2) Print() {
	fmt.Printf("Vec2<%.2f,%.2f>\n", self[0], self[1])
}
