package prefabs

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	o "github.com/danielherschel/raylib-test/game/objects"
	"github.com/danielherschel/raylib-test/game/schemas"
	u "github.com/danielherschel/raylib-test/game/utils"
)

// Level
func NewLevelFromFile(path string) *Level {
	levelData := schemas.NewLevelSchemaFromFile(path)

	// Extract player data
	playerPosition := levelData.PlayerStartData.Position
	playerDirection := levelData.PlayerStartData.Direction
	player := NewPlayer(o.Transform{
		Position:  rl.NewVector2(float32(playerPosition[0]), float32(playerPosition[1])),
		Direction: rl.NewVector2(float32(playerDirection[0]), float32(playerDirection[1])),
	})

	// Extract game objects data
	gameObjectsData := levelData.GameObjectsData
	var gameObjects []o.IGameObject
	for _, gameObjectData := range gameObjectsData {
		objectPosition := gameObjectData.Position

		switch gameObjectData.Type {
		case "barrel":
			gameObjects = append(gameObjects, NewBarrel(float32(objectPosition[0]), float32(objectPosition[1])))
		case "pillar":
			gameObjects = append(gameObjects, NewPillar(float32(objectPosition[0]), float32(objectPosition[1])))
		}
	}

	// Load map data
	walls := NewWalls(levelData.WorldMap)

	floorImage, ceilingImage := rl.LoadImage(u.TEXTURE_STONE_BRICKS), rl.LoadImage(u.TEXTURE_WOOD)
	floorTexture := rl.LoadImageColors(floorImage)
	ceilingTexture := rl.LoadImageColors(ceilingImage)
	u.UnloadImages(floorImage, ceilingImage)

	floorCeiling := NewFloorCeiling(floorTexture, ceilingTexture)

	return &Level{
		Id:           levelData.ID,
		Name:         levelData.Name,
		WorldMap:     levelData.WorldMap,
		Player:       player,
		GameObjects:  gameObjects,
		Walls:        walls,
		FloorCeiling: floorCeiling,
	}
}

type Level struct {
	Id       int
	Name     string
	WorldMap [][]int

	Player       *Player
	GameObjects  o.GameObjects
	Walls        Walls
	FloorCeiling FloorCeiling
}

func (l *Level) Close() {
	l.Walls.Close()
	l.FloorCeiling.Close()
	l.GameObjects.Close()
}
