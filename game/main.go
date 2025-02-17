package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	p "github.com/danielherschel/raylib-test/game/prefabs"
	u "github.com/danielherschel/raylib-test/game/utils"
)

const (
	CROSSHAIR_SIZE = 64
)

func main() {
	// Initialize window
	rl.InitWindow(u.SCREEN_WIDTH, u.SCREEN_HEIGHT, "Raycaster")
	rl.SetTargetFPS(u.FRAME_RATE)
	rl.SetBlendMode(rl.BlendAlpha)
	rl.DisableCursor()

	level := p.NewLevel(u.LEVEL_1_PATH)
	defer level.Close()

	// HUD initialization
	crosshair := rl.LoadTexture("assets/sprites/crosshair/x.png")

	// Main loop
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		level.MainLoop()

		// HUD
		rl.DrawTexturePro(
			crosshair,
			rl.NewRectangle(0, 0, float32(crosshair.Width), float32(crosshair.Height)),
			rl.NewRectangle(float32(u.SCREEN_WIDTH/2-CROSSHAIR_SIZE/2), float32(u.SCREEN_HEIGHT/2-CROSSHAIR_SIZE/2), CROSSHAIR_SIZE, CROSSHAIR_SIZE),
			rl.NewVector2(0, 0),
			0,
			rl.White,
		)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
