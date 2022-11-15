package FlyweightPattern

type CoOrdinates struct {
	x float64
	y float64
	z float64
}

type Vector struct {
	i float64
	j float64
	k float64
}

type MovingParticle struct {
	particle *IParticle
	coOrds   CoOrdinates
	vector   Vector
	speed    float64
}

func MovingParticleFactory(particle IParticle) MovingParticle {
	return MovingParticle{particle: &particle}
}

func (particle *MovingParticle) move(coOrds CoOrdinates, vector Vector, speed float64) {
	particle.coOrds = coOrds
	particle.vector = vector
	particle.speed = speed
}

func (particle *MovingParticle) getPosition() CoOrdinates {
	return particle.coOrds
}

func (particle *MovingParticle) getSpeed() float64 {
	return particle.speed
}

func (particle *MovingParticle) getVector() Vector {
	return particle.vector
}

func (particle *MovingParticle) getColor() string {
	return (*particle.particle).getColor()
}

func (particle *MovingParticle) getSprite() string {
	return (*particle.particle).getSprite()
}
