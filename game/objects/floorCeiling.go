package objects

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"

	u "github.com/danielherschel/raylib-test/game/utils"
)

func NewFloorCeiling(floorTexture []color.RGBA, ceilingTexture []color.RGBA) FloorCeiling {
	floorCeilImage := rl.GenImageColor(u.SCREEN_WIDTH, u.SCREEN_HEIGHT, rl.White)
	floorCeilTexture := rl.LoadTextureFromImage(floorCeilImage)
	rl.UnloadImage(floorCeilImage)

	pixels := make([]color.RGBA, u.SCREEN_WIDTH*u.SCREEN_HEIGHT)

	return FloorCeiling{
		floorTexture:     floorTexture,
		ceilingTexture:   ceilingTexture,
		pixels:           pixels,
		floorCeilTexture: floorCeilTexture,
	}
}

type FloorCeiling struct {
	floorTexture     []color.RGBA
	ceilingTexture   []color.RGBA
	pixels           []color.RGBA
	floorCeilTexture rl.Texture2D
}

func (f FloorCeiling) Draw(camera Camera) {
	position := camera.Position
	dir := camera.Direction
	plane := camera.Plane

	for y := 0; y < u.SCREEN_HEIGHT; y++ {
		rayDir0 := rl.NewVector2(dir.X-plane.X, dir.Y-plane.Y)
		rayDir1 := rl.NewVector2(dir.X+plane.X, dir.Y+plane.Y)

		p := float32(y) - u.SCREEN_HEIGHT/2

		posZ := float32(0.5 * u.SCREEN_HEIGHT)

		rowDistance := posZ / p

		floorStep := rl.NewVector2(
			rowDistance*(rayDir1.X-rayDir0.X)/u.SCREEN_WIDTH,
			rowDistance*(rayDir1.Y-rayDir0.Y)/u.SCREEN_WIDTH,
		)

		floor := rl.NewVector2(position.X+rowDistance*rayDir0.X, position.Y+rowDistance*rayDir0.Y)

		for x := 0; x < u.SCREEN_WIDTH; x++ {
			cellX := int(floor.X)
			cellY := int(floor.Y)

			tx := int32(u.TEXTURE_WIDTH*(floor.X-float32(cellX))) & (u.TEXTURE_WIDTH - 1)
			ty := int32(u.TEXTURE_HEIGHT*(floor.Y-float32(cellY))) & (u.TEXTURE_HEIGHT - 1)

			floor.X += floorStep.X
			floor.Y += floorStep.Y

			color := f.floorTexture[u.TEXTURE_WIDTH*ty+tx]
			f.pixels[u.SCREEN_WIDTH*y+x] = color

			color = f.ceilingTexture[u.TEXTURE_WIDTH*ty+tx]
			f.pixels[u.SCREEN_WIDTH*(u.SCREEN_HEIGHT-y-1)+x] = color
		}
	}

	rl.UpdateTexture(f.floorCeilTexture, f.pixels)
	rl.DrawTexture(f.floorCeilTexture, 0, 0, rl.Gray)
}

func (f FloorCeiling) Close() {
	rl.UnloadTexture(f.floorCeilTexture)
	rl.UnloadImageColors(f.floorTexture)
	rl.UnloadImageColors(f.ceilingTexture)
}
