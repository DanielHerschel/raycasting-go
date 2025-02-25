package objects

import (
	// "math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type IHittable interface {
	GetHitBox() HitBox
}

type IDestroyable interface {
	ShouldDestroy() bool
}

func NewHitBox(transform Transform, size float32) HitBox {
	boundingBox := rl.NewBoundingBox(
		rl.NewVector3(transform.Position.X-size/2, transform.Position.Y-size/2, 0),
		rl.NewVector3(transform.Position.X+size/2, transform.Position.Y+size/2, 0),
	)

	return HitBox{
		Transform:   transform,
		Size:        size,
		BoundingBox: boundingBox,
	}
}

type HitBox struct {
	Transform
	Size float32
	rl.BoundingBox
}
