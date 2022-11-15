package FlyweightPattern

type SilverBullet struct {
}

func (bullet SilverBullet) getColor() string {
	return "Silver"
}

func (bullet SilverBullet) getSprite() string {
	return "https://upload.wikimedia.org/wikipedia/commons/6/64/Silver_bullet_%28rotated%29.png"
}

func SilverBulletFactory() MovingParticle {
	return MovingParticleFactory(SilverBullet{})
}
