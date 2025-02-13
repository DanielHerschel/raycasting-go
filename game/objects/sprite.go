package objects

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"

	u "github.com/danielherschel/raylib-test/game/utils"
)

func NewSprite(transform Transform, texture rl.Texture2D) Sprite {
	return Sprite{Transform: transform, Texture: texture}
}

type Sprite struct {
	Transform
	Texture rl.Texture2D
}

func (s Sprite) Draw(camera Camera) {
	spritePos := rl.NewVector2(s.Position.X-camera.Position.X, s.Position.Y-camera.Position.Y)

	invDet := 1.0 / (camera.Plane.X*camera.Direction.Y - camera.Direction.X*camera.Plane.Y)

	transformVec := rl.NewVector2(
		invDet*(camera.Direction.Y*spritePos.X-camera.Direction.X*spritePos.Y),
		invDet*(-camera.Plane.Y*spritePos.X+camera.Plane.X*spritePos.Y),
	)

	spriteScreenX := int((u.SCREEN_WIDTH / 2) * (1 + transformVec.X/transformVec.Y))

	spriteHeight := int(math.Abs(u.SCREEN_HEIGHT / float64(transformVec.Y)))

	// Calc height of the sprite
	drawStartY := -spriteHeight/2 + u.SCREEN_HEIGHT/2
	drawEndY := spriteHeight/2 + u.SCREEN_HEIGHT/2

	// Calc width of the sprite
	spriteWidth := int(math.Abs(u.SCREEN_HEIGHT / float64(transformVec.Y)))
	drawStartX := -spriteWidth/2 + spriteScreenX
	drawEndX := spriteWidth/2 + spriteScreenX

	// Draw the sprite
	for stripe := drawStartX; stripe < drawEndX; stripe++ {
		texX := int(256 * float64((stripe-(-spriteWidth/2+spriteScreenX))*u.TEXTURE_WIDTH) / float64(spriteWidth) / 256)
		if transformVec.Y > 0 && stripe > 0 && stripe < u.SCREEN_WIDTH && transformVec.Y < camera.ZBuffer[stripe] {
			rl.DrawTexturePro(
				s.Texture,
				rl.NewRectangle(float32(texX), 0, 1, u.TEXTURE_HEIGHT),
				rl.NewRectangle(float32(stripe), float32(drawStartY), 1, float32(drawEndY-drawStartY)),
				rl.NewVector2(0.0, 0.0),
				0.0,
				rl.Gray,
			)
		}
	}
}

func (s Sprite) Close() {
	rl.UnloadTexture(s.Texture)
}

func SortSprites(camera Camera, sprites []Sprite) (sorted []Sprite) {
	spriteData := make(map[int]float32, len(sprites))
	for i := 0; i < len(sprites); i++ {
		spriteData[i] = ((camera.Position.X-sprites[i].Position.X)*(camera.Position.X-sprites[i].Position.X) + (camera.Position.Y-sprites[i].Position.Y)*(camera.Position.Y-sprites[i].Position.Y))
	}

	pairList := u.SortMap(spriteData)

	for i := 0; i < len(pairList); i++ {
		index := pairList[i].Key
		sorted = append(sorted, sprites[index])
	}
	return
}

func UnloadSprites(sprites []Sprite) {
	for _, sprite := range sprites {
		sprite.Close()
	}
}
