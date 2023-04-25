package gui

import (
	"log"

	"fyne.io/fyne/v2/widget"
	"github.com/adehlbom/WebPenAuto/report"
)

func NewExportButton(results []report.TestResult) *widget.Button {
	exportButton := widget.NewButton("Export PDF", func() {
		err := report.GenerateReport(results, "report.pdf")
		if err != nil {
			log.Println("Error generating PDF report:", err)
		} else {
			log.Println("PDF report generated successfully.")
		}
	})
	return exportButton
}
