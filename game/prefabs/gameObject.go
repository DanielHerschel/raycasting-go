package prefabs

import (
	o "github.com/danielherschel/raylib-test/game/objects"
	u "github.com/danielherschel/raylib-test/game/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type IGameObject interface {
	GetTransform() o.Transform
	Update(frameTime float64, currentLevel Level)
	Close()
}

func SortGameObjectsByDistanceToPoint(pos rl.Vector2, gameObjects GameObjects) (sorted GameObjects) {
	spriteData := make(map[int]float32, len(gameObjects))
	for i := 0; i < len(gameObjects); i++ {
		spriteData[i] = ((pos.X-gameObjects[i].GetTransform().Position.X)*(pos.X-gameObjects[i].GetTransform().Position.X) + (pos.Y-gameObjects[i].GetTransform().Position.Y)*(pos.Y-gameObjects[i].GetTransform().Position.Y))
	}

	pairList := u.SortMap(spriteData)

	for i := 0; i < len(pairList); i++ {
		index := pairList[i].Key
		sorted = append(sorted, gameObjects[index])
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
