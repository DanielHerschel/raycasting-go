package objects

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type IHittable interface {
	GetHitBox() HitBox
	OnHit()
}

type IDestroyable interface {
	ShouldDestroy() bool
}

func NewHitBox(x, y, size float32) HitBox {
	return HitBox{
		Transform: Transform{
			Position: rl.NewVector2(x, y),
		},
		Size: size,
	}
}

type HitBox struct {
	Transform
	Size float32
}

func (h HitBox) CheckCollision(otherRay Transform) bool {
	m := (otherRay.Direction.Y) / (otherRay.Direction.X)
	perpM := -1 / m

	halfSize := float64(h.Size / 2)
	dx := float32(halfSize / math.Sqrt(float64(1+perpM*perpM)))
	dy := perpM * dx

	point1 := rl.NewVector2(h.Position.X-dx, h.Position.Y-dy)
	point2 := rl.NewVector2(h.Position.X+dx, h.Position.Y+dy)

	return rl.CheckCollisionLines(
		point1,
		point2,
		rl.NewVector2(otherRay.Position.X, otherRay.Position.Y),
		rl.NewVector2(otherRay.Position.X+otherRay.Direction.X*1e30, otherRay.Position.Y+otherRay.Direction.Y*1e30),
		nil,
	)
}
