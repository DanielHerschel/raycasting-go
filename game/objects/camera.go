package objects

import (
	"math"

	u "github.com/danielherschel/raylib-test/game/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewCamera(transform Transform, plane rl.Vector2) *Camera {
	return &Camera{Transform: transform, Plane: plane, ZBuffer: make([]float32, u.SCREEN_WIDTH)}
}

type Camera struct {
	Transform
	Plane   rl.Vector2
	ZBuffer []float32
}

func (c *Camera) Update(frameTime float64, worldMap [][]int) {
	moveSpeed := float32(frameTime * 2.0)
	if rl.IsKeyDown(rl.KeyLeftShift) {
		moveSpeed *= 1.6
	}
	rotSpeed := rl.GetMouseDelta().X * float32(frameTime*0.2)

	if rl.IsKeyDown(rl.KeyW) {
		if worldMap[int(c.Position.X+c.Direction.X*moveSpeed)][int(c.Position.Y)] == 0 {
			c.Position.X += c.Direction.X * moveSpeed
		}
		if worldMap[int(c.Position.X)][int(c.Position.Y+c.Direction.Y*moveSpeed)] == 0 {
			c.Position.Y += c.Direction.Y * moveSpeed
		}
	}
	if rl.IsKeyDown(rl.KeyS) {
		if worldMap[int(c.Position.X-c.Direction.X*moveSpeed)][int(c.Position.Y)] == 0 {
			c.Position.X -= c.Direction.X * moveSpeed
		}
		if worldMap[int(c.Position.X)][int(c.Position.Y-c.Direction.Y*moveSpeed)] == 0 {
			c.Position.Y -= c.Direction.Y * moveSpeed
		}
	}
	if rl.IsKeyDown(rl.KeyD) {
		walkDir := rl.NewVector2(-c.Direction.Y, c.Direction.X)
		if worldMap[int(c.Position.X-walkDir.X*moveSpeed)][int(c.Position.Y)] == 0 {
			c.Position.X -= walkDir.X * moveSpeed * 0.6
		}
		if worldMap[int(c.Position.X)][int(c.Position.Y+walkDir.Y*moveSpeed)] == 0 {
			c.Position.Y -= walkDir.Y * moveSpeed * 0.6
		}
	}
	if rl.IsKeyDown(rl.KeyA) {
		walkDir := rl.NewVector2(c.Direction.Y, -c.Direction.X)
		if worldMap[int(c.Position.X-walkDir.X*moveSpeed)][int(c.Position.Y)] == 0 {
			c.Position.X -= walkDir.X * moveSpeed * 0.6
		}
		if worldMap[int(c.Position.X)][int(c.Position.Y+walkDir.Y*moveSpeed)] == 0 {
			c.Position.Y -= walkDir.Y * moveSpeed * 0.6
		}
	}
	if rotSpeed != 0 {
		oldDirX := c.Direction.X
		c.Direction.X = c.Direction.X*float32(math.Cos(float64(-rotSpeed))) - c.Direction.Y*float32(math.Sin(float64(-rotSpeed)))
		c.Direction.Y = oldDirX*float32(math.Sin(float64(-rotSpeed))) + c.Direction.Y*float32(math.Cos(float64(-rotSpeed)))
		oldPlaneX := c.Plane.X
		c.Plane.X = c.Plane.X*float32(math.Cos(float64(-rotSpeed))) - c.Plane.Y*float32(math.Sin(float64(-rotSpeed)))
		c.Plane.Y = oldPlaneX*float32(math.Sin(float64(-rotSpeed))) + c.Plane.Y*float32(math.Cos(float64(-rotSpeed)))
	}
}
