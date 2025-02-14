package prefabs

import (
	o "github.com/danielherschel/raylib-test/game/objects"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewPillar(x, y float32) Pillar {
	pillarTexture := rl.LoadTexture("assets/sprites/pillar.png")

	return Pillar{
		o.NewSprite(
			o.NewTransform(rl.NewVector2(x, y), rl.NewVector2(0.0, 0.0)), 
			pillarTexture,
		),
	}
}

type Pillar struct {
	o.Sprite
}
