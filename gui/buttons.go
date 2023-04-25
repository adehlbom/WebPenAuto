package gui

import (
	"fmt"

	"fyne.io/fyne/v2/widget"
)

func NewScanButton(urlEntry *widget.Entry, output *widget.Label) *widget.Button {
	scanButton := widget.NewButton("Scan", func() {
		url := urlEntry.Text
		if url == "" {
			output.SetText("Please enter a target URL.")
			return
		}
		// Add your web app penetration testing code here.
		output.SetText(fmt.Sprintf("Scanning URL: %s", url))
	})
	return scanButton
}
