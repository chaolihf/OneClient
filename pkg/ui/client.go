package ui

import (
	"fmt"

	"com.chinatelecom.oneops.client/pkg/terminal"
	"com.chinatelecom.oneops.client/pkg/ui/component/osrbrowser"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func fileDraggable(p fyne.Position, files []fyne.URI) {
	if len(files) > 0 {
		fmt.Println(files[0].Path())
	}
}

func ShowRemoteDesktop() {
	a := app.New()
	myWindow := a.NewWindow("终端")

	openDesktopButton := widget.NewButton("打开远程桌面", func() {
		w, err := terminal.ShowTerminal(a, "rdp", "", "", "", "")
		if err == nil {
			w.Show()
		} else {
			fmt.Println(err)
		}
	})
	openOsrBrowserButton := widget.NewButton("打开浏览器", func() {
		w, err := osrbrowser.ShowBrowser(a, 800, 600)
		if err == nil {
			w.Show()
		} else {
			fmt.Println(err)
		}
	})
	content := container.New(layout.NewVBoxLayout(), openDesktopButton, openOsrBrowserButton)
	myWindow.SetContent(content)
	myWindow.SetOnDropped(fileDraggable)
	myWindow.ShowAndRun()
}
