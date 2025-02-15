package prefabs

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"

	o "github.com/danielherschel/raylib-test/game/objects"
	u "github.com/danielherschel/raylib-test/game/utils"
)

// Hard coded level data, will be moved to a file later
type LevelData struct {
	Id          int
	Name        string
	WorldMap    [][]int
	Player      o.Transform
	GameObjects []o.IGameObject
}

func loadLevelDataFromFile(path string) LevelData {
	// Open the file
	levelFile, err := os.Open(path)
	defer levelFile.Close()
	if err != nil {
		panic(err)
	}

	fileInfo, err := levelFile.Stat()
	if err != nil {
		panic(err)
	}

	fileSize := fileInfo.Size()
	fileData := make([]byte, fileSize)
	_, err = levelFile.Read(fileData)
	if err != nil {
		panic(err)
	}

	// Parse the JSON data
	var levelData map[string]interface{}
	err = json.Unmarshal(fileData, &levelData)
	if err != nil {
		panic(err)
	}

	// Extract the level data
	id := int(levelData["id"].(float64))
	name := levelData["name"].(string)

	// Extract world map data
	worldMapData := levelData["worldMap"].([]interface{})
	var worldMap [][]int
	for _, row := range worldMapData {
		row := row.([]interface{})
		var intRow []int
		for _, cell := range row {
			intRow = append(intRow, int(cell.(float64)))
		}
		worldMap = append(worldMap, intRow)
	}

	// Extract player data
	playerData := levelData["player"].(map[string]interface{})
	playerPosition := playerData["position"].([]interface{})
	playerDirection := playerData["direction"].([]interface{})
	player := o.Transform{
		Position:  rl.NewVector2(float32(playerPosition[0].(float64)), float32(playerPosition[1].(float64))),
		Direction: rl.NewVector2(float32(playerDirection[0].(float64)), float32(playerDirection[1].(float64))),
	}

	// Extract game objects data
	gameObjectsData := levelData["gameObjects"].([]interface{})
	var gameObjects []o.IGameObject
	for _, gameObjectData := range gameObjectsData {
		gameObject := gameObjectData.(map[string]interface{})
		objectPosition := gameObject["position"].([]interface{})

		switch gameObject["type"].(string) {
		case "barrel":
			gameObjects = append(gameObjects, NewBarrel(float32(objectPosition[0].(float64)), float32(objectPosition[1].(float64))))
		case "pillar":
			gameObjects = append(gameObjects, NewPillar(float32(objectPosition[0].(float64)), float32(objectPosition[1].(float64))))
		}
	}

	return LevelData{
		Id:          id,
		Name:        name,
		WorldMap:    worldMap,
		Player:      player,
		GameObjects: gameObjects,
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

// Level struct

func NewLevel(levelPath string) *Level {
	levelData := loadLevelDataFromFile(levelPath)
	fmt.Print(levelData)

	// Texture loading and initialization
	wallsImages, wallTextures := getWalls()

	floorTexture := rl.LoadImageColors(wallsImages[3])
	ceilingTexture := rl.LoadImageColors(wallsImages[6])
	u.UnloadImages(wallsImages)

	// Load map data
	worldMap := levelData.WorldMap

	walls := NewWalls(worldMap, wallTextures)

	floorCeiling := NewFloorCeiling(floorTexture, ceilingTexture)

	// Camera settings
	camera := o.NewCamera(
		levelData.Player,
		rl.NewVector2(0.0, 0.66),
	)

	// Load Game Objects
	gameObjects := levelData.GameObjects

	// Time and physics iunitialization
	currentTime, oldTime := time.Now().UnixMilli(), int64(0)

	return &Level{
		WorldMap:     worldMap,
		Walls:        walls,
		FloorCeiling: floorCeiling,
		GameObjects:  gameObjects,
		Camera:       camera,
		currentTime:  currentTime,
		oldTime:      oldTime,
	}
}

type Level struct {
	WorldMap     [][]int
	Walls        Walls
	FloorCeiling FloorCeiling
	GameObjects  []o.IGameObject

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
	l.GameObjects = o.SortGameObjectsByDistanceToCamera(*l.Camera, l.GameObjects)
	for _, sprite := range l.GameObjects {
		sprite.GetSprite().Draw(*l.Camera)
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
	o.UnloadGameObjects(l.GameObjects)
}
