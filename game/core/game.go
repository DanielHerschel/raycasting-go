package core

import (
	"fmt"
	"time"

	o "github.com/danielherschel/raylib-test/game/objects"
	p "github.com/danielherschel/raylib-test/game/prefabs"
	u "github.com/danielherschel/raylib-test/game/utils"
	ph "github.com/danielherschel/raylib-test/game/physics"

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
	// Timing for FPS counter
	g.frameTime = g.getFrameTime()
	rl.DrawText(fmt.Sprintf("FPS: %d", int(1.0/g.frameTime)), 10, 10, 30, rl.White)

	playerCamera := *g.CurrentLevel.Player.Camera

	// Draw world
	g.CurrentLevel.FloorCeiling.Draw(playerCamera)
	g.CurrentLevel.Walls.Draw(playerCamera)

	// Draw Sprites
	g.updateGameObjects()

	// Update camera
	g.CurrentLevel.Player.Update(g.frameTime, g.CurrentLevel.WorldMap)
}

func (g *Game) updateGameObjects() {
	var indicesToRemove []int

	player := g.CurrentLevel.Player

	g.CurrentLevel.GameObjects = o.SortGameObjectsByDistanceToPoint(g.CurrentLevel.Player.Position, g.CurrentLevel.GameObjects)

	// check what is the crosshair is looking at
	hittables := g.getAllHittables()
	hittable := ph.CastRay(player.Position, player.Direction, hittables)
	if hittable != nil {
        hittable.OnHit()
    }

	for index, gameObject := range g.CurrentLevel.GameObjects {
		// Destroy destroyable objects
		toDraw := true
		if destroyable, ok := gameObject.(o.IDestroyable); ok {
			if destroyable.ShouldDestroy() {
				indicesToRemove = append(indicesToRemove, index)
				toDraw = false
			}
		}

		// Draw the gameobject's sprite
		if toDraw {
			if sprite, ok := gameObject.(o.ISprite); ok {
				sprite.GetSprite().Draw(*player.Camera)
			}
		}
	}

	// Remove destroyable objects in reverse order
	for i := len(indicesToRemove) - 1; i >= 0; i-- {
		index := indicesToRemove[i]
		g.CurrentLevel.GameObjects[index].Close()
		g.CurrentLevel.GameObjects = g.CurrentLevel.GameObjects.Remove(index)
	}
}

func (g *Game) getAllHittables() (hittables []o.IHittable) {
	// Add the game objects to the hittables list
	for _, obj := range g.CurrentLevel.GameObjects {
		if hittable, ok := obj.(o.IHittable); ok {
            hittables = append(hittables, hittable)
        }
	}

	// Add the walls to the hittables list
	for _, wall := range g.CurrentLevel.Walls.HitBoxes {
		hittables = append(hittables, wall)
	}

	// TODO: add the enemies to the hittables list

	return
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

