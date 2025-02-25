package prefabs

import (
	o "github.com/danielherschel/raylib-test/game/objects"
	u "github.com/danielherschel/raylib-test/game/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewBarrel(x, y float32) *Barrel {
	barrelTexture := rl.LoadTexture("assets/sprites/barrel.png")

	position := o.NewTransform(rl.NewVector2(x, y), rl.NewVector2(0.0, 0.0))

	return &Barrel{
		position,
		o.NewSprite(
			position,
			barrelTexture,
		),
		o.NewHitBox(position, 0.5),
		u.BARREL_HEALTH,
		false,
	}
}

type Barrel struct {
	o.Transform
	o.Sprite
	o.HitBox

	Health     int
	ShouldDest bool
}

// IGameObject functions
func (b *Barrel) Update(frameTime float64, currentLevel Level) {
	if !b.IsAlive() {
		b.ShouldDest = true
	}
}

func (b Barrel) GetTransform() o.Transform {
	return b.Transform
}

func (b Barrel) Close() {
	b.Sprite.Close()
}

// ISprite functions
func (b Barrel) GetSprite() o.Sprite {
	return b.Sprite
}

// IHittable functions
func (b *Barrel) GetHitBox() o.HitBox {
	return b.HitBox
}

// IDestroyable functions
func (b Barrel) ShouldDestroy() bool {
	return b.ShouldDest
}

// IDamageable functions

func (b *Barrel) TakeDamage(amount int) {
	b.Health -= amount
}

func (b Barrel) IsAlive() bool {
	return b.Health > 0
}
