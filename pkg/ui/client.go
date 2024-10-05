package ui

import (
	"fmt"

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
	w := a.NewWindow("远程桌面")
	w.SetOnDropped(fileDraggable)
	w.SetContent(widget.NewLabel("可以拖入上传文件"))
	w.ShowAndRun()
}
