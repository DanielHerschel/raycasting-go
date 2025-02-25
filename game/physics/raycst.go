package physics

import (
	o "github.com/danielherschel/raylib-test/game/objects"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func CastRay(origin o.IHittable, direction rl.Vector2, hittables []o.IHittable) o.IHittable {

	originPosition := origin.GetHitBox().Position

	ray := rl.NewRay(rl.NewVector3(originPosition.X, originPosition.Y, 0), rl.NewVector3(direction.X, direction.Y, 0))

	var closestHittable o.IHittable
	closestHittableDistance := float32(1e30)

	for _, hittable := range hittables {
		collision := rl.GetRayCollisionBox(ray, hittable.GetHitBox().BoundingBox)
		if collision.Hit && collision.Distance < closestHittableDistance {
			closestHittableDistance = collision.Distance
			closestHittable = hittable
		}
	}

	return closestHittable
}
