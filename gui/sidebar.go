package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func NewSidebar(onSelect func(string)) *widget.List {
	tools := []string{
		"Tool 1",
		"Tool 2",
		"Tool 3",
		// Add more tools here
	}

	toolList := widget.NewList(
		func() int {
			return len(tools)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(tools[id])
		},
	)

	toolList.OnSelected = func(id widget.ListItemID) {
		onSelect(tools[id])
	}

	return toolList
}
