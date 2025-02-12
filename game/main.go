package main

import (
	"fmt"
	"image/color"
	"math"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
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

func castWalls(dirX float64, dirY float64, planeX float64, planeY float64, posX float64, posY float64, worldMap [][]int, walls []rl.Texture2D) {
	for x := 0; x < SCREEN_WIDTH; x++ {
		cameraX := 2*float64(x)/float64(SCREEN_WIDTH) - 1
		rayDirX := dirX + planeX*cameraX
		rayDirY := dirY + planeY*cameraX

		mapX, mapY := int(posX), int(posY)
		var sideDistX float64
		var sideDistY float64

		var deltaDistX float64
		var deltaDistY float64
		var prepWallDist float64

		if rayDirX == 0 {
			deltaDistX = 1e30
		} else {
			deltaDistX = math.Abs(1.0 / rayDirX)
		}
		if rayDirY == 0 {
			deltaDistY = 1e30
		} else {
			deltaDistY = math.Abs(1.0 / rayDirY)
		}

		var stepX int
		var stepY int

		hit := 0
		var side int

		// Set step direction and set the distances to the next closest square
		if rayDirX < 0 {
			stepX = -1
			sideDistX = (posX - float64(mapX)) * deltaDistX
		} else {
			stepX = 1
			sideDistX = (float64(mapX+1) - posX) * deltaDistX
		}
		if rayDirY < 0 {
			stepY = -1
			sideDistY = (posY - float64(mapY)) * deltaDistY
		} else {
			stepY = 1
			sideDistY = (float64(mapY+1) - posY) * deltaDistY
		}

		for hit == 0 {
			// Jump to the next square in the X direction or Y direction
			if sideDistX < sideDistY {
				sideDistX += deltaDistX
				mapX += stepX
				side = 0
			} else {
				sideDistY += deltaDistY
				mapY += stepY
				side = 1
			}

			// Check if ray hits a wall
			if worldMap[mapX][mapY] > 0 {
				hit = 1
			}
		}

		var wallX float64

		if side == 0 {
			prepWallDist = sideDistX - deltaDistX
			wallX = posY + prepWallDist*rayDirY
		} else {
			prepWallDist = sideDistY - deltaDistY
			wallX = posX + prepWallDist*rayDirX
		}
		wallX -= math.Floor(wallX)

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

func castCeiling(dirX float64, dirY float64, planeX float64, planeY float64, posX float64, posY float64, floorTexture []color.RGBA, ceilingTexture []color.RGBA, pixels []color.RGBA, floorCeilTexture rl.Texture2D) {
	for y := 0; y < SCREEN_HEIGHT; y++ {
		rayDirX0 := dirX - planeX
		rayDirY0 := dirY - planeY
		rayDirX1 := dirX + planeX
		rayDirY1 := dirY + planeY

		p := float64(y) - SCREEN_HEIGHT/2

		posZ := 0.5 * SCREEN_HEIGHT

		rowDistance := posZ / p

		floorStepX := rowDistance * (rayDirX1 - rayDirX0) / SCREEN_WIDTH
		floorStepY := rowDistance * (rayDirY1 - rayDirY0) / SCREEN_WIDTH

		floorX := posX + rowDistance*rayDirX0
		floorY := posY + rowDistance*rayDirY0

		for x := 0; x < SCREEN_WIDTH; x++ {
			cellX := int(floorX)
			cellY := int(floorY)

			tx := int32(TEXTURE_WIDTH*(floorX-float64(cellX))) & (TEXTURE_WIDTH - 1)
			ty := int32(TEXTURE_HEIGHT*(floorY-float64(cellY))) & (TEXTURE_HEIGHT - 1)

			floorX += floorStepX
			floorY += floorStepY

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
	posX, posY := 22.0, 12.0
	dirX, dirY := -1.0, 0.0
	planeX, planeY := 0.0, 0.66

	// Time and physics initialization
	currentTime, oldTime := time.Now().UnixMilli(), int64(0)

	// Main loop
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		castCeiling(dirX, dirY, planeX, planeY, posX, posY, floorTexture, ceilingTexture, pixels, floorCeilTexture)
		castWalls(dirX, dirY, planeX, planeY, posX, posY, worldMap, walls)

		// Timing for FPS counter
		oldTime = currentTime
		currentTime = time.Now().UnixMilli()
		frameTime := (float64(currentTime) - float64(oldTime)) / 1000.0
		rl.DrawText(fmt.Sprintf("FPS: %d", int(1.0/frameTime)), 10, 10, 30, rl.White)

		moveSpeed := frameTime * 3.0
		rotSpeed := frameTime * 3.0

		if rl.IsKeyDown(rl.KeyUp) {
			if worldMap[int(posX+dirX*moveSpeed)][int(posY)] == 0 {
				posX += dirX * moveSpeed
			}
			if worldMap[int(posX)][int(posY+dirY*moveSpeed)] == 0 {
				posY += dirY * moveSpeed
			}
		}
		if rl.IsKeyDown(rl.KeyDown) {
			if worldMap[int(posX-dirX*moveSpeed)][int(posY)] == 0 {
				posX -= dirX * moveSpeed
			}
			if worldMap[int(posX)][int(posY-dirY*moveSpeed)] == 0 {
				posY -= dirY * moveSpeed
			}
		}
		if rl.IsKeyDown(rl.KeyRight) {
			oldDirX := dirX
			dirX = dirX*math.Cos(-rotSpeed) - dirY*math.Sin(-rotSpeed)
			dirY = oldDirX*math.Sin(-rotSpeed) + dirY*math.Cos(-rotSpeed)
			oldPlaneX := planeX
			planeX = planeX*math.Cos(-rotSpeed) - planeY*math.Sin(-rotSpeed)
			planeY = oldPlaneX*math.Sin(-rotSpeed) + planeY*math.Cos(-rotSpeed)
		}
		if rl.IsKeyDown(rl.KeyLeft) {
			oldDirX := dirX
			dirX = dirX*math.Cos(rotSpeed) - dirY*math.Sin(rotSpeed)
			dirY = oldDirX*math.Sin(rotSpeed) + dirY*math.Cos(rotSpeed)
			oldPlaneX := planeX
			planeX = planeX*math.Cos(rotSpeed) - planeY*math.Sin(rotSpeed)
			planeY = oldPlaneX*math.Sin(rotSpeed) + planeY*math.Cos(rotSpeed)
		}

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
