package loops

import "fmt"

type Mat4 [16]float32

func NewIdentityMat4() Mat4 {
	return Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
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
