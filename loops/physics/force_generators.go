package physics

import "github.com/leafo/weeklyloops/loops"

var defaultGravity = GravityGenerator{15}

type ForceGenerator interface {
	Apply(particle *Particle3d)
}

type ForceGeneratorFunc func(*Particle3d)

func (self ForceGeneratorFunc) Apply(p *Particle3d) {
	self(p)
}

type GravityGenerator struct {
	gravity float64
}

func (self GravityGenerator) Apply(p *Particle3d) {
	p.ApplyForce(loops.Vec3{0, -float32(self.gravity), 0}.Scale(float32(p.Mass())))
}

type DragGenerator struct {
	k1, k2 float64
}

func (self DragGenerator) Apply(p *Particle3d) {
	dir := p.Vel
	amount := float64(dir.Length())
	amount = self.k1*amount + self.k2*amount*amount
	force := dir.Normalize().Scale(float32(-amount))
	p.ApplyForce(force)
}
