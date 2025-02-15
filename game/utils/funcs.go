package utils

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func UnloadTextures(textures ...rl.Texture2D) {
	for _, texture := range textures {
		rl.UnloadTexture(texture)
	}
}

func UnloadImages(images ...*rl.Image) {
	for _, image := range images {
		rl.UnloadImage(image)
	}
}
