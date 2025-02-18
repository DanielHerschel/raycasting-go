package objects

import (
	u "github.com/danielherschel/raylib-test/game/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type IGameObject interface {
	GetTransform() Transform
	Close()
}

func SortGameObjectsByDistanceToPoint(pos rl.Vector2, sprites GameObjects) (sorted GameObjects) {
	spriteData := make(map[int]float32, len(sprites))
	for i := 0; i < len(sprites); i++ {
		spriteData[i] = ((pos.X-sprites[i].GetTransform().Position.X)*(pos.X-sprites[i].GetTransform().Position.X) + (pos.Y-sprites[i].GetTransform().Position.Y)*(pos.Y-sprites[i].GetTransform().Position.Y))
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
