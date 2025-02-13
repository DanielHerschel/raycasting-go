package objects

type Drawable interface {
	Draw(camera Camera)
	Close()
}
