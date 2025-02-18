package prefabs

import (
	o "github.com/danielherschel/raylib-test/game/objects"
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
		2,
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
func (b *Barrel) OnHit() {
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		b.Health--
		if b.Health <= 0 {
			b.ShouldDest = true
		}
	}
}

func (b *Barrel) GetHitBox() o.HitBox {
	return b.HitBox
}

// IDestroyable functions
func (b Barrel) ShouldDestroy() bool {
	return b.ShouldDest
}
