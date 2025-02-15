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
	}
}

type Barrel struct {
	o.Sprite
}

func (b Barrel) GetSprite() o.Sprite {
	return b.Sprite
}

func (b Barrel) Close() {
	b.Sprite.Close()
}
