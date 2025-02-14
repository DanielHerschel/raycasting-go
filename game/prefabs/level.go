package prefabs

import (
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"

	o "github.com/danielherschel/raylib-test/game/objects"
	u "github.com/danielherschel/raylib-test/game/utils"
)

// Hard coded level data, will be moved to a file later

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

func getSprites() (sprites []o.ISprite) {
	sprites = append(sprites, NewBarrel(21.5, 1.5))  // barrel
	sprites = append(sprites, NewBarrel(15.5, 1.5))  // barrel
	sprites = append(sprites, NewBarrel(16.0, 1.8))  // barrel
	sprites = append(sprites, NewBarrel(16.2, 1.2))  // barrel
	sprites = append(sprites, NewBarrel(3.5, 2.5))   // barrel
	sprites = append(sprites, NewBarrel(9.5, 15.5))  // barrel
	sprites = append(sprites, NewBarrel(10.0, 15.1)) // barrel
	sprites = append(sprites, NewBarrel(10.5, 15.8)) // barrel

	sprites = append(sprites, NewPillar(18.5, 10.5)) // pillar
	sprites = append(sprites, NewPillar(18.5, 11.5)) // pillar
	sprites = append(sprites, NewPillar(18.5, 12.5)) // pillar

	return
}

// Level struct

func NewLevel() *Level {
	// Texture loading and initialization
	wallsImages, wallTextures := getWalls()

	floorTexture := rl.LoadImageColors(wallsImages[3])
	ceilingTexture := rl.LoadImageColors(wallsImages[6])
	u.UnloadImages(wallsImages)

	// Load map data
	worldMap := getWorldMap()

	walls := NewWalls(worldMap, wallTextures)

	floorCeiling := NewFloorCeiling(floorTexture, ceilingTexture)

	// Camera settings
	camera := o.NewCamera(
		o.NewTransform(rl.NewVector2(22.0, 12.0), rl.NewVector2(-1.0, 0.0)),
		rl.NewVector2(0.0, 0.66),
	)

	// Sprites
	sprites := getSprites()

	// Time and physics iunitialization
	currentTime, oldTime := time.Now().UnixMilli(), int64(0)

	return &Level{
		WorldMap:     worldMap,
		Walls:        walls,
		FloorCeiling: floorCeiling,
		Sprites:      sprites,
		Camera:       camera,
		currentTime:  currentTime,
		oldTime:      oldTime,
	}
}

type Level struct {
	WorldMap     [][]int
	Walls        Walls
	FloorCeiling FloorCeiling
	Sprites      []o.ISprite

	Camera *o.Camera

	// Time and physics
	currentTime int64
	oldTime     int64
	frameTime   float64
}

func (l *Level) MainLoop() {
	// Draw world
	l.FloorCeiling.Draw(*l.Camera)
	l.Walls.Draw(*l.Camera)

	// Draw Sprites
	l.drawSprites()

	// Timing for FPS counter
	l.frameTime = l.getFrameTime()
	rl.DrawText(fmt.Sprintf("FPS: %d", int(1.0/l.frameTime)), 10, 10, 30, rl.White)

	// Update camera
	l.Camera.Update(l.frameTime, l.WorldMap)
}

func (l *Level) drawSprites() {
	l.Sprites = o.SortSprites(*l.Camera, l.Sprites)
	for _, sprite := range l.Sprites {
		sprite.Draw(*l.Camera)
	}
}

func (l *Level) getFrameTime() float64 {
	l.oldTime = l.currentTime
	l.currentTime = time.Now().UnixMilli()
	return float64(l.currentTime-l.oldTime) / 1000.0
}

func (l *Level) Close() {
	l.Walls.Close()
	l.FloorCeiling.Close()
	o.UnloadSprites(l.Sprites)
}
