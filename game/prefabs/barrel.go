package prefabs

import (
	o "github.com/danielherschel/raylib-test/game/objects"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewBarrel(x, y float32) *Barrel {
	barrelTexture := rl.LoadTexture("assets/sprites/barrel.png")

	return &Barrel{
		o.NewSprite(
			o.NewTransform(rl.NewVector2(x, y), rl.NewVector2(0.0, 0.0)),
			barrelTexture,
		),
		o.NewHitBox(x, y, 0.5),
		2,
		false,
	}
}

type Barrel struct {
	o.Sprite
	o.HitBox

	Health     int
	ShouldDest bool
}

// GameObject functions
func (b Barrel) GetSprite() o.Sprite {
	return b.Sprite
}

func (b Barrel) Close() {
	b.Sprite.Close()
}

// IHittable functions
func (b *Barrel) OnHit() {
	if rl.IsKeyPressed(rl.KeySpace) {
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
