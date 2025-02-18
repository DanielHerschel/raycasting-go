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

// Level data from file
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
	if err != nil {
		panic(err)
	}
	defer levelFile.Close()

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

// Level struct
func NewLevel(levelFilePath string) *Level {
	levelData := loadLevelDataFromFile(levelFilePath)
	fmt.Print(levelData)

	// Load map data
	worldMap := levelData.WorldMap

	walls := NewWalls(worldMap)

	floorImage, ceilingImage := rl.LoadImage(u.TEXTURE_STONE_BRICKS), rl.LoadImage(u.TEXTURE_WOOD)
	floorTexture := rl.LoadImageColors(floorImage)
	ceilingTexture := rl.LoadImageColors(ceilingImage)
	u.UnloadImages(floorImage, ceilingImage)

	floorCeiling := NewFloorCeiling(floorTexture, ceilingTexture)

	// Camera settings
	camera := o.NewCamera(
		levelData.Player,
		u.CAMERA_FOV,
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
	GameObjects  o.GameObjects

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
	l.updateGameObjects()

	// Timing for FPS counter
	l.frameTime = l.getFrameTime()
	rl.DrawText(fmt.Sprintf("FPS: %d", int(1.0/l.frameTime)), 10, 10, 30, rl.White)

	// Update camera
	l.Camera.Update(l.frameTime, l.WorldMap)
}

func (l *Level) updateGameObjects() {
	var indicesToRemove []int
	var gameObjectsHit []o.IHittable

	l.GameObjects = o.SortGameObjectsByDistanceToCamera(*l.Camera, l.GameObjects)

	for index, gameObject := range l.GameObjects {
		// Check for collision
		if hittable, ok := gameObject.(o.IHittable); ok {
			if hittable.GetHitBox().CheckCollision(l.Camera.Transform) {
				gameObjectsHit = append(gameObjectsHit, hittable)
			}
		}

		// Destroy destroyable objects
		toDraw := true
		if destroyable, ok := gameObject.(o.IDestroyable); ok {
			if destroyable.ShouldDestroy() {
				indicesToRemove = append(indicesToRemove, index)
				toDraw = false
			}
		}

		// Draw object
		if toDraw {
			if sprite, ok := gameObject.(o.ISprite); ok {
				sprite.GetSprite().Draw(*l.Camera)
			}
		}
	}
	// Run the OnHit function of the last object hit - the closest one to the camera
	if len(gameObjectsHit) > 0 {
		gameObjectsHit[len(gameObjectsHit)-1].OnHit()
	}

	// Remove destroyable objects in reverse order
	for i := len(indicesToRemove) - 1; i >= 0; i-- {
		index := indicesToRemove[i]
		l.GameObjects[index].Close()
		l.GameObjects = l.GameObjects.Remove(index)
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
	l.GameObjects.Close()
}
