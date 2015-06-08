package loops

import (
	"fmt"
	"math"
)

type Mat4 [16]float32

func NewIdentityMat4() Mat4 {
	return Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

func NewTranslateMatrix(tx, ty, tz float32) Mat4 {
	return Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		tx, ty, tz, 1,
	}
}

func NewScaleMatrix(sx, sy, sz float32) Mat4 {
	return Mat4{
		sx, 0, 0, 0,
		0, sy, 0, 0,
		0, 0, sz, 0,
		0, 0, 0, 1,
	}
}

func NewRotate2DMatrix(theta float32) Mat4 {
	c := float32(math.Cos(float64(theta)))
	s := float32(math.Sin(float64(theta)))

	return Mat4{
		c, s, 0, 0,
		-s, c, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

func NewRotateMatrix(theta, x, y, z float32) Mat4 {
	dir := NewVec3(x, y, z).Normalize()
	c := float32(math.Cos(float64(theta)))
	s := float32(math.Sin(float64(theta)))

	x = dir[0]
	y = dir[1]
	z = dir[2]
	c1 := 1 - c

	return Mat4{
		x*x*c1 + c, x*y*c1 - z*s, x*z*c1 + y*s, 0,
		y*x*c1 + z*s, y*y*c1 + c, y*z*c1 - x*s, 0,
		x*z*c1 - y*s, y*z*c1 + x*s, z*z*c1 + c, 0,
		0, 0, 0, 1,
	}.Transpose()
}

// fov: field of view in degrees
// aspect: width/height
func NewPerspectiveMatrix(fovy, aspect, zNear, zFar float64) Mat4 {
	fovy = fovy / 180.0 * math.Pi // deg to rad
	f := math.Tan(fovy / 2.0)
	nearFar := zNear - zFar

	return Mat4{
		float32(f / aspect), 0, 0, 0,
		0, float32(f), 0, 0,
		0, 0, float32((zFar + zNear) / nearFar), float32(2 * zFar * zNear / nearFar),
		0, 0, -1, 0,
	}.Transpose()
}

func (self Mat4) Mul(other Mat4) Mat4 {
	return Mat4{
		self[0]*other[0] + self[1]*other[4] + self[2]*other[8] + self[3]*other[12],
		self[0]*other[1] + self[1]*other[5] + self[2]*other[9] + self[3]*other[13],
		self[0]*other[2] + self[1]*other[6] + self[2]*other[10] + self[3]*other[14],
		self[0]*other[3] + self[1]*other[7] + self[2]*other[11] + self[3]*other[15],
		self[4]*other[0] + self[5]*other[4] + self[6]*other[8] + self[7]*other[12],
		self[4]*other[1] + self[5]*other[5] + self[6]*other[9] + self[7]*other[13],
		self[4]*other[2] + self[5]*other[6] + self[6]*other[10] + self[7]*other[14],
		self[4]*other[3] + self[5]*other[7] + self[6]*other[11] + self[7]*other[15],
		self[8]*other[0] + self[9]*other[4] + self[10]*other[8] + self[11]*other[12],
		self[8]*other[1] + self[9]*other[5] + self[10]*other[9] + self[11]*other[13],
		self[8]*other[2] + self[9]*other[6] + self[10]*other[10] + self[11]*other[14],
		self[8]*other[3] + self[9]*other[7] + self[10]*other[11] + self[11]*other[15],
		self[12]*other[0] + self[13]*other[4] + self[14]*other[8] + self[15]*other[12],
		self[12]*other[1] + self[13]*other[5] + self[14]*other[9] + self[15]*other[13],
		self[12]*other[2] + self[13]*other[6] + self[14]*other[10] + self[15]*other[14],
		self[12]*other[3] + self[13]*other[7] + self[14]*other[11] + self[15]*other[15],
	}
}

func (self Mat4) Print() {
	for i, val := range self {
		if i != 0 && i%4 == 0 {
			fmt.Print("\n")
		}

		fmt.Printf("%.2f, ", val)
	}

	fmt.Print("\n")
}

func (self Mat4) Translate(tx, ty, tz float32) Mat4 {
	return self.Mul(NewTranslateMatrix(tx, ty, tz))
}

func (self Mat4) TranslateVec3(v Vec3) Mat4 {
	return self.Mul(NewTranslateMatrix(v[0], v[1], v[2]))
}

func (self Mat4) Scale(sx, sy, sz float32) Mat4 {
	return self.Mul(NewScaleMatrix(sx, sy, sz))
}

func (self Mat4) Rotate(theta float32) Mat4 {
	return self.Mul(NewRotate2DMatrix(theta))
}

func (self Mat4) Rotate3d(theta, x, y, z float32) Mat4 {
	return self.Mul(NewRotateMatrix(theta, x, y, z))
}

func (self Mat4) Transpose() Mat4 {
	return Mat4{
		self[0], self[4], self[8], self[12],
		self[1], self[5], self[9], self[13],
		self[2], self[6], self[10], self[14],
		self[3], self[7], self[11], self[15],
	}
}
