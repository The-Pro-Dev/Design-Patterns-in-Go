package FlyweightPattern

import "fmt"

func displayBullet(bullet MovingParticle) {
	fmt.Println("The Bullet with a color of", bullet.getColor(), "is currently at", bullet.getPosition(), "going in direction", bullet.getVector(), "and moving at a speed of", bullet.getSpeed(), "m/s")
}

func Main() {
	fmt.Println("** Flyweight Pattern **")

	bullet1 := SilverBulletFactory()
	bullet1.move(CoOrdinates{0, 0, 0}, Vector{1, 1, 1}, 5.0)
	displayBullet((bullet1))

	bullet2 := SilverBulletFactory()
	bullet2.move(CoOrdinates{8, 9, 8}, Vector{1, -5, 6}, 9.0)
	displayBullet(bullet2)
}
