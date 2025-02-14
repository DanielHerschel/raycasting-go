package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	p "github.com/danielherschel/raylib-test/game/prefabs"
	u "github.com/danielherschel/raylib-test/game/utils"
)

func main() {
	// Initialize window
	rl.InitWindow(u.SCREEN_WIDTH, u.SCREEN_HEIGHT, "Raycaster")
	rl.SetTargetFPS(u.FRAME_RATE)
	rl.SetBlendMode(rl.BlendAlpha)

	level := p.NewLevel()
	defer level.Close()

	// Main loop
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		level.MainLoop()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
