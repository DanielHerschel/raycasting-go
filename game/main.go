package main

import (
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"

	o "github.com/danielherschel/raylib-test/game/objects"
	u "github.com/danielherschel/raylib-test/game/utils"
)

func getWorldMap() [][]int {
	return [][]int{
		{8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 4, 4, 6, 4, 4, 6, 4, 6, 4, 4, 4, 6, 4},
		{8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4},
		{8, 0, 3, 3, 0, 0, 0, 0, 0, 8, 8, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 6},
		{8, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 6},
		{8, 0, 3, 3, 0, 0, 0, 0, 0, 8, 8, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4},
		{8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8, 4, 0, 0, 0, 0, 0, 6, 6, 6, 0, 6, 4, 6},
		{8, 8, 8, 8, 0, 8, 8, 8, 8, 8, 8, 4, 4, 4, 4, 4, 4, 6, 0, 0, 0, 0, 0, 6},
		{7, 7, 7, 7, 0, 7, 7, 7, 7, 0, 8, 0, 8, 0, 8, 0, 8, 4, 0, 4, 0, 6, 0, 6},
		{7, 7, 0, 0, 0, 0, 0, 0, 7, 8, 0, 8, 0, 8, 0, 8, 8, 6, 0, 0, 0, 0, 0, 6},
		{7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8, 6, 0, 0, 0, 0, 0, 4},
		{7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8, 6, 0, 6, 0, 6, 0, 6},
		{7, 7, 0, 0, 0, 0, 0, 0, 7, 8, 0, 8, 0, 8, 0, 8, 8, 6, 4, 6, 0, 6, 6, 6},
		{7, 7, 7, 7, 0, 7, 7, 7, 7, 8, 8, 4, 0, 6, 8, 4, 8, 3, 3, 3, 0, 3, 3, 3},
		{2, 2, 2, 2, 0, 2, 2, 2, 2, 4, 6, 4, 0, 0, 6, 0, 6, 3, 0, 0, 0, 0, 0, 3},
		{2, 2, 0, 0, 0, 0, 0, 2, 2, 4, 0, 0, 0, 0, 0, 0, 4, 3, 0, 0, 0, 0, 0, 3},
		{2, 0, 0, 0, 0, 0, 0, 0, 2, 4, 0, 0, 0, 0, 0, 0, 4, 3, 0, 0, 0, 0, 0, 3},
		{1, 0, 0, 0, 0, 0, 0, 0, 1, 4, 4, 4, 4, 4, 6, 0, 6, 3, 3, 0, 0, 0, 3, 3},
		{2, 0, 0, 0, 0, 0, 0, 0, 2, 2, 2, 1, 2, 2, 2, 6, 6, 0, 0, 5, 0, 5, 0, 5},
		{2, 2, 0, 0, 0, 0, 0, 2, 2, 2, 0, 0, 0, 2, 2, 0, 5, 0, 5, 0, 0, 0, 5, 5},
		{2, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 2, 5, 0, 5, 0, 5, 0, 5, 0, 5},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5},
		{2, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 2, 5, 0, 5, 0, 5, 0, 5, 0, 5},
		{2, 2, 0, 0, 0, 0, 0, 2, 2, 2, 0, 0, 0, 2, 2, 0, 5, 0, 5, 0, 0, 0, 5, 5},
		{2, 2, 2, 2, 1, 2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5},
	}
}

func getWalls() (wallsImages []*rl.Image, wallsTextures []rl.Texture2D) {
	wallsImages = append(wallsImages, rl.LoadImage("assets/textures/banner.png"))
	wallsImages = append(wallsImages, rl.LoadImage("assets/textures/redbricks.png"))
	wallsImages = append(wallsImages, rl.LoadImage("assets/textures/purplemeat.png"))
	wallsImages = append(wallsImages, rl.LoadImage("assets/textures/stonebricks.png"))
	wallsImages = append(wallsImages, rl.LoadImage("assets/textures/bluebricks.png"))
	wallsImages = append(wallsImages, rl.LoadImage("assets/textures/mossbricks.png"))
	wallsImages = append(wallsImages, rl.LoadImage("assets/textures/wood.png"))
	wallsImages = append(wallsImages, rl.LoadImage("assets/textures/bricks.png"))

	for _, image := range wallsImages {
		wallsTextures = append(wallsTextures, rl.LoadTextureFromImage(image))
	}

	return
}

func main() {
	// Initialize window
	rl.InitWindow(u.SCREEN_WIDTH, u.SCREEN_HEIGHT, "Raycaster")
	rl.SetTargetFPS(u.FRAME_RATE)

	// Texture loading and initialization
	wallsImages, wallTextures := getWalls()

	floorTexture := rl.LoadImageColors(wallsImages[3])
	ceilingTexture := rl.LoadImageColors(wallsImages[6])
	u.UnloadImages(wallsImages)

	// Load map data
	worldMap := getWorldMap()

	walls := o.NewWalls(worldMap, wallTextures)
	defer walls.Close()

	floorCeiling := o.NewFloorCeiling(floorTexture, ceilingTexture)
	defer walls.Close()

	// Camera settings
	camera := o.NewCamera(
		o.NewTransform(rl.NewVector2(22.0, 12.0), rl.NewVector2(-1.0, 0.0)),
		rl.NewVector2(0.0, 0.66),
	)

	// Time and physics initialization
	currentTime, oldTime := time.Now().UnixMilli(), int64(0)

	// Main loop
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		floorCeiling.Draw(*camera)
		walls.Draw(*camera)

		// Timing for FPS counter
		oldTime = currentTime
		currentTime = time.Now().UnixMilli()
		frameTime := (float64(currentTime) - float64(oldTime)) / 1000.0
		rl.DrawText(fmt.Sprintf("FPS: %d", int(1.0/frameTime)), 10, 10, 30, rl.White)

		camera.Update(frameTime, worldMap)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
