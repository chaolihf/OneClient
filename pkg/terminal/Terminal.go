package terminal

import (
	"fmt"
	"strconv"
	"strings"

	"com.chinatelecom.oneops.client/pkg/terminal/bring"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const (
	guacdAddress  = "localhost:4822"
	defaultWidth  = 800
	defaultHeight = 600
)

// Creates and initialize Bring's Session and Client
func createBringClient(protocol, hostname, port, username, password string) (*bring.Client, error) {
	client, err := bring.NewClient(guacdAddress, protocol, map[string]string{
		"hostname":    hostname,
		"port":        port,
		"username":    username,
		"password":    password,
		"width":       strconv.Itoa(defaultWidth),
		"height":      strconv.Itoa(defaultHeight),
		"ignore-cert": "true",
		"security":    "any",
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}

func ShowTerminal(bringApp fyne.App, protocol, hostname, port, username, password string) (fyne.Window, error) {
	client, err := createBringClient(protocol, hostname, port, username, password)
	if err != nil {
		return nil, err
	} else {
		title := fmt.Sprintf("%s (%s:%s)", strings.ToUpper(protocol), hostname, port)
		w := bringApp.NewWindow(title)
		bringDisplay := NewBringDisplay(client, defaultWidth, defaultHeight)
		w.CenterOnScreen()
		leftContent := container.New(layout.NewHBoxLayout(), widget.NewButton("退出", func() {
			bringApp.Quit()
		}))
		content := container.New(layout.NewVBoxLayout(), leftContent, bringDisplay)
		w.SetContent(content)
		w.Canvas().Focus(bringDisplay)
		return w, nil
	}
}
