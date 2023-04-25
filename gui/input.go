package gui

import (
	"fyne.io/fyne/v2/widget"
)

func NewURLInput() *widget.Entry {
	urlEntry := widget.NewEntry()
	urlEntry.SetPlaceHolder("Enter target URL")
	return urlEntry
}
