package prefabs

import (
	"time"

	o "github.com/danielherschel/raylib-test/game/objects"
	ph "github.com/danielherschel/raylib-test/game/physics"
	u "github.com/danielherschel/raylib-test/game/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewTurret(x, y float32) *Turret {
	enemyTexture := rl.LoadTexture("assets/sprites/turret.png")

	transform := o.NewTransform(rl.NewVector2(x, y), rl.NewVector2(0.0, 0.0))

	return &Turret{
		Transform:  transform,
		HitBox:     o.NewHitBox(transform, 0.2),
		Sprite:     o.NewSprite(transform, enemyTexture),
		Health:     u.TURRET_HEALTH,
		ShouldDest: false,
		lookRadius: u.TURRET_LOOK_RADIUS,
	}
}

type Turret struct {
	o.Transform
	o.HitBox
	o.Sprite

	Health     int
	ShouldDest bool
	lookRadius float32

	lastShot time.Time
}

// IGameObject functions
func (t *Turret) Update(frameTime float64, currentLevel Level) {
	if !t.IsAlive() {
		t.ShouldDest = true
		return
	}

	player := currentLevel.Player
	directionToPlayer := rl.Vector2Normalize(rl.Vector2Subtract(player.Position, t.Position))

	hit := ph.CastRay(t, directionToPlayer, currentLevel.GetAllHittables())
	if hit != nil {
		if time.Since(t.lastShot) >= time.Second*u.TURRET_FIRE_RATE {
			player.TakeDamage(u.TURRET_DAMAGE)
			t.lastShot = time.Now()
		}
	}
}

func (t *Turret) GetTransform() o.Transform {
	return t.Transform
}

func (t *Turret) Close() {
	t.Sprite.Close()
}

// ISprite functions
func (t *Turret) GetSprite() o.Sprite {
	return t.Sprite
}

// IHittable functions
func (t *Turret) GetHitBox() o.HitBox {
	return t.HitBox
}

// IDestroyable functions
func (t *Turret) ShouldDestroy() bool {
	return t.ShouldDest
}

// IDamageable functions

func (t *Turret) TakeDamage(amount int) bool {
	t.Health -= amount
	return true
}

func (t *Turret) IsAlive() bool {
	return t.Health > 0
}
