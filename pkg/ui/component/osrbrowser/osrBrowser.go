package osrbrowser

import (
	"image"
	"image/color"

	"com.chinatelecom.oneops.client/pkg/ui/component"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type osrBrowserRenderer struct {
	objects []fyne.CanvasObject
	remote  *OsrBrowser
}

func (r *osrBrowserRenderer) MinSize() fyne.Size {
	return r.remote.MinSize()
}

func (r *osrBrowserRenderer) Layout(size fyne.Size) {
	if len(r.objects) == 0 {
		return
	}

	r.objects[0].Resize(size)
}

func (r *osrBrowserRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *osrBrowserRenderer) BackgroundColor() color.Color {
	return color.Transparent
}

func (r *osrBrowserRenderer) Refresh() {
	if len(r.objects) == 0 {
		raster := canvas.NewImageFromImage(r.remote.Display)
		raster.FillMode = canvas.ImageFillContain
		r.objects = append(r.objects, raster)
	} else {
		r.objects[0].(*canvas.Image).Image = r.remote.Display
	}
	r.Layout(r.remote.Size())
	canvas.Refresh(r.remote)
}

func (r *osrBrowserRenderer) Destroy() {
}

// Custom Widget that represents the remote computer being controlled
type OsrBrowser struct {
	widget.BaseWidget
	keyboardHandler
	mouseHandler

	lastUpdate int64

	Display image.Image
	Client  component.WindowlessClient
}

// Creates a new BringDisplay and does all the heavy lifting, setting up all event handlers
func NewOsrBrowser(client component.WindowlessClient, width, height int) *OsrBrowser {
	empty := image.NewNRGBA(image.Rect(0, 0, width-1, height-1))

	b := &OsrBrowser{
		Client: client,
	}
	b.keyboardHandler.display = b
	b.mouseHandler.display = b

	b.SetDisplay(empty)
	b.Client.OnSync(func(img image.Image, ts int64) {
		if ts == b.lastUpdate {
			return
		}
		b.lastUpdate = ts
		b.SetDisplay(img)
	})
	return b
}

func (b *OsrBrowser) MinSize() fyne.Size {
	b.ExtendBaseWidget(b)
	return fyne.Size{
		Width:  float32(b.Display.Bounds().Dx()),
		Height: float32(b.Display.Bounds().Dy()),
	}
}

func (b *OsrBrowser) CreateRenderer() fyne.WidgetRenderer {
	return &osrBrowserRenderer{
		objects: []fyne.CanvasObject{},
		remote:  b,
	}
}

// This forces a display refresh if there are any pending updates, instead of waiting for the next
// OnSync event. Ex: when moving the mouse we want instant feedback of the new cursor position
func (b *OsrBrowser) updateDisplay() {
	img, ts := b.Client.Screen()
	if ts != b.lastUpdate {
		b.SetDisplay(img)
		b.lastUpdate = ts
	}
}

func (b *OsrBrowser) SetDisplay(img image.Image) {
	b.Display = img
	b.Refresh()
}
