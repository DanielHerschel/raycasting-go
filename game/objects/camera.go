package objects

import (
	"math"

	u "github.com/danielherschel/raylib-test/game/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewCamera(transform Transform, fov float32) *Camera {
	camera := &Camera{Transform: transform, Plane: rl.NewVector2(0, 0), ZBuffer: make([]float32, u.SCREEN_WIDTH)}
	camera.SetFOV(fov)
	return camera
}

type Camera struct {
	Transform
	Plane   rl.Vector2
	ZBuffer []float32
}

// SetFOV sets the camera's field of view (FOV) in degrees.
func (c *Camera) SetFOV(fov float32) {
	// Calculate the new length of the Plane vector based on the desired FOV.
	// The FOV is the angle between the left and right edges of the view.
	// The Plane vector length is calculated as tan(FOV / 2).
	planeLength := float32(math.Tan(float64(fov) * math.Pi / 360.0))
	c.Plane = rl.NewVector2(c.Direction.Y*planeLength, -c.Direction.X*planeLength)
}
