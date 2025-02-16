package prefabs

import (
	o "github.com/danielherschel/raylib-test/game/objects"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewBarrel(x, y float32) Barrel {
	barrelTexture := rl.LoadTexture("assets/sprites/barrel.png")

	return Barrel{
		o.NewSprite(
			o.NewTransform(rl.NewVector2(x, y), rl.NewVector2(0.0, 0.0)),
			barrelTexture,
		),
		o.NewHitBox(x, y, 0.5),
	}
}

type Barrel struct {
	o.Sprite
	o.HitBox
}

func (b Barrel) GetSprite() o.Sprite {
	return b.Sprite
}

func (b Barrel) Close() {
	b.Sprite.Close()
}

func (b Barrel) GetHitBox() o.HitBox {
	return b.HitBox
}
