package objects

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewCamera(transform Transform, plane rl.Vector2) *Camera {
	return &Camera{Transform: transform, Plane: plane}
}

type Camera struct {
	Transform
	Plane     rl.Vector2
}

func (c *Camera) Update(frameTime float64, worldMap [][]int) {
	moveSpeed := float32(frameTime * 3.0)
	rotSpeed := float32(frameTime * 3.0)

	if rl.IsKeyDown(rl.KeyUp) {
		if worldMap[int(c.Position.X+c.Direction.X*moveSpeed)][int(c.Position.Y)] == 0 {
			c.Position.X += c.Direction.X * moveSpeed
		}
		if worldMap[int(c.Position.X)][int(c.Position.Y+c.Direction.Y*moveSpeed)] == 0 {
			c.Position.Y += c.Direction.Y * moveSpeed
		}
	}
	if rl.IsKeyDown(rl.KeyDown) {
		if worldMap[int(c.Position.X-c.Direction.X*moveSpeed)][int(c.Position.Y)] == 0 {
			c.Position.X -= c.Direction.X * moveSpeed
		}
		if worldMap[int(c.Position.X)][int(c.Position.Y-c.Direction.Y*moveSpeed)] == 0 {
			c.Position.Y -= c.Direction.Y * moveSpeed
		}
	}
	if rl.IsKeyDown(rl.KeyRight) {
		oldDirX := c.Direction.X
		c.Direction.X = c.Direction.X*float32(math.Cos(float64(-rotSpeed))) - c.Direction.Y*float32(math.Sin(float64(-rotSpeed)))
		c.Direction.Y = oldDirX*float32(math.Sin(float64(-rotSpeed))) + c.Direction.Y*float32(math.Cos(float64(-rotSpeed)))
		oldPlaneX := c.Plane.X
		c.Plane.X = c.Plane.X*float32(math.Cos(float64(-rotSpeed))) - c.Plane.Y*float32(math.Sin(float64(-rotSpeed)))
		c.Plane.Y = oldPlaneX*float32(math.Sin(float64(-rotSpeed))) + c.Plane.Y*float32(math.Cos(float64(-rotSpeed)))
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		oldDirX := c.Direction.X
		c.Transform.Direction.X = c.Direction.X*float32(math.Cos(float64(rotSpeed))) - c.Direction.Y*float32(math.Sin(float64(rotSpeed)))
		c.Transform.Direction.Y = oldDirX*float32(math.Sin(float64(rotSpeed))) + c.Direction.Y*float32(math.Cos(float64(rotSpeed)))
		oldPlaneX := c.Plane.X
		c.Plane.X = c.Plane.X*float32(math.Cos(float64(rotSpeed))) - c.Plane.Y*float32(math.Sin(float64(rotSpeed)))
		c.Plane.Y = oldPlaneX*float32(math.Sin(float64(rotSpeed))) + c.Plane.Y*float32(math.Cos(float64(rotSpeed)))
	}
}
