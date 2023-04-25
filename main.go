package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/adehlbom/WebPenAuto/gui"
	"github.com/adehlbom/WebPenAuto/report"
)

func main() {
	a := app.New()
	w := a.NewWindow("Web App Penetration Testing")

	// Create the export button
	exportButton := gui.NewExportButton([]report.TestResult{})

	// Pass the export button to SetupMainWindow
	gui.SetupMainWindow(w, exportButton)

	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}
