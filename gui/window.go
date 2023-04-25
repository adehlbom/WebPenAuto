package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func SetupMainWindow(window fyne.Window) {
	urlEntry := widget.NewEntry()
	urlEntry.SetPlaceHolder("Enter target URL")

	output := widget.NewLabel("")

	scanButton := widget.NewButton("Scan", func() {
		url := urlEntry.Text
		if url == "" {
			output.SetText("Please enter a target URL.")
			return
		}
		// Add your web app penetration testing code here.
		output.SetText(fmt.Sprintf("Scanning URL: %s", url))
	})

	content := container.NewVBox(
		urlEntry,
		scanButton,
		widget.NewLabel("Output:"),
		output,
	)

	window.SetContent(content)
}
