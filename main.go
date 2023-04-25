package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/v2/app"
	"github.com/adehlbom/WebPenAuto/gui"
)

func main() {
	application := app.New()
	window := application.NewWindow("Web App Penetration Testing Tool")

	gui.SetupMainWindow(window)

	window.Resize(fyne.NewSize(400, 300))
	window.ShowAndRun()
}
