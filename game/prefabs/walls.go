package prefabs

import (
	"math"

	u "github.com/danielherschel/raylib-test/game/utils"
	o "github.com/danielherschel/raylib-test/game/objects"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewWalls(worldMap [][]int, wallTextures []rl.Texture2D) Walls {
	return Walls{worldMap: worldMap, wallTextures: wallTextures}
}

type Walls struct {
	worldMap     [][]int
	wallTextures []rl.Texture2D
}

func (w Walls) Draw(camera o.Camera) {
	position := camera.Position
	dir := camera.Direction
	plane := camera.Plane
	
	for x := 0; x < u.SCREEN_WIDTH; x++ {
		cameraX := 2*float32(x)/float32(u.SCREEN_WIDTH) - 1
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
			if w.worldMap[mapX][mapY] > 0 {
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

		lineHeight := int(u.SCREEN_HEIGHT / prepWallDist)

		drawStart := int(-float64(lineHeight)/2 + u.SCREEN_HEIGHT/2)
		drawEnd := int(float64(lineHeight)/2 + u.SCREEN_HEIGHT/2)

		wallType := w.worldMap[mapX][mapY] - 1
		texturePos := float32(wallX * u.TEXTURE_WIDTH)

		rl.DrawTexturePro(
			w.wallTextures[wallType],
			rl.NewRectangle(texturePos, 0, 1, u.TEXTURE_HEIGHT),
			rl.NewRectangle(float32(x), float32(drawStart), 1, float32(drawEnd-drawStart)),
			rl.NewVector2(0.0, 0.0),
			0.0,
			rl.Gray,
		)

		// Save wall distance in this camera x position in the buffer
		camera.ZBuffer[x] = prepWallDist
	}
}

func (w Walls) Close() {
	u.UnloadTextures(w.wallTextures)
}
