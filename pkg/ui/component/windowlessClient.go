package component

import (
	"image"

	"com.chinatelecom.oneops.client/pkg/terminal/bring"
)

type WindowlessClient interface {
	SendMouse(p image.Point, pressedButtons ...bring.MouseButton) error
	OnSync(f bring.OnSyncFunc)
	Screen() (image image.Image, lastUpdate int64)
	SendKey(key bring.KeyCode, pressed bool) error
}
