package prefabs

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	o "github.com/danielherschel/raylib-test/game/objects"
	"github.com/danielherschel/raylib-test/game/schemas"
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
	var gameObjects []IGameObject
	for _, gameObjectData := range gameObjectsData {
		objectPosition := gameObjectData.Position

		switch gameObjectData.Type {
		case "barrel":
			gameObjects = append(gameObjects, NewBarrel(float32(objectPosition[0]), float32(objectPosition[1])))
		case "pillar":
			gameObjects = append(gameObjects, NewPillar(float32(objectPosition[0]), float32(objectPosition[1])))
		case "turret":
			gameObjects = append(gameObjects, NewTurret(float32(objectPosition[0]), float32(objectPosition[1])))
		}
	}

	// Load map data
	walls := NewWalls(levelData.WorldMap)
	floorCeiling := NewFloorCeiling()

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
	GameObjects  GameObjects
	Walls        Walls
	FloorCeiling FloorCeiling
}

func (l *Level) Close() {
	l.Walls.Close()
	l.FloorCeiling.Close()
	l.GameObjects.Close()
}

func (l *Level) GetAllHittables() (hittables []o.IHittable) {
	// Add the game objects to the hittables list
	for _, obj := range l.GameObjects {
		if hittable, ok := obj.(o.IHittable); ok {
			hittables = append(hittables, hittable)
		}
	}

	// Add the walls to the hittables list
	for _, wall := range l.Walls.HitBoxes {
		hittables = append(hittables, wall)
	}

	// TODO: add the enemies to the hittables list

	return
}
