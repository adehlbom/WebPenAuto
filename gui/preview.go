package gui

import (
	"io/ioutil"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/jung-kurt/gofpdf"
)

// NewPreviewPanel creates a new panel that displays a preview of the PDF report.
func NewPreviewPanel() *fyne.Container {
	// Create a temporary file to store the PDF report.
	tmpfile, err := ioutil.TempFile("", "report*.pdf")
	if err != nil {
		log.Fatal("Error creating temporary file:", err)
	}
	defer os.Remove(tmpfile.Name())

	// Generate a sample PDF report and save it to the temporary file.
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, World!")
	err = pdf.OutputFileAndClose(tmpfile.Name())
	if err != nil {
		log.Fatal("Error generating PDF report:", err)
	}

	// Load the PDF report from the temporary file and create a static resource.
	bytes, err := ioutil.ReadFile(tmpfile.Name())
	if err != nil {
		log.Fatal("Error reading PDF report:", err)
	}
	res := fyne.NewStaticResource("report.pdf", bytes)

	// Create a canvas to display the static resource.
	img := canvas.NewImageFromResource(res)
	img.FillMode = canvas.ImageFillContain
	img.SetMinSize(fyne.NewSize(200, 300))
	canvas := fyne.NewContainerWithLayout(layout.NewMaxLayout(), img)

	// Create a container that holds the canvas and adds a label.
	label := widget.NewLabel("PDF Preview")
	container := container.NewVBox(label, canvas)
	return container
}
