package loops

import "math"

type Vec3 [3]float32

func NewVec3(x, y, z float32) Vec3 {
	return Vec3{x, y, z}
}

func (self Vec3) Length() float32 {
	pow2 := self[0]*self[0] + self[1]*self[1] + self[2]*self[2]

	if pow2 == 0 {
		return 0
	}

	return float32(math.Sqrt(float64(pow2)))
}

func (self Vec3) Normalize() Vec3 {
	length := self.Length()
	if length == 0 {
		return Vec3{}
	}

	return Vec3{
		self[0] / length,
		self[1] / length,
		self[2] / length,
	}
}

func (self Vec3) Add(other Vec3) Vec3 {
	return Vec3{
		self[0] + other[0],
		self[1] + other[1],
		self[2] + other[2],
	}
}

func (self Vec3) Scale(s float32) Vec3 {
	return Vec3{
		self[0] * s,
		self[1] * s,
		self[2] * s,
	}
}
