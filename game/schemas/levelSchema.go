package schemas

import (
	"encoding/json"
	"os"
)

// Level Schema
func NewLevelSchemaFromFile(path string) (levelData LevelSchema) {
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
	err = json.Unmarshal(fileData, &levelData)
	if err != nil {
		panic(err)
	}
	return
}

type LevelSchema struct {
	ID              int                `json:"id"`
	Name            string             `json:"name"`
	WorldMap        [][]int            `json:"worldMap"`
	PlayerStartData PlayerStartSchema  `json:"player"`
	GameObjectsData []GameObjectSchema `json:"gameObjects"`
}

type PlayerStartSchema struct {
	Position  []float64 `json:"position"`
	Direction []float64 `json:"direction"`
}

type GameObjectSchema struct {
	Type     string    `json:"type"`
	Position []float64 `json:"position"`
}