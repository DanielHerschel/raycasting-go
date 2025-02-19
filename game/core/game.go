package core

import (
	"fmt"
	"time"

	o "github.com/danielherschel/raylib-test/game/objects"
	p "github.com/danielherschel/raylib-test/game/prefabs"
	u "github.com/danielherschel/raylib-test/game/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Game struct
func NewGame() *Game {
	levels := []*p.Level{p.NewLevelFromFile(u.LEVEL_1_PATH)}
	currentLevel := levels[0]


	// Time and physics iunitialization
	currentTime, oldTime := time.Now().UnixMilli(), int64(0)

	return &Game{
		Levels: levels,
		CurrentLevel: currentLevel,
		currentTime:  currentTime,
		oldTime:      oldTime,
	}
}

type Game struct {
	Levels []*p.Level
	CurrentLevel *p.Level

	// Time and physics
	currentTime int64
	oldTime     int64
	frameTime   float64
}

func (g *Game) MainLoop() {
	playerCamera := *g.CurrentLevel.Player.Camera

	// Draw world
	g.CurrentLevel.FloorCeiling.Draw(playerCamera)
	g.CurrentLevel.Walls.Draw(playerCamera)

	// Draw Sprites
	g.updateGameObjects()

	// Timing for FPS counter
	g.frameTime = g.getFrameTime()
	rl.DrawText(fmt.Sprintf("FPS: %d", int(1.0/g.frameTime)), 10, 10, 30, rl.White)

	// Update camera
	g.CurrentLevel.Player.Update(g.frameTime, g.CurrentLevel.WorldMap)
}

func (g *Game) updateGameObjects() {
	var indicesToRemove []int
	var gameObjectsHit []o.IHittable

	g.CurrentLevel.GameObjects = o.SortGameObjectsByDistanceToPoint(g.CurrentLevel.Player.Position, g.CurrentLevel.GameObjects)

	for index, gameObject := range g.CurrentLevel.GameObjects {
		// Check for crosshair collision
		if hittable, ok := gameObject.(o.IHittable); ok {
			if hittable.GetHitBox().CheckCollision(g.CurrentLevel.Player.Transform) {
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

		// Draw sprites
		if toDraw {
			if sprite, ok := gameObject.(o.ISprite); ok {
				sprite.GetSprite().Draw(*g.CurrentLevel.Player.Camera)
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
		g.CurrentLevel.GameObjects[index].Close()
		g.CurrentLevel.GameObjects = g.CurrentLevel.GameObjects.Remove(index)
	}
}

func (g *Game) getFrameTime() float64 {
	g.oldTime = g.currentTime
	g.currentTime = time.Now().UnixMilli()
	return float64(g.currentTime-g.oldTime) / 1000.0
}

func (g *Game) Close() {
	for _, level := range g.Levels {
        level.Close()
    }
}

