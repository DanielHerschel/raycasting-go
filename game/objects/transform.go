package objects

import rl "github.com/gen2brain/raylib-go/raylib"

func NewTransform(position rl.Vector2, direction rl.Vector2) Transform {
	return Transform{Position: position, Direction: direction}
}

type Transform struct {
	Position rl.Vector2
	Direction rl.Vector2 // Also rotation
}
