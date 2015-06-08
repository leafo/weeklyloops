package physics

import "github.com/leafo/weeklyloops/loops"

type Particle3d struct {
	Pos         loops.Vec3
	Vel         loops.Vec3
	accel       loops.Vec3
	inverseMass float64

	forces []ForceGenerator
}

func NewParticle3d(mass, x, y, z float64) *Particle3d {
	forces := make([]ForceGenerator, 0)
	forces = append(forces, defaultGravity)

	return &Particle3d{
		Pos:         loops.Vec3{float32(x), float32(y), float32(z)},
		forces:      forces,
		inverseMass: 1 / mass,
	}
}

func (self *Particle3d) SetMass(mass float64) {
	self.inverseMass = 1 / mass
}

func (self *Particle3d) Mass() float64 {
	return 1 / self.inverseMass
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

	self.Pos = self.Pos.Add(self.Vel.Scale(float32(dt)))
	self.accel = loops.Vec3{}

	for _, fg := range self.forces {
		fg.Apply(self)
	}

	self.Vel = self.Vel.Add(self.accel.Scale(float32(dt)))
}
