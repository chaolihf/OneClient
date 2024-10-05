package ui

import (
	"fmt"

	"com.chinatelecom.oneops.client/pkg/terminal"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
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

	content := widget.NewButton("打开远程桌面", func() {
		w, err := terminal.ShowTerminal(a, "rdp", "192.168.1.128", "3389", "lichao", "lc2013!")
		if err == nil {
			w.Show()
		} else {
			fmt.Println(err)
		}
	})
	myWindow.SetContent(content)
	myWindow.SetOnDropped(fileDraggable)
	myWindow.ShowAndRun()
}
