package report

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

type TestResult struct {
	Name        string
	Description string
	Severity    string
	Methodology string
}

func GenerateReport(results []TestResult, outputPath string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetTitle("Security Report", false)
	pdf.SetAuthor("Your Application Name", false)

	for _, result := range results {
		pdf.AddPage()
		pdf.SetFont("Arial", "B", 16)
		pdf.Cell(0, 10, fmt.Sprintf("Name: %s", result.Name))
		pdf.Ln(20)

		pdf.SetFont("Arial", "", 12)
		pdf.Cell(0, 10, fmt.Sprintf("Description: %s", result.Description))
		pdf.Ln(20)

		pdf.Cell(0, 10, fmt.Sprintf("Severity: %s", result.Severity))
		pdf.Ln(20)

		pdf.Cell(0, 10, "Methodology:")
		pdf.Ln(20)
		pdf.MultiCell(0, 10, result.Methodology, "0", "L", false)
		pdf.Ln(20)
	}

	return pdf.OutputFileAndClose(outputPath)
}
