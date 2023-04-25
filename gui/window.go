package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func SetupMainWindow(window fyne.Window, exportButton *widget.Button) {
	urlEntry := NewURLInput()

	output := NewOutputLabel()

	scanButton := NewScanButton(urlEntry, output)

	sidebar := NewSidebar(func(toolName string) {
		output.SetText("Selected tool: " + toolName)
		// You can add code here to handle tool selection and update the main content accordingly
	})
	previewPanel := NewPreviewPanel()

	content := container.NewBorder(
		nil, nil, sidebar, nil,
		container.NewVBox(
			urlEntry,
			scanButton,
			widget.NewLabel("Output:"),
			output,
			previewPanel,
			exportButton,
		),
	)

	window.SetContent(content)
}
