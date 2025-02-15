package objects

import (

	u "github.com/danielherschel/raylib-test/game/utils"
)

type IGameObject interface {
	GetSprite() Sprite
	Close()
}

func SortGameObjectsByDistanceToCamera(camera Camera, sprites []IGameObject) (sorted []IGameObject) {
	spriteData := make(map[int]float32, len(sprites))
	for i := 0; i < len(sprites); i++ {
		spriteData[i] = ((camera.Position.X-sprites[i].GetSprite().Position.X)*(camera.Position.X-sprites[i].GetSprite().Position.X) + (camera.Position.Y-sprites[i].GetSprite().Position.Y)*(camera.Position.Y-sprites[i].GetSprite().Position.Y))
	}

	pairList := u.SortMap(spriteData)

	for i := 0; i < len(pairList); i++ {
		index := pairList[i].Key
		sorted = append(sorted, sprites[index])
	}
	return
}

func UnloadGameObjects(gameObjects []IGameObject) {
	for i := 0; i < len(gameObjects); i++ {
		gameObjects[i].Close()
	}
}
