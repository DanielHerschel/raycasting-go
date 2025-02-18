package prefabs

import (
	"math"

	o "github.com/danielherschel/raylib-test/game/objects"
	u "github.com/danielherschel/raylib-test/game/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewPlayer(transform o.Transform) *Player {
	return &Player{
		transform,
		o.NewCamera(transform, u.CAMERA_FOV),
	}
}

type Player struct {
	o.Transform
	Camera *o.Camera
}

func (p *Player) Update(frameTime float64, worldMap [][]int) {
	p.HandleWalking(frameTime, worldMap)
	p.HanldeCameraRotation(float32(frameTime))

	// Sync the camera transform with the player transform
	p.Camera.Transform = p.Transform
}

func (p *Player) HandleWalking(frameTime float64, worldMap [][]int) {
	moveSpeed := float32(frameTime * 2.0)
	if rl.IsKeyDown(rl.KeyLeftShift) {
		moveSpeed *= 1.6
	}

	if rl.IsKeyDown(rl.KeyW) {
		if worldMap[int(p.Position.X+p.Direction.X*moveSpeed)][int(p.Position.Y)] == 0 {
			p.Position.X += p.Direction.X * moveSpeed
		}
		if worldMap[int(p.Position.X)][int(p.Position.Y+p.Direction.Y*moveSpeed)] == 0 {
			p.Position.Y += p.Direction.Y * moveSpeed
		}
	}
	if rl.IsKeyDown(rl.KeyS) {
		if worldMap[int(p.Position.X-p.Direction.X*moveSpeed)][int(p.Position.Y)] == 0 {
			p.Position.X -= p.Direction.X * moveSpeed
		}
		if worldMap[int(p.Position.X)][int(p.Position.Y-p.Direction.Y*moveSpeed)] == 0 {
			p.Position.Y -= p.Direction.Y * moveSpeed
		}
	}
	if rl.IsKeyDown(rl.KeyD) {
		walkDir := rl.NewVector2(-p.Direction.Y, p.Direction.X)
		if worldMap[int(p.Position.X-walkDir.X*moveSpeed)][int(p.Position.Y)] == 0 {
			p.Position.X -= walkDir.X * moveSpeed * 0.6
		}
		if worldMap[int(p.Position.X)][int(p.Position.Y-walkDir.Y*moveSpeed)] == 0 {
			p.Position.Y -= walkDir.Y * moveSpeed * 0.6
		}
	}
	if rl.IsKeyDown(rl.KeyA) {
		walkDir := rl.NewVector2(p.Direction.Y, -p.Direction.X)
		if worldMap[int(p.Position.X-walkDir.X*moveSpeed)][int(p.Position.Y)] == 0 {
			p.Position.X -= walkDir.X * moveSpeed * 0.6
		}
		if worldMap[int(p.Position.X)][int(p.Position.Y-walkDir.Y*moveSpeed)] == 0 {
			p.Position.Y -= walkDir.Y * moveSpeed * 0.6
		}
	}
}

func (p *Player) HanldeCameraRotation(frameTime float32) {
	rotSpeed := rl.GetMouseDelta().X * float32(frameTime*u.CAMERA_SPEED)
	if rotSpeed != 0 {
		oldDirX := p.Direction.X
		p.Direction.X = p.Direction.X*float32(math.Cos(float64(-rotSpeed))) - p.Direction.Y*float32(math.Sin(float64(-rotSpeed)))
		p.Direction.Y = oldDirX*float32(math.Sin(float64(-rotSpeed))) + p.Direction.Y*float32(math.Cos(float64(-rotSpeed)))
		oldPlaneX := p.Camera.Plane.X
		p.Camera.Plane.X = p.Camera.Plane.X*float32(math.Cos(float64(-rotSpeed))) - p.Camera.Plane.Y*float32(math.Sin(float64(-rotSpeed)))
		p.Camera.Plane.Y = oldPlaneX*float32(math.Sin(float64(-rotSpeed))) + p.Camera.Plane.Y*float32(math.Cos(float64(-rotSpeed)))
	}
}

// IGameObject functions
func (p Player) GetTransform() o.Transform {
	return p.Transform
}

func (p Player) Close() {
}
