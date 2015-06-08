package physics

import "github.com/leafo/weeklyloops/loops"

type ForceGenerator interface {
	Apply(particle *Particle3d)
}

type Particle3d struct {
	Pos         loops.Vec3
	vel         loops.Vec3
	accel       loops.Vec3
	inverseMass float64

	forces []ForceGenerator
}

func NewParticle3d(mass, x, y, z float64) *Particle3d {
	return &Particle3d{
		Pos:         loops.Vec3{float32(x), float32(y), float32(z)},
		forces:      make([]ForceGenerator, 0),
		inverseMass: 1 / mass,
	}
}

func (self *Particle3d) SetMass(mass float64) {
	self.inverseMass = 1 / mass
}

func (self *Particle3d) ApplyForce(force loops.Vec3) {
	if self.inverseMass == 0 {
		return
	}

	self.accel = self.accel.Add(force.Scale(float32(self.inverseMass)))
}

func (self *Particle3d) Update(dt float64) {
	if self.inverseMass == 0 {
		return
	}

	self.Pos = self.Pos.Add(self.vel.Scale(float32(dt)))
	self.accel = loops.Vec3{}

	for _, fg := range self.forces {
		fg.Apply(self)
	}

	self.vel = self.vel.Add(self.accel.Scale(float32(dt)))
}
