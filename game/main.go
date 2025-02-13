package main

import (
	"fmt"
	"image/color"
	"math"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"

	o "github.com/danielherschel/raylib-test/game/objects"
)

const (
	SCREEN_WIDTH   = 1920
	SCREEN_HEIGHT  = 1080
	FRAME_RATE     = 60
	TEXTURE_WIDTH  = 64
	TEXTURE_HEIGHT = 64
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

func unloadTextures(textures []rl.Texture2D) {
	for _, texture := range textures {
		rl.UnloadTexture(texture)
	}
}

func unloadImages(images []*rl.Image) {
	for _, image := range images {
		rl.UnloadImage(image)
	}
}

func castWalls(camera o.Camera, worldMap [][]int, walls []rl.Texture2D) {
	position := camera.Position
	dir := camera.Direction
	plane := camera.Plane
	
	for x := 0; x < SCREEN_WIDTH; x++ {
		cameraX := 2*float32(x)/float32(SCREEN_WIDTH) - 1
		rayDir := rl.NewVector2(dir.X+plane.X*cameraX, dir.Y+plane.Y*cameraX)

		mapX, mapY := int(position.X), int(position.Y)
		var sideDist rl.Vector2

		var deltaDist rl.Vector2
		var prepWallDist float32

		if rayDir.X == 0 {
			deltaDist.X = 1e30
		} else {
			deltaDist.X = float32(math.Abs(float64(1.0 / rayDir.X)))
		}
		if rayDir.Y == 0 {
			deltaDist.Y = 1e30
		} else {
			deltaDist.Y = float32(math.Abs(float64(1.0 / rayDir.Y)))
		}

		var stepX int
		var stepY int

		hit := 0
		var side int

		// Set step direction and set the distances to the next closest square
		if rayDir.X < 0 {
			stepX = -1
			sideDist.X = (position.X - float32(mapX)) * deltaDist.X
		} else {
			stepX = 1
			sideDist.X = (float32(mapX+1) - position.X) * deltaDist.X
		}
		if rayDir.Y < 0 {
			stepY = -1
			sideDist.Y = (position.Y - float32(mapY)) * deltaDist.Y
		} else {
			stepY = 1
			sideDist.Y = (float32(mapY+1) - position.Y) * deltaDist.Y
		}

		for hit == 0 {
			// Jump to the next square in the X direction or Y direction
			if sideDist.X < sideDist.Y {
				sideDist.X += deltaDist.X
				mapX += stepX
				side = 0
			} else {
				sideDist.Y += deltaDist.Y
				mapY += stepY
				side = 1
			}

			// Check if ray hits a wall
			if worldMap[mapX][mapY] > 0 {
				hit = 1
			}
		}

		var wallX float32

		if side == 0 {
			prepWallDist = sideDist.X - deltaDist.X
			wallX = position.Y + prepWallDist*rayDir.Y
		} else {
			prepWallDist = sideDist.Y - deltaDist.Y
			wallX = position.X + prepWallDist*rayDir.X
		}
		wallX -= float32(math.Floor(float64(wallX)))

		lineHeight := int(SCREEN_HEIGHT / prepWallDist)

		drawStart := int(-float64(lineHeight)/2 + SCREEN_HEIGHT/2)
		// if drawStart < 0 {
		// 	drawStart = 0
		// }
		drawEnd := int(float64(lineHeight)/2 + SCREEN_HEIGHT/2)
		// if drawEnd >= SCREEN_HEIGHT {
		// 	drawEnd = SCREEN_HEIGHT
		// }

		wallType := worldMap[mapX][mapY] - 1
		texturePos := float32(wallX * TEXTURE_WIDTH)

		rl.DrawTexturePro(
			walls[wallType],
			rl.NewRectangle(texturePos, 0, 1, TEXTURE_HEIGHT),
			rl.NewRectangle(float32(x), float32(drawStart), 1, float32(drawEnd-drawStart)),
			rl.NewVector2(0.0, 0.0),
			0.0,
			rl.Gray,
		)
	}
}

func castCeiling(camera o.Camera, floorTexture []color.RGBA, ceilingTexture []color.RGBA, pixels []color.RGBA, floorCeilTexture rl.Texture2D) {
	position := camera.Position
	dir := camera.Direction
	plane := camera.Plane
	
	for y := 0; y < SCREEN_HEIGHT; y++ {
		rayDir0 := rl.NewVector2(dir.X-plane.X, dir.Y-plane.Y)
		rayDir1 := rl.NewVector2(dir.X+plane.X, dir.Y+plane.Y)

		p := float32(y) - SCREEN_HEIGHT/2

		posZ := float32(0.5 * SCREEN_HEIGHT)

		rowDistance := posZ / p

		floorStep := rl.NewVector2(
			rowDistance*(rayDir1.X-rayDir0.X)/SCREEN_WIDTH,
			rowDistance*(rayDir1.Y-rayDir0.Y)/SCREEN_WIDTH,
		)

		floor := rl.NewVector2(position.X+rowDistance*rayDir0.X, position.Y+rowDistance*rayDir0.Y)

		for x := 0; x < SCREEN_WIDTH; x++ {
			cellX := int(floor.X)
			cellY := int(floor.Y)

			tx := int32(TEXTURE_WIDTH*(floor.X-float32(cellX))) & (TEXTURE_WIDTH - 1)
			ty := int32(TEXTURE_HEIGHT*(floor.Y-float32(cellY))) & (TEXTURE_HEIGHT - 1)

			floor.X += floorStep.X
			floor.Y += floorStep.Y

			color := floorTexture[TEXTURE_WIDTH*ty+tx]
			pixels[SCREEN_WIDTH*y+x] = color

			color = ceilingTexture[TEXTURE_WIDTH*ty+tx]
			pixels[SCREEN_WIDTH*(SCREEN_HEIGHT-y-1)+x] = color
		}
	}

	rl.UpdateTexture(floorCeilTexture, pixels)
	rl.DrawTexture(floorCeilTexture, 0, 0, rl.Gray)
}

func main() {
	// Initialize window
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "Raycaster")
	rl.SetTargetFPS(FRAME_RATE)

	// Load map data
	worldMap := getWorldMap()

	// Texture loading and initialization
	wallsImages, walls := getWalls()
	floorCeilImage := rl.GenImageColor(SCREEN_WIDTH, SCREEN_HEIGHT, rl.White)
	floorCeilTexture := rl.LoadTextureFromImage(floorCeilImage)

	pixels := make([]color.RGBA, SCREEN_WIDTH*SCREEN_HEIGHT)

	floorTexture := rl.LoadImageColors(wallsImages[3])
	ceilingTexture := rl.LoadImageColors(wallsImages[6])

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

		castCeiling(*camera, floorTexture, ceilingTexture, pixels, floorCeilTexture)
		castWalls(*camera, worldMap, walls)

		// Timing for FPS counter
		oldTime = currentTime
		currentTime = time.Now().UnixMilli()
		frameTime := (float64(currentTime) - float64(oldTime)) / 1000.0
		rl.DrawText(fmt.Sprintf("FPS: %d", int(1.0/frameTime)), 10, 10, 30, rl.White)

		camera.Update(frameTime, worldMap)

		rl.EndDrawing()
	}

	rl.UnloadImageColors(floorTexture)
	rl.UnloadImageColors(ceilingTexture)
	rl.UnloadImage(floorCeilImage)
	rl.UnloadTexture(floorCeilTexture)

	unloadImages(wallsImages)
	unloadTextures(walls)

	rl.CloseWindow()
}
