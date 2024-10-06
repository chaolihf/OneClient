package osrbrowser

import (
	"fmt"
	"image"

	"com.chinatelecom.oneops.client/pkg/terminal/bring"
	"fyne.io/fyne/v2"
)

type OsrBrowserClient struct {
}

func (r *OsrBrowserClient) SendMouse(p image.Point, pressedButtons ...bring.MouseButton) error {
	fmt.Println("SendMouse")
	//ui.GoSendMouseEvent()
	return nil
}

func (r *OsrBrowserClient) OnSync(f bring.OnSyncFunc) {
	fmt.Println("OnSync")
}

func (r *OsrBrowserClient) Screen() (image image.Image, lastUpdate int64) {
	fmt.Println("Screen")
	return nil, 0
}

func (r *OsrBrowserClient) SendKey(key bring.KeyCode, pressed bool) error {
	fmt.Println("SendKey")
	return nil
}

func ShowBrowser(browserApp fyne.App, width, height int) (fyne.Window, error) {
	title := "OffScreen Render Browser"
	w := browserApp.NewWindow(title)
	osrBrowserClient := OsrBrowserClient{}
	osrBrowser := NewOsrBrowser(&osrBrowserClient, width, height)
	w.CenterOnScreen()
	w.SetContent(osrBrowser)
	w.Canvas().Focus(osrBrowser)
	return w, nil
}
