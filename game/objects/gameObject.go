package objects

import (
	u "github.com/danielherschel/raylib-test/game/utils"
)

type IGameObject interface {
	GetSprite() Sprite
	Close()
}

func SortGameObjectsByDistanceToCamera(camera Camera, sprites GameObjects) (sorted GameObjects) {
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

type GameObjects []IGameObject

func (g GameObjects) Close() {
	for i := 0; i < len(g); i++ {
		g[i].Close()
	}
}

func (g GameObjects) Remove(index int) GameObjects {
	return append(g[:index], g[index+1:]...)
}
